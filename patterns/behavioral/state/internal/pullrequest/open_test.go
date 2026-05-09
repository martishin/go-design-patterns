package pullrequest_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/state/internal/pullrequest"
)

func TestOpenState_ApproveTransitionsToApproved(t *testing.T) {
	pullRequest := openTestPullRequest(t)

	if err := pullRequest.Approve(); err != nil {
		t.Fatalf("Approve() returned error: %v", err)
	}

	if pullRequest.StateName() != "approved" {
		t.Fatalf("got state %q, want %q", pullRequest.StateName(), "approved")
	}
}

func TestOpenState_RequestChangesTransitionsToDraft(t *testing.T) {
	pullRequest := openTestPullRequest(t)

	if err := pullRequest.RequestChanges(); err != nil {
		t.Fatalf("RequestChanges() returned error: %v", err)
	}

	if pullRequest.StateName() != "draft" {
		t.Fatalf("got state %q, want %q", pullRequest.StateName(), "draft")
	}
}

func TestOpenState_RejectsInvalidActions(t *testing.T) {
	tests := []struct {
		name   string
		action func(*pullrequest.PullRequest) error
	}{
		{name: "open", action: (*pullrequest.PullRequest).Open},
		{name: "merge", action: (*pullrequest.PullRequest).Merge},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pullRequest := openTestPullRequest(t)

			err := tt.action(pullRequest)
			if !errors.Is(err, pullrequest.ErrInvalidTransition) {
				t.Fatalf("got error %v, want %v", err, pullrequest.ErrInvalidTransition)
			}
			if pullRequest.StateName() != "open" {
				t.Fatalf("got state %q, want %q", pullRequest.StateName(), "open")
			}
		})
	}
}

func openTestPullRequest(t *testing.T) *pullrequest.PullRequest {
	t.Helper()

	pullRequest := newTestPullRequest(t)
	if err := pullRequest.Open(); err != nil {
		t.Fatalf("Open() returned error: %v", err)
	}

	return pullRequest
}
