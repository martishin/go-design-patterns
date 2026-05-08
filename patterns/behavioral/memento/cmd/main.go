package main

import (
	"fmt"
	"log"

	"github.com/martishin/go-design-patterns/patterns/behavioral/memento/internal/game"
)

func main() {
	player, err := game.NewPlayer("Rook", 3, 80, "iron sword", "health potion")
	if err != nil {
		log.Fatal(err)
	}

	history := game.NewHistory()

	fmt.Printf("Current state: %s\n", player)

	beforeCave := player.Save("before cave")
	history.Push(beforeCave)
	fmt.Printf("Saved checkpoint: %s\n", beforeCave.Name())

	player.LevelUp()
	player.TakeDamage(35)
	player.AddItem("ancient key")
	fmt.Printf("Current state: %s\n", player)

	beforeBoss := player.Save("before boss fight")
	history.Push(beforeBoss)
	fmt.Printf("Saved checkpoint: %s\n", beforeBoss.Name())

	player.TakeDamage(60)
	player.AddItem("dragon scale")
	fmt.Printf("Current state: %s\n", player)

	restoreLast(player, history)
	fmt.Printf("Restored state: %s\n", player)

	restoreLast(player, history)
	fmt.Printf("Restored state: %s\n", player)
}

func restoreLast(player *game.Player, history *game.History) {
	saved := history.Pop()
	if saved == nil {
		return
	}

	fmt.Printf("Restoring checkpoint: %s\n", saved.Name())
	if err := player.Restore(saved); err != nil {
		log.Fatal(err)
	}
}
