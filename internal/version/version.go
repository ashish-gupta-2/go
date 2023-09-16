package version

var (
	name    = "Test"
	version = "v0.0.0"
)

// GetName returns application name.
func GetName() string {
	return name
}

// GetVersion returns application version.
func GetVersion() string {
	return version
}
