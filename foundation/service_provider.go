package foundation

// ServiceProvider service provider interface
type ServiceProvider interface {
	// Init init
	Init(app Application)
	// Register register
	Register()
	// Boot boot
	Boot()
}
