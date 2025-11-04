package htmlgui

import (
	"fmt"
	"io"

	"github.com/martishin/go-design-patterns/patterns/creational/factorymethod/pkg/gui"
)

// HTMLButton is a concrete product.
type HTMLButton struct{}

func (b HTMLButton) Render(w io.Writer) {
	fmt.Fprintln(w, "<button>Test Button</button>")
	b.OnClick(w)
}

func (b HTMLButton) OnClick(w io.Writer) {
	fmt.Fprintln(w, "Click! Button says - 'Hello World!'")
}

// HTMLDialog is a concrete creator.
type HTMLDialog struct{ gui.DialogBase }

// New returns a Dialog that creates HTML buttons.
func New() gui.Dialog {
	return HTMLDialog{
		DialogBase: gui.DialogBase{
			Create: func() gui.Button {
				return HTMLButton{}
			},
		},
	}
}
