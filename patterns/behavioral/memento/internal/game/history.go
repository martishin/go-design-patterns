package game

import "github.com/martishin/go-design-patterns/patterns/behavioral/memento/pkg/checkpoint"

type History struct {
	checkpoints []checkpoint.Checkpoint
}

func NewHistory() *History {
	return &History{}
}

func (h *History) Push(saved checkpoint.Checkpoint) {
	h.checkpoints = append(h.checkpoints, saved)
}

func (h *History) Pop() checkpoint.Checkpoint {
	if len(h.checkpoints) == 0 {
		return nil
	}

	lastIndex := len(h.checkpoints) - 1
	saved := h.checkpoints[lastIndex]
	h.checkpoints = h.checkpoints[:lastIndex]

	return saved
}

func (h *History) Len() int {
	return len(h.checkpoints)
}
