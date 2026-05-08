package chat_test

import (
	"errors"
	"reflect"
	"testing"

	internalchat "github.com/martishin/go-design-patterns/patterns/behavioral/mediator/internal/chat"
	chatapi "github.com/martishin/go-design-patterns/patterns/behavioral/mediator/pkg/chat"
)

type mediatorSpy struct {
	name    string
	sender  chatapi.Participant
	message string
}

func (m *mediatorSpy) Name() string {
	return m.name
}

func (m *mediatorSpy) Send(sender chatapi.Participant, message string) {
	m.sender = sender
	m.message = message
}

func TestNewUser_ReturnsErrorForEmptyName(t *testing.T) {
	_, err := internalchat.NewUser("")

	if !errors.Is(err, internalchat.ErrEmptyUserName) {
		t.Fatalf("got error %v, want %v", err, internalchat.ErrEmptyUserName)
	}
}

func TestUser_ImplementsParticipant(t *testing.T) {
	user, err := internalchat.NewUser("alice")
	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}

	var participant chatapi.Participant = user
	if participant.Name() != "alice" {
		t.Fatalf("got participant name %q, want %q", participant.Name(), "alice")
	}
}

func TestUser_SendDelegatesToMediator(t *testing.T) {
	user, err := internalchat.NewUser("alice")
	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}

	room := &mediatorSpy{name: "general"}
	if err := user.Send(room, "Hi team"); err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}

	if room.sender != user {
		t.Fatalf("got sender %v, want %v", room.sender, user)
	}
	if room.message != "Hi team" {
		t.Fatalf("got message %q, want %q", room.message, "Hi team")
	}
}

func TestUser_SendReturnsErrorForNilMediator(t *testing.T) {
	user, err := internalchat.NewUser("alice")
	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}

	err = user.Send(nil, "Hi team")

	if !errors.Is(err, internalchat.ErrNilMediator) {
		t.Fatalf("got error %v, want %v", err, internalchat.ErrNilMediator)
	}
}

func TestUser_ReceiveStoresRoomMessage(t *testing.T) {
	user, err := internalchat.NewUser("bob")
	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}

	user.Receive("general", "Hi team")

	want := []string{"[general] bob receives: Hi team"}
	if !reflect.DeepEqual(user.Inbox(), want) {
		t.Fatalf("got inbox %v, want %v", user.Inbox(), want)
	}
}

func TestUser_CanParticipateInMultipleRooms(t *testing.T) {
	alice, err := internalchat.NewUser("alice")
	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}
	bob, err := internalchat.NewUser("bob")
	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}
	carol, err := internalchat.NewUser("carol")
	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}

	team, err := internalchat.NewRoom("team")
	if err != nil {
		t.Fatalf("NewRoom() returned error: %v", err)
	}
	incident, err := internalchat.NewRoom("incident")
	if err != nil {
		t.Fatalf("NewRoom() returned error: %v", err)
	}

	for _, participant := range []chatapi.Participant{alice, bob} {
		if err := team.Join(participant); err != nil {
			t.Fatalf("Join() returned error: %v", err)
		}
	}
	for _, participant := range []chatapi.Participant{alice, bob, carol} {
		if err := incident.Join(participant); err != nil {
			t.Fatalf("Join() returned error: %v", err)
		}
	}

	if err := alice.Send(team, "Hi team"); err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}
	if err := alice.Send(incident, "API is down"); err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}
	if err := bob.Send(team, "Yes, already looking into it"); err != nil {
		t.Fatalf("Send() returned error: %v", err)
	}

	wantBobInbox := []string{
		"[team] bob receives: Hi team",
		"[incident] bob receives: API is down",
	}
	if !reflect.DeepEqual(bob.Inbox(), wantBobInbox) {
		t.Fatalf("got bob inbox %v, want %v", bob.Inbox(), wantBobInbox)
	}

	wantCarolInbox := []string{
		"[incident] carol receives: API is down",
	}
	if !reflect.DeepEqual(carol.Inbox(), wantCarolInbox) {
		t.Fatalf("got carol inbox %v, want %v", carol.Inbox(), wantCarolInbox)
	}
}

func TestUser_InboxReturnsCopy(t *testing.T) {
	user, err := internalchat.NewUser("bob")
	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}

	user.Receive("general", "Hi team")
	inbox := user.Inbox()
	inbox[0] = "changed"

	if user.Inbox()[0] != "[general] bob receives: Hi team" {
		t.Fatalf("inbox was mutated through returned slice: %v", user.Inbox())
	}
}
