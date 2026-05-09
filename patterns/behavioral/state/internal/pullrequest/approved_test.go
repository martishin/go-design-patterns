package pullrequest_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/state/internal/pullrequest"
)

func TestApprovedState_MergeTransitionsToMerged(t *testing.T) {
	pullRequest := approvedTestPullRequest(t)

	if err := pullRequest.Merge(); err != nil {
		t.Fatalf("Merge() returned error: %v", err)
	}

	if pullRequest.StateName() != "merged" {
		t.Fatalf("got state %q, want %q", pullRequest.StateName(), "merged")
	}
}

func TestApprovedState_RequestChangesTransitionsToOpen(t *testing.T) {
	pullRequest := approvedTestPullRequest(t)

	if err := pullRequest.RequestChanges(); err != nil {
		t.Fatalf("RequestChanges() returned error: %v", err)
	}

	if pullRequest.StateName() != "open" {
		t.Fatalf("got state %q, want %q", pullRequest.StateName(), "open")
	}
}

func TestApprovedState_RejectsInvalidActions(t *testing.T) {
	tests := []struct {
		name   string
		action func(*pullrequest.PullRequest) error
	}{
		{name: "open", action: (*pullrequest.PullRequest).Open},
		{name: "approve", action: (*pullrequest.PullRequest).Approve},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pullRequest := approvedTestPullRequest(t)

			err := tt.action(pullRequest)
			if !errors.Is(err, pullrequest.ErrInvalidTransition) {
				t.Fatalf("got error %v, want %v", err, pullrequest.ErrInvalidTransition)
			}
			if pullRequest.StateName() != "approved" {
				t.Fatalf("got state %q, want %q", pullRequest.StateName(), "approved")
			}
		})
	}
}

func approvedTestPullRequest(t *testing.T) *pullrequest.PullRequest {
	t.Helper()

	pullRequest := openTestPullRequest(t)
	if err := pullRequest.Approve(); err != nil {
		t.Fatalf("Approve() returned error: %v", err)
	}

	return pullRequest
}
