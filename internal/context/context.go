package context

const (
	// NamespaceRoot is the root namespace.
	NamespaceRoot = "root"
)

// Key represents the unique context key.
type Key struct{}

// Value represents the context value. It is primarily used for storing
// context values for logging purpose.
type Value struct {
	Namespace string
	TransID   string
	Endpoint  string
}
