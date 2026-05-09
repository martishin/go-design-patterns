package pullrequest

type DraftState struct {
	pullRequest *PullRequest
}

func (s *DraftState) Name() string {
	return "draft"
}

func (s *DraftState) Open() error {
	s.pullRequest.setState(s.pullRequest.open)
	return nil
}

func (s *DraftState) Approve() error {
	return invalidTransition(s.Name(), "approve")
}

func (s *DraftState) RequestChanges() error {
	return invalidTransition(s.Name(), "request changes")
}

func (s *DraftState) Merge() error {
	return invalidTransition(s.Name(), "merge")
}
