package chat_test

import (
	"errors"
	"reflect"
	"testing"

	internalchat "github.com/martishin/go-design-patterns/patterns/behavioral/mediator/internal/chat"
	chatapi "github.com/martishin/go-design-patterns/patterns/behavioral/mediator/pkg/chat"
)

type participantSpy struct {
	name     string
	received []delivery
}

type delivery struct {
	room    string
	message string
}

func (p *participantSpy) Name() string {
	return p.name
}

func (p *participantSpy) Receive(room string, message string) {
	p.received = append(p.received, delivery{room: room, message: message})
}

func TestNewRoom_ReturnsErrorForEmptyName(t *testing.T) {
	_, err := internalchat.NewRoom("")

	if !errors.Is(err, internalchat.ErrEmptyRoomName) {
		t.Fatalf("got error %v, want %v", err, internalchat.ErrEmptyRoomName)
	}
}

func TestRoom_ImplementsMediator(t *testing.T) {
	room, err := internalchat.NewRoom("general")
	if err != nil {
		t.Fatalf("NewRoom() returned error: %v", err)
	}

	var mediator chatapi.Mediator = room
	if mediator.Name() != "general" {
		t.Fatalf("got mediator name %q, want %q", mediator.Name(), "general")
	}
}

func TestRoom_JoinReturnsErrorForNilParticipant(t *testing.T) {
	room, err := internalchat.NewRoom("general")
	if err != nil {
		t.Fatalf("NewRoom() returned error: %v", err)
	}

	err = room.Join(nil)

	if !errors.Is(err, internalchat.ErrNilParticipant) {
		t.Fatalf("got error %v, want %v", err, internalchat.ErrNilParticipant)
	}
}

func TestRoom_SendBroadcastsToParticipantsExceptSender(t *testing.T) {
	room, err := internalchat.NewRoom("general")
	if err != nil {
		t.Fatalf("NewRoom() returned error: %v", err)
	}

	alice := &participantSpy{name: "alice"}
	bob := &participantSpy{name: "bob"}
	carol := &participantSpy{name: "carol"}

	if err := room.Join(alice); err != nil {
		t.Fatalf("Join() returned error: %v", err)
	}
	if err := room.Join(bob); err != nil {
		t.Fatalf("Join() returned error: %v", err)
	}
	if err := room.Join(carol); err != nil {
		t.Fatalf("Join() returned error: %v", err)
	}

	room.Send(alice, "Hi team")

	if len(alice.received) != 0 {
		t.Fatalf("sender received messages %v, want none", alice.received)
	}

	want := []delivery{{room: "general", message: "Hi team"}}
	if !reflect.DeepEqual(bob.received, want) {
		t.Fatalf("got bob messages %v, want %v", bob.received, want)
	}
	if !reflect.DeepEqual(carol.received, want) {
		t.Fatalf("got carol messages %v, want %v", carol.received, want)
	}
}

func TestRoom_HistoryReturnsSentMessages(t *testing.T) {
	room, err := internalchat.NewRoom("general")
	if err != nil {
		t.Fatalf("NewRoom() returned error: %v", err)
	}

	alice := &participantSpy{name: "alice"}
	room.Send(alice, "Hi team")

	want := []string{"[general] alice sends: Hi team"}
	if !reflect.DeepEqual(room.History(), want) {
		t.Fatalf("got history %v, want %v", room.History(), want)
	}
}

func TestRoom_HistoryReturnsCopy(t *testing.T) {
	room, err := internalchat.NewRoom("general")
	if err != nil {
		t.Fatalf("NewRoom() returned error: %v", err)
	}

	alice := &participantSpy{name: "alice"}
	room.Send(alice, "Hi team")

	history := room.History()
	history[0] = "changed"

	if room.History()[0] != "[general] alice sends: Hi team" {
		t.Fatalf("history was mutated through returned slice: %v", room.History())
	}
}
