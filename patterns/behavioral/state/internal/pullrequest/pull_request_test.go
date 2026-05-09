package pullrequest_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/state/internal/pullrequest"
)

func TestNewPullRequest_ReturnsErrorForEmptyTitle(t *testing.T) {
	_, err := pullrequest.NewPullRequest("")

	if !errors.Is(err, pullrequest.ErrEmptyTitle) {
		t.Fatalf("got error %v, want %v", err, pullrequest.ErrEmptyTitle)
	}
}

func TestNewPullRequest_StartsInDraftState(t *testing.T) {
	pullRequest := newTestPullRequest(t)

	if pullRequest.Title() != "Add cache invalidation" {
		t.Fatalf("got title %q, want %q", pullRequest.Title(), "Add cache invalidation")
	}
	if pullRequest.StateName() != "draft" {
		t.Fatalf("got state %q, want %q", pullRequest.StateName(), "draft")
	}
}

func TestPullRequest_DelegatesActionsToCurrentState(t *testing.T) {
	pullRequest := newTestPullRequest(t)

	if err := pullRequest.Open(); err != nil {
		t.Fatalf("Open() returned error: %v", err)
	}
	if err := pullRequest.Approve(); err != nil {
		t.Fatalf("Approve() returned error: %v", err)
	}
	if err := pullRequest.Merge(); err != nil {
		t.Fatalf("Merge() returned error: %v", err)
	}

	if pullRequest.StateName() != "merged" {
		t.Fatalf("got state %q, want %q", pullRequest.StateName(), "merged")
	}
}

func newTestPullRequest(t *testing.T) *pullrequest.PullRequest {
	t.Helper()

	pullRequest, err := pullrequest.NewPullRequest("Add cache invalidation")
	if err != nil {
		t.Fatalf("NewPullRequest() returned error: %v", err)
	}

	return pullRequest
}
