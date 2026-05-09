package pullrequest

import (
	"errors"
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/behavioral/state/pkg/review"
)

var (
	ErrEmptyTitle        = errors.New("pull request: title is empty")
	ErrInvalidTransition = errors.New("pull request: invalid state transition")
)

type PullRequest struct {
	title    string
	state    review.State
	draft    review.State
	open     review.State
	approved review.State
	merged   review.State
}

func NewPullRequest(title string) (*PullRequest, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}

	pullRequest := &PullRequest{title: title}

	pullRequest.draft = &DraftState{pullRequest: pullRequest}
	pullRequest.open = &OpenState{pullRequest: pullRequest}
	pullRequest.approved = &ApprovedState{pullRequest: pullRequest}
	pullRequest.merged = &MergedState{pullRequest: pullRequest}
	pullRequest.state = pullRequest.draft

	return pullRequest, nil
}

func (p *PullRequest) Title() string {
	return p.title
}

func (p *PullRequest) StateName() string {
	return p.state.Name()
}

func (p *PullRequest) Open() error {
	return p.state.Open()
}

func (p *PullRequest) Approve() error {
	return p.state.Approve()
}

func (p *PullRequest) RequestChanges() error {
	return p.state.RequestChanges()
}

func (p *PullRequest) Merge() error {
	return p.state.Merge()
}

func (p *PullRequest) setState(state review.State) {
	p.state = state
}

func invalidTransition(state string, action string) error {
	return fmt.Errorf("%w: cannot %s while %s", ErrInvalidTransition, action, state)
}
