package core

import (
	"go.mondoo.io/mondoo/motor/providers/container/docker_engine"
	"go.mondoo.io/mondoo/motor/providers/container/docker_snapshot"
	"go.mondoo.io/mondoo/motor/providers/tar"
)

func (v *mqlPlatformVirtualization) id() (string, error) {
	return "platform.virtualization", nil
}

func (v *mqlPlatformVirtualization) GetIsContainer() (bool, error) {
	switch v.MotorRuntime.Motor.Provider.(type) {
	case *tar.Provider:
		return true, nil
	case *docker_snapshot.DockerSnapshotProvider:
		return true, nil
	case *docker_engine.Provider:
		return true, nil
	}

	return false, nil
}
