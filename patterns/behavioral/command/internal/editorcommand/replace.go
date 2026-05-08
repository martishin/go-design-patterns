package editorcommand

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/pkg/command"
)

type ReplaceTextCommand struct {
	baseCommand
	start int
	end   int
	text  string
}

func NewReplaceTextCommand(editor *editor.Editor, start int, end int, text string) command.Command {
	return &ReplaceTextCommand{
		baseCommand: newBaseCommand(editor),
		start:       start,
		end:         end,
		text:        text,
	}
}

func (c *ReplaceTextCommand) Name() string {
	return "replace text"
}

func (c *ReplaceTextCommand) Description() string {
	return fmt.Sprintf("replace range [%d:%d] with %q", c.start, c.end, c.text)
}

func (c *ReplaceTextCommand) Execute() (bool, error) {
	c.saveBackup()

	if err := c.editor.Replace(c.start, c.end, c.text); err != nil {
		return false, err
	}

	return true, nil
}
