package editorcommand

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/pkg/command"
)

type CopyTextCommand struct {
	editor *editor.Editor
	start  int
	end    int
}

func NewCopyTextCommand(editor *editor.Editor, start int, end int) command.Command {
	return &CopyTextCommand{
		editor: editor,
		start:  start,
		end:    end,
	}
}

func (c *CopyTextCommand) Name() string {
	return "copy text"
}

func (c *CopyTextCommand) Description() string {
	return fmt.Sprintf("copy range [%d:%d] to clipboard", c.start, c.end)
}

func (c *CopyTextCommand) Execute() (bool, error) {
	selection, err := c.editor.Selection(c.start, c.end)
	if err != nil {
		return false, err
	}

	c.editor.SetClipboard(selection)
	return false, nil
}

func (c *CopyTextCommand) Undo() {}
