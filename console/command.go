package console

// Command is the interface for a command.
type Command interface {
	// Signature returns the signature of the command.
	Signature() string
	// Description returns the description of the command.
	Description() string
	// Help returns the help information of the command.
	Help() string
	// Example returns the example of the command.
	Example() string
	// Group returns the group of the command.
	Group() string
	// Handle handles the command.
	Handle(input Input)
	// Args returns the excepted args of the command.
	Args() func(input Input) error
	// Flags returns the flags of the command.
	Flags() []Flag
	// PersistentFlags returns the persistent flags of the command.
	PersistentFlags() []Flag
}

type Group interface {
	// ID returns the id of the group.
	ID() string
	// Name returns the name of the group.
	Name() string
	// Commands returns the commands of the group.
	Commands() []Command
}
