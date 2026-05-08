package main

import (
	"fmt"
	"log"

	internalchat "github.com/martishin/go-design-patterns/patterns/behavioral/mediator/internal/chat"
	chatapi "github.com/martishin/go-design-patterns/patterns/behavioral/mediator/pkg/chat"
)

func main() {
	alice, err := internalchat.NewUser("Alice")
	if err != nil {
		log.Fatal(err)
	}
	bob, err := internalchat.NewUser("Bob")
	if err != nil {
		log.Fatal(err)
	}
	carol, err := internalchat.NewUser("Carol")
	if err != nil {
		log.Fatal(err)
	}

	team, err := internalchat.NewRoom("team")
	if err != nil {
		log.Fatal(err)
	}
	incident, err := internalchat.NewRoom("incident")
	if err != nil {
		log.Fatal(err)
	}

	joinAll(team, alice, bob)
	joinAll(incident, alice, bob, carol)

	sendMessage(incident, alice, "API is down")
	printLatest(bob.Inbox())
	printLatest(carol.Inbox())

	fmt.Println()
	sendMessage(team, alice, "Hi team, API is down")
	printLatest(bob.Inbox())

	fmt.Println()
	sendMessage(team, bob, "Yes, already looking into it")
	printLatest(alice.Inbox())

	fmt.Println()
	printHistory("incident", incident.History())

	fmt.Println()
	printHistory("team", team.History())
}

func joinAll(room *internalchat.Room, participants ...chatapi.Participant) {
	for _, participant := range participants {
		if err := room.Join(participant); err != nil {
			log.Fatal(err)
		}
	}
}

func sendMessage(room *internalchat.Room, sender *internalchat.User, message string) {
	fmt.Printf("[%s] %s sends: %s\n", room.Name(), sender.Name(), message)
	if err := sender.Send(room, message); err != nil {
		log.Fatal(err)
	}
}

func printLatest(messages []string) {
	if len(messages) == 0 {
		return
	}

	fmt.Println(messages[len(messages)-1])
}

func printHistory(room string, history []string) {
	fmt.Printf("%s history:\n", room)
	for _, message := range history {
		fmt.Printf("- %s\n", message)
	}
}
