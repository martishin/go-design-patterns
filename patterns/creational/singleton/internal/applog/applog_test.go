package applog_test

import (
	"sync"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/creational/singleton/internal/applog"
)

func TestGet_ReturnsSamePointer(t *testing.T) {
	a := applog.Get()
	b := applog.Get()

	if a == nil || b == nil {
		t.Fatalf("expected non-nil logger")
	}
	if a != b {
		t.Fatalf("expected same logger pointer, got %p and %p", a, b)
	}
}

func TestGet_ReturnsSamePointer_Concurrent(t *testing.T) {
	const goroutines = 64

	base := applog.Get()

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for range goroutines {
		go func() {
			defer wg.Done()
			got := applog.Get()
			if got != base {
				t.Errorf("expected same logger pointer, got %p and %p", got, base)
			}
		}()
	}

	wg.Wait()
}
