package windowsgui

import (
	"fmt"
	"io"

	"github.com/martishin/go-design-patterns/patterns/creational/factorymethod/pkg/gui"
)

// WindowsButton is a concrete product.
type WindowsButton struct{}

func (b WindowsButton) Render(w io.Writer) {
	fmt.Fprintln(w, "Drawing a Windows button")
	b.OnClick(w)
}

func (b WindowsButton) OnClick(w io.Writer) {
	fmt.Fprintln(w, "Click! Hello, Windows!")
}

// WindowsDialog is a concrete creator.
type WindowsDialog struct {
	gui.DialogBase
}

// New returns a Dialog that creates Windows buttons.
func New() gui.Dialog {
	return WindowsDialog{
		DialogBase: gui.DialogBase{
			Create: func() gui.Button {
				return WindowsButton{}
			},
		},
	}
}
