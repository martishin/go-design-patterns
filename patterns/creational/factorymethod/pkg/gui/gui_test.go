package gui_test

import (
	"io"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/creational/factorymethod/pkg/gui"
)

type fakeButton struct {
	rendered bool
	clicked  bool
}

func (f *fakeButton) Render(_ io.Writer) {
	f.rendered = true
	f.OnClick(io.Discard)
}

func (f *fakeButton) OnClick(_ io.Writer) {
	f.clicked = true
}

func TestDialogBase_RenderInvokesFactoryAndProduct(t *testing.T) {
	f := &fakeButton{}
	d := gui.DialogBase{
		Create: func() gui.Button {
			return f
		},
	}

	d.Render(io.Discard)

	if !f.rendered {
		t.Fatal("expected Button.Render to be called")
	}
	if !f.clicked {
		t.Fatal("expected Button.OnClick to be called")
	}
}
