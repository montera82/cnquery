// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

package sbom

//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative sbom.proto

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/mitchellh/hashstructure/v2"
	"go.mondoo.com/cnquery/v10"
	"go.mondoo.com/cnquery/v10/explorer"
	"go.mondoo.com/cnquery/v10/mrn"
)

// SBOMQueryPack is a protobuf message that contains the SBOM query pack
//
//go:embed sbom.mql.yaml
var sbomQueryPack []byte

func QueryPack() (*explorer.Bundle, error) {
	return explorer.BundleFromYAML(sbomQueryPack)
}

// GenerateBom generates a BOM from a cnspec json report collection
func GenerateBom(r *ReportCollectionJson) ([]Sbom, error) {
	if r == nil {
		return nil, nil
	}

	generator := &Generator{
		Vendor:  "Mondoo, Inc.",
		Name:    "cnquery",
		Version: cnquery.Version,
		Url:     "https://mondoo.com",
	}
	now := time.Now().UTC().Format(time.RFC3339)

	boms := []Sbom{}
	for mrn := range r.Assets {
		asset := r.Assets[mrn]

		bom := Sbom{
			Generator: generator,
			Timestamp: now,
			Status:    Status_STATUS_SUCCEEDED,
		}

		bom.Asset = &Asset{
			Name:        asset.Name,
			PlatformIds: nil,
			Platform:    &Platform{},
			Labels:      map[string]string{},
			ExternalIds: []*ExternalID{},
		}

		bom.Packages = []*Package{}

		// extract os packages and python packages
		dataPoints := r.Data[mrn]
		for k := range dataPoints {
			rb := BomReport{}
			err := json.Unmarshal(dataPoints[k], &rb)
			if err != nil {
				return nil, err
			}
			if rb.Asset != nil {
				bom.Asset.Name = rb.Asset.Name
				bom.Asset.Platform.Name = rb.Asset.Platform
				bom.Asset.Platform.Version = rb.Asset.Version
				bom.Asset.Platform.Arch = rb.Asset.Arch
				bom.Asset.Platform.Cpes = rb.Asset.CPEs
				bom.Asset.Platform.Labels = rb.Asset.Labels
				bom.Asset.PlatformIds = enrichPlatformIds(rb.Asset.IDs)
			}
			if rb.Packages != nil {
				for _, pkg := range rb.Packages {
					bom.Packages = append(bom.Packages, &Package{
						Name:    pkg.Name,
						Version: pkg.Version,
						Purl:    pkg.Purl,
						Cpes:    pkg.CPEs,
						Type:    pkg.Format,
					})
				}
			}
			if rb.PythonPackages != nil {
				for _, pkg := range rb.PythonPackages {
					bom.Packages = append(bom.Packages, &Package{
						Name:     pkg.Name,
						Version:  pkg.Version,
						Purl:     pkg.Purl,
						Cpes:     pkg.CPEs,
						Location: pkg.FilePath,
						Type:     "pypi",
					})
				}
			}
		}
		boms = append(boms, bom)
	}
	return boms, nil
}

func (b *Package) Hash() (string, error) {
	hash, err := hashstructure.Hash(b, hashstructure.FormatV2, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%016x", hash), nil
}

// enrichPlatformIds adds the platform id based on cnquery ids
// - AWS EC2 instance ARN
func enrichPlatformIds(ids []string) []string {
	platformIds := []string{}
	for i := range ids {
		platformIds = append(platformIds, ids[i])

		// handle AWS EC2 instance platform identifier and generate AWS ARN as additional identifier
		// EC2 arns have the following format arn:aws:ec2:<REGION>:<ACCOUNT_ID>:instance/<instance-id>
		// //platformid.api.mondoo.app/runtime/aws/ec2/v1/accounts/12345678910/regions/us-east-1/instances/i-1234567890abcdef0
		if strings.HasPrefix(ids[i], "//platformid.api.mondoo.app/runtime/aws/ec2/v1") {
			ec2mrn, err := mrn.NewMRN(ids[i])
			if err != nil {
				continue
			}

			accountID, _ := ec2mrn.ResourceID("accounts")
			region, _ := ec2mrn.ResourceID("regions")
			instanceID, _ := ec2mrn.ResourceID("instances")

			if accountID != "" && region != "" && instanceID != "" {
				platformIds = append(platformIds, fmt.Sprintf("arn:aws:ec2:%s:%s:instance/%s", region, accountID, instanceID))
			}
		}
	}
	return platformIds
}
