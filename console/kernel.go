package console

type Kernel interface {
	AddGroup(id, name string, fn func(group Group))
	AddCommand(command Command)
	AddFlag(flag Flag)
	AddPersistentFlag(flag Flag)
	Call(cmd string, args ...string) error
	Run() error
}
