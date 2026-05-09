package main

import (
	"fmt"
	"log"

	"github.com/martishin/go-design-patterns/patterns/behavioral/state/internal/pullrequest"
)

func main() {
	pr, err := pullrequest.NewPullRequest("Add cache invalidation")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current state: %s\n", pr.StateName())

	runAction("Open pull request", pr.Open, pr)
	runAction("Merge pull request", pr.Merge, pr)
	runAction("Approve pull request", pr.Approve, pr)
	runAction("Merge pull request", pr.Merge, pr)
	runAction("Request changes", pr.RequestChanges, pr)
}

func runAction(label string, action func() error, pr *pullrequest.PullRequest) {
	if err := action(); err != nil {
		fmt.Printf("%s -> error: %s\n", label, err)
		return
	}

	fmt.Printf("%s -> state: %s\n", label, pr.StateName())
}
