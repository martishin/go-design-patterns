package pullrequest

type OpenState struct {
	pullRequest *PullRequest
}

func (s *OpenState) Name() string {
	return "open"
}

func (s *OpenState) Open() error {
	return invalidTransition(s.Name(), "open")
}

func (s *OpenState) Approve() error {
	s.pullRequest.setState(s.pullRequest.approved)
	return nil
}

func (s *OpenState) RequestChanges() error {
	s.pullRequest.setState(s.pullRequest.draft)
	return nil
}

func (s *OpenState) Merge() error {
	return invalidTransition(s.Name(), "merge")
}
