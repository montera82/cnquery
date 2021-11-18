package aws

import (
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"go.mondoo.io/mondoo/lumi/resources/smbios"
	"go.mondoo.io/mondoo/motor/motorid/awsec2"
	"go.mondoo.io/mondoo/motor/platform"
	"go.mondoo.io/mondoo/motor/transports"
)

const (
	awsIdentifierFileLinux = "/sys/class/dmi/id/product_version"
)

func Detect(t transports.Transport, p *platform.Platform) string {
	productVersion := ""
	if p.IsFamily("linux") {
		// Fetching the product version from the smbios manager is slow
		// because it iterates through files we don't need to check. This
		// is an optimzation for our sshfs. Also, be aware that on linux,
		// you may not have access to all the smbios things under /sys, so
		// you want to make sure to only check the product_version file
		content, err := afero.ReadFile(t.FS(), awsIdentifierFileLinux)
		if err != nil {
			log.Debug().Err(err).Msgf("unable to read %s", awsIdentifierFileLinux)
			return ""
		}
		productVersion = string(content)
	} else {
		mgr, err := smbios.ResolveManager(t, p)
		if err != nil {
			return ""
		}
		info, err := mgr.Info()
		if err != nil {
			log.Debug().Err(err).Msg("failed to query smbios")
			return ""
		}
		productVersion = info.SysInfo.Version
	}

	if strings.Contains(productVersion, "amazon") {
		mdsvc, err := awsec2.Resolve(t, p)
		if err != nil {
			log.Debug().Err(err).Msg("failed to get metadata resolver")
			return ""
		}
		id, err := mdsvc.InstanceID()
		if err != nil {
			log.Debug().Err(err).
				Str("transport", t.Kind().String()).
				Strs("platform", p.GetFamily()).
				Msg("failed to get aws platform id")
			return ""
		}
		return id
	}

	return ""
}