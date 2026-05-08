package game_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/memento/internal/game"
)

func TestHistory_PopReturnsNilWhenEmpty(t *testing.T) {
	history := game.NewHistory()

	if history.Pop() != nil {
		t.Fatal("got checkpoint, want nil")
	}
}

func TestHistory_PopReturnsLastCheckpoint(t *testing.T) {
	history := game.NewHistory()
	player := newTestPlayer(t)
	first := player.Save("before cave")
	second := player.Save("before boss fight")

	history.Push(first)
	history.Push(second)

	got := history.Pop()
	if got != second {
		t.Fatalf("got checkpoint %v, want %v", got, second)
	}
	if history.Len() != 1 {
		t.Fatalf("got history length %d, want %d", history.Len(), 1)
	}
}
