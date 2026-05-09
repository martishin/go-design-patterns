package pullrequest_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/state/internal/pullrequest"
)

func TestDraftState_OpenTransitionsToOpen(t *testing.T) {
	pullRequest := newTestPullRequest(t)

	if err := pullRequest.Open(); err != nil {
		t.Fatalf("Open() returned error: %v", err)
	}

	if pullRequest.StateName() != "open" {
		t.Fatalf("got state %q, want %q", pullRequest.StateName(), "open")
	}
}

func TestDraftState_RejectsInvalidActions(t *testing.T) {
	tests := []struct {
		name   string
		action func(*pullrequest.PullRequest) error
	}{
		{name: "approve", action: (*pullrequest.PullRequest).Approve},
		{name: "request changes", action: (*pullrequest.PullRequest).RequestChanges},
		{name: "merge", action: (*pullrequest.PullRequest).Merge},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pullRequest := newTestPullRequest(t)

			err := tt.action(pullRequest)
			if !errors.Is(err, pullrequest.ErrInvalidTransition) {
				t.Fatalf("got error %v, want %v", err, pullrequest.ErrInvalidTransition)
			}
			if pullRequest.StateName() != "draft" {
				t.Fatalf("got state %q, want %q", pullRequest.StateName(), "draft")
			}
		})
	}
}
