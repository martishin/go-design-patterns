package pullrequest

type MergedState struct {
	pullRequest *PullRequest
}

func (s *MergedState) Name() string {
	return "merged"
}

func (s *MergedState) Open() error {
	return invalidTransition(s.Name(), "open")
}

func (s *MergedState) Approve() error {
	return invalidTransition(s.Name(), "approve")
}

func (s *MergedState) RequestChanges() error {
	return invalidTransition(s.Name(), "request changes")
}

func (s *MergedState) Merge() error {
	return invalidTransition(s.Name(), "merge")
}
