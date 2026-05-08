package game_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/memento/internal/game"
	"github.com/martishin/go-design-patterns/patterns/behavioral/memento/pkg/checkpoint"
)

type foreignCheckpoint struct{}

func (f foreignCheckpoint) Name() string {
	return "foreign"
}

func TestNewPlayer_ReturnsErrorForInvalidState(t *testing.T) {
	tests := []struct {
		name    string
		player  string
		level   int
		health  int
		wantErr error
	}{
		{name: "empty name", player: "", level: 1, health: 100, wantErr: game.ErrEmptyPlayerName},
		{name: "invalid level", player: "Rook", level: 0, health: 100, wantErr: game.ErrInvalidLevel},
		{name: "invalid health", player: "Rook", level: 1, health: 101, wantErr: game.ErrInvalidHealth},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := game.NewPlayer(tt.player, tt.level, tt.health)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got error %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlayer_SaveReturnsCheckpointMetadata(t *testing.T) {
	player := newTestPlayer(t)

	saved := player.Save("before boss fight")

	var metadata checkpoint.Checkpoint = saved
	if metadata.Name() != "before boss fight" {
		t.Fatalf("got checkpoint name %q, want %q", metadata.Name(), "before boss fight")
	}
}

func TestPlayer_RestoreRestoresSavedState(t *testing.T) {
	player := newTestPlayer(t)
	saved := player.Save("before boss fight")

	player.LevelUp()
	player.TakeDamage(70)
	player.AddItem("ancient key")

	if err := player.Restore(saved); err != nil {
		t.Fatalf("Restore() returned error: %v", err)
	}

	if player.Level() != 3 {
		t.Fatalf("got level %d, want %d", player.Level(), 3)
	}
	if player.Health() != 80 {
		t.Fatalf("got health %d, want %d", player.Health(), 80)
	}

	wantInventory := []string{"iron sword", "health potion"}
	if !reflect.DeepEqual(player.Inventory(), wantInventory) {
		t.Fatalf("got inventory %v, want %v", player.Inventory(), wantInventory)
	}
}

func TestPlayer_RestoreReturnsErrorForForeignCheckpoint(t *testing.T) {
	player := newTestPlayer(t)

	err := player.Restore(foreignCheckpoint{})

	if !errors.Is(err, game.ErrInvalidSnapshot) {
		t.Fatalf("got error %v, want %v", err, game.ErrInvalidSnapshot)
	}
}

func TestPlayer_InventoryReturnsCopy(t *testing.T) {
	player := newTestPlayer(t)

	inventory := player.Inventory()
	inventory[0] = "mutated"

	if player.Inventory()[0] != "iron sword" {
		t.Fatalf("inventory was mutated through returned slice: %v", player.Inventory())
	}
}

func TestPlayer_SavedCheckpointIsImmutable(t *testing.T) {
	player := newTestPlayer(t)
	saved := player.Save("before boss fight")

	inventory := player.Inventory()
	inventory[0] = "mutated"
	player.AddItem("ancient key")

	if err := player.Restore(saved); err != nil {
		t.Fatalf("Restore() returned error: %v", err)
	}

	wantInventory := []string{"iron sword", "health potion"}
	if !reflect.DeepEqual(player.Inventory(), wantInventory) {
		t.Fatalf("got inventory %v, want %v", player.Inventory(), wantInventory)
	}
}

func TestPlayer_StringReturnsReadableState(t *testing.T) {
	player := newTestPlayer(t)

	want := "Rook: level=3, health=80, inventory=[health potion, iron sword]"
	if player.String() != want {
		t.Fatalf("got string %q, want %q", player.String(), want)
	}
}

func newTestPlayer(t *testing.T) *game.Player {
	t.Helper()

	player, err := game.NewPlayer("Rook", 3, 80, "iron sword", "health potion")
	if err != nil {
		t.Fatalf("NewPlayer() returned error: %v", err)
	}

	return player
}
