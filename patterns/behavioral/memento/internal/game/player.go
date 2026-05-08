package game

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/martishin/go-design-patterns/patterns/behavioral/memento/pkg/checkpoint"
)

var (
	ErrEmptyPlayerName = errors.New("game: player name is empty")
	ErrInvalidHealth   = errors.New("game: health must be between 0 and 100")
	ErrInvalidLevel    = errors.New("game: level must be positive")
	ErrInvalidSnapshot = errors.New("game: invalid checkpoint")
)

type Player struct {
	name      string
	level     int
	health    int
	inventory []string
}

type playerCheckpoint struct {
	name      string
	level     int
	health    int
	inventory []string
}

func NewPlayer(name string, level int, health int, inventory ...string) (*Player, error) {
	if name == "" {
		return nil, ErrEmptyPlayerName
	}
	if level <= 0 {
		return nil, ErrInvalidLevel
	}
	if health < 0 || health > 100 {
		return nil, ErrInvalidHealth
	}

	return &Player{
		name:      name,
		level:     level,
		health:    health,
		inventory: append([]string(nil), inventory...),
	}, nil
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Level() int {
	return p.level
}

func (p *Player) Health() int {
	return p.health
}

func (p *Player) Inventory() []string {
	return append([]string(nil), p.inventory...)
}

func (p *Player) LevelUp() {
	p.level++
}

func (p *Player) TakeDamage(damage int) {
	p.health -= damage
	if p.health < 0 {
		p.health = 0
	}
}

func (p *Player) Heal(amount int) {
	p.health += amount
	if p.health > 100 {
		p.health = 100
	}
}

func (p *Player) AddItem(item string) {
	p.inventory = append(p.inventory, item)
}

func (p *Player) Save(name string) checkpoint.Checkpoint {
	return &playerCheckpoint{
		name:      name,
		level:     p.level,
		health:    p.health,
		inventory: append([]string(nil), p.inventory...),
	}
}

func (p *Player) Restore(saved checkpoint.Checkpoint) error {
	snapshot, ok := saved.(*playerCheckpoint)
	if !ok {
		return ErrInvalidSnapshot
	}

	p.level = snapshot.level
	p.health = snapshot.health
	p.inventory = append([]string(nil), snapshot.inventory...)

	return nil
}

func (p *Player) String() string {
	inventory := append([]string(nil), p.inventory...)
	sort.Strings(inventory)

	return fmt.Sprintf(
		"%s: level=%d, health=%d, inventory=[%s]",
		p.name,
		p.level,
		p.health,
		strings.Join(inventory, ", "),
	)
}

func (c *playerCheckpoint) Name() string {
	return c.name
}
