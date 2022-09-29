package commands

type Command interface {
	Invokes() []string
	Description() string
	AdminRequired() bool
	Execute(ctx *Context) error
}