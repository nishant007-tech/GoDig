package plugin

// Plugin defines a common interface.
type Plugin interface {
	Name() string
	Init() error
}
