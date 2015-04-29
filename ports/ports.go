package ports

// A Port is an entry point to a Sheath application, providing a way to interact with the defined use cases.
// It could be an HTTP service, a CLI utility, etc.
type Port interface {
	Expose() error
}
