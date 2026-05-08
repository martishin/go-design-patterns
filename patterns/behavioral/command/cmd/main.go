package main

import (
	"fmt"
	"log"

	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/application"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/internal/editorcommand"
	"github.com/martishin/go-design-patterns/patterns/behavioral/command/pkg/command"
)

func main() {
	document := editor.NewEditor("deploy api")
	app := application.NewApplication(command.NewHistory())

	commands := []command.Command{
		editorcommand.NewInsertTextCommand(document, len(document.Text()), " to production"),
		editorcommand.NewCopyTextCommand(document, 0, 6),
		editorcommand.NewReplaceTextCommand(document, 0, 6, "release"),
	}

	fmt.Printf("Initial text: %q\n", document.Text())

	for _, cmd := range commands {
		fmt.Printf("Executing command: %s - %s\n", cmd.Name(), cmd.Description())
		if err := app.Execute(cmd); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Result text: %q\n", document.Text())
	fmt.Printf("Clipboard: %q\n", document.Clipboard())

	undoneCommand := app.Undo()
	if undoneCommand != nil {
		fmt.Printf("Undo command: %s - %s\n", undoneCommand.Name(), undoneCommand.Description())
	}
	fmt.Printf("Text after undo: %q\n", document.Text())
}
