package chat

import (
	"errors"
	"fmt"

	chatapi "github.com/martishin/go-design-patterns/patterns/behavioral/mediator/pkg/chat"
)

var (
	ErrEmptyUserName = errors.New("chat: user name is empty")
	ErrNilMediator   = errors.New("chat: nil mediator")
)

type User struct {
	name  string
	inbox []string
}

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, ErrEmptyUserName
	}

	return &User{name: name}, nil
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Send(room chatapi.Mediator, message string) error {
	if room == nil {
		return ErrNilMediator
	}

	room.Send(u, message)
	return nil
}

func (u *User) Receive(room string, message string) {
	u.inbox = append(u.inbox, fmt.Sprintf("[%s] %s receives: %s", room, u.name, message))
}

func (u *User) Inbox() []string {
	return append([]string(nil), u.inbox...)
}
