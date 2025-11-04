package main

import (
	"os"

	"github.com/martishin/go-design-patterns/patterns/creational/factorymethod/internal/windowsgui"
)

func main() {
	// dialog := htmlgui.New()
	dialog := windowsgui.New()
	dialog.Render(os.Stdout)
	dialog.Refresh(os.Stdout)
}
