package command

type History struct {
	commands []Command
}

func NewHistory() *History {
	return &History{}
}

func (h *History) Push(command Command) {
	h.commands = append(h.commands, command)
}

func (h *History) Pop() Command {
	if len(h.commands) == 0 {
		return nil
	}

	last := h.commands[len(h.commands)-1]
	h.commands = h.commands[:len(h.commands)-1]

	return last
}

func (h *History) Len() int {
	return len(h.commands)
}
