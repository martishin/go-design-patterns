package editorcommand

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/pkg/command"
)

type InsertTextCommand struct {
	baseCommand
	position int
	text     string
}

func NewInsertTextCommand(editor *editor.Editor, position int, text string) command.Command {
	return &InsertTextCommand{
		baseCommand: newBaseCommand(editor),
		position:    position,
		text:        text,
	}
}

func (c *InsertTextCommand) Name() string {
	return "insert text"
}

func (c *InsertTextCommand) Description() string {
	return fmt.Sprintf("insert %q at position %d", c.text, c.position)
}

func (c *InsertTextCommand) Execute() (bool, error) {
	c.saveBackup()

	if err := c.editor.Insert(c.position, c.text); err != nil {
		return false, err
	}

	return true, nil
}
