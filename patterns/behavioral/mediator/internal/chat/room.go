package chat

import (
	"errors"
	"fmt"

	chatapi "github.com/martishin/go-design-patterns/patterns/behavioral/mediator/pkg/chat"
)

var (
	ErrEmptyRoomName  = errors.New("chat: room name is empty")
	ErrNilParticipant = errors.New("chat: nil participant")
)

type Room struct {
	name         string
	participants []chatapi.Participant
	history      []string
}

func NewRoom(name string) (*Room, error) {
	if name == "" {
		return nil, ErrEmptyRoomName
	}

	return &Room{name: name}, nil
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) Join(participant chatapi.Participant) error {
	if participant == nil {
		return ErrNilParticipant
	}

	r.participants = append(r.participants, participant)
	return nil
}

func (r *Room) Send(sender chatapi.Participant, message string) {
	entry := fmt.Sprintf("[%s] %s sends: %s", r.name, sender.Name(), message)
	r.history = append(r.history, entry)

	for _, participant := range r.participants {
		if participant == sender {
			continue
		}

		participant.Receive(r.name, message)
	}
}

func (r *Room) History() []string {
	return append([]string(nil), r.history...)
}
