package application

import "github.com/martishin/go-design-patterns/patterns/behavioral/command/pkg/command"

type Application struct {
	history *command.History
}

func NewApplication(history *command.History) *Application {
	return &Application{history: history}
}

func (a *Application) Execute(command command.Command) error {
	changed, err := command.Execute()
	if err != nil {
		return err
	}

	if changed {
		a.history.Push(command)
	}

	return nil
}

func (a *Application) Undo() command.Command {
	last := a.history.Pop()
	if last == nil {
		return nil
	}

	last.Undo()
	return last
}

func (a *Application) HistoryLen() int {
	return a.history.Len()
}
