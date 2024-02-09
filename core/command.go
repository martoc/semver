package core

// Command is an interface that represents a command.
// It defines the Execute method, which executes the command and returns a string and an error.
type Command interface {
	Execute() (string, error)
}
