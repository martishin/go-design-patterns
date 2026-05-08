package chat_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/mediator/pkg/chat"
)

type participantStub struct{}

func (p participantStub) Name() string {
	return "alice"
}

func (p participantStub) Receive(_ string, _ string) {}

type mediatorStub struct{}

func (m mediatorStub) Name() string {
	return "general"
}

func (m mediatorStub) Send(_ chat.Participant, _ string) {}

func TestParticipant_InterfaceCanBeImplemented(t *testing.T) {
	var participant chat.Participant = participantStub{}

	if participant.Name() != "alice" {
		t.Fatalf("got participant name %q, want %q", participant.Name(), "alice")
	}
}

func TestMediator_InterfaceCanBeImplemented(t *testing.T) {
	var mediator chat.Mediator = mediatorStub{}

	if mediator.Name() != "general" {
		t.Fatalf("got mediator name %q, want %q", mediator.Name(), "general")
	}
}
