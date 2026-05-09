package pullrequest

type ApprovedState struct {
	pullRequest *PullRequest
}

func (s *ApprovedState) Name() string {
	return "approved"
}

func (s *ApprovedState) Open() error {
	return invalidTransition(s.Name(), "open")
}

func (s *ApprovedState) Approve() error {
	return invalidTransition(s.Name(), "approve")
}

func (s *ApprovedState) RequestChanges() error {
	s.pullRequest.setState(s.pullRequest.open)
	return nil
}

func (s *ApprovedState) Merge() error {
	s.pullRequest.setState(s.pullRequest.merged)
	return nil
}
