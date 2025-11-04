package gui

import (
	"fmt"
	"io"
)

// Button is the abstract product.
type Button interface {
	Render(io.Writer)
	OnClick(io.Writer)
}

// Dialog is the abstract creator.
// Concrete dialogs override creation by supplying Create().
type Dialog interface {
	CreateButton() Button
	Render(io.Writer)
	Refresh(io.Writer)
}

// DialogBase provides the shared template behavior (Render/Refresh)
// and delegates button creation to the injected Create function
// This is the Go way to express "override the factory method".
type DialogBase struct {
	Create func() Button
}

func (d DialogBase) CreateButton() Button {
	return d.Create()
}

func (d DialogBase) Render(w io.Writer) {
	btn := d.CreateButton()
	btn.Render(w)
}

func (DialogBase) Refresh(w io.Writer) {
	fmt.Fprintln(w, "Dialog - Refresh")
}
