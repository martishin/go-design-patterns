package command

type Command interface {
	Name() string
	Description() string
	Execute() (bool, error)
	Undo()
}
