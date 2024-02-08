package core

type Command interface {
	Execute() (string, error)
}
