package editorcommand

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/pkg/command"
)

type DeleteTextCommand struct {
	baseCommand
	start int
	end   int
}

func NewDeleteTextCommand(editor *editor.Editor, start int, end int) command.Command {
	return &DeleteTextCommand{
		baseCommand: newBaseCommand(editor),
		start:       start,
		end:         end,
	}
}

func (c *DeleteTextCommand) Name() string {
	return "delete text"
}

func (c *DeleteTextCommand) Description() string {
	return fmt.Sprintf("delete range [%d:%d]", c.start, c.end)
}

func (c *DeleteTextCommand) Execute() (bool, error) {
	c.saveBackup()

	if _, err := c.editor.Delete(c.start, c.end); err != nil {
		return false, err
	}

	return true, nil
}
