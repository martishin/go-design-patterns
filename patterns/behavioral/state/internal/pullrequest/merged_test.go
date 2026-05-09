package pullrequest_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/state/internal/pullrequest"
)

func TestMergedState_RejectsAllActions(t *testing.T) {
	tests := []struct {
		name   string
		action func(*pullrequest.PullRequest) error
	}{
		{name: "open", action: (*pullrequest.PullRequest).Open},
		{name: "approve", action: (*pullrequest.PullRequest).Approve},
		{name: "request changes", action: (*pullrequest.PullRequest).RequestChanges},
		{name: "merge", action: (*pullrequest.PullRequest).Merge},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pullRequest := mergedTestPullRequest(t)

			err := tt.action(pullRequest)
			if !errors.Is(err, pullrequest.ErrInvalidTransition) {
				t.Fatalf("got error %v, want %v", err, pullrequest.ErrInvalidTransition)
			}
			if pullRequest.StateName() != "merged" {
				t.Fatalf("got state %q, want %q", pullRequest.StateName(), "merged")
			}
		})
	}
}

func mergedTestPullRequest(t *testing.T) *pullrequest.PullRequest {
	t.Helper()

	pullRequest := approvedTestPullRequest(t)
	if err := pullRequest.Merge(); err != nil {
		t.Fatalf("Merge() returned error: %v", err)
	}

	return pullRequest
}
