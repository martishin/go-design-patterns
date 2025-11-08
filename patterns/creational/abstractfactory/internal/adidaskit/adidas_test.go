package adidaskit_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/creational/abstractfactory/internal/adidaskit"
)

func TestAdidasFactory_ProducesAdidasItems(t *testing.T) {
	f := adidaskit.New()

	shoe := f.Shoe(44)
	shirt := f.Shirt(52)

	if shoe.Logo() != "adidas" || shirt.Logo() != "adidas" {
		t.Fatalf("expected adidas logos, got shoe=%q shirt=%q", shoe.Logo(), shirt.Logo())
	}

	if shoe.Size() != 44 || shirt.Size() != 52 {
		t.Fatalf("unexpected sizes: shoe=%q shirt=%q", shoe.Size(), shirt.Size())
	}
}
