package editorcommand

import "github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"

type baseCommand struct {
	editor *editor.Editor
	backup string
}

func newBaseCommand(editor *editor.Editor) baseCommand {
	return baseCommand{editor: editor}
}

func (c *baseCommand) saveBackup() {
	c.backup = c.editor.Text()
}

func (c *baseCommand) Undo() {
	c.editor.SetText(c.backup)
}
