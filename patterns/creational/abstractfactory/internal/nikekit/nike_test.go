package nikekit_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/creational/abstractfactory/internal/nikekit"
)

func TestNikeFactory_ProducesNikeItems(t *testing.T) {
	f := nikekit.New()

	shoe := f.Shoe(43)
	shirt := f.Shirt(50)

	if shoe.Logo() != "nike" || shirt.Logo() != "nike" {
		t.Fatalf("expected nike logos, got shoe=%q shirt=%q", shoe.Logo(), shirt.Logo())
	}

	if shoe.Size() != 43 || shirt.Size() != 50 {
		t.Fatalf("unepercted sizes: shoe=%d shirt=%d", shoe.Size(), shirt.Size())
	}
}
