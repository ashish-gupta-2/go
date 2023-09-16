package models

// Version represents application build and version info.
type Version struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}
