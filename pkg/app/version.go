package app

import (
	"ashish.com/m/internal/version"
	"ashish.com/m/pkg/models"
)

// Version returns application build and version info.
func Version() models.Version {
	return models.Version{
		Name:    version.GetName(),
		Version: version.GetVersion(),
	}
}
