package registry

import "errors"

var (
	ErrInvalidArtifactRequest = errors.New("registry: service and version are required")
)

type Artifact struct {
	Service string
	Version string
	Image   string
	Digest  string
	SizeMB  int
}

type ArtifactRegistry interface {
	Get(service string, version string) (Artifact, error)
}
