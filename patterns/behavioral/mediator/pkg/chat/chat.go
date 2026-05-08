package chat

type Mediator interface {
	Name() string
	Send(sender Participant, message string)
}

type Participant interface {
	Name() string
	Receive(room string, message string)
}
