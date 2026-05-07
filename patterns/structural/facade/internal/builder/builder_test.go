package builder_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/structural/facade/internal/builder"
	"github.com/martishin/go-design-patterns/patterns/structural/facade/pkg/deploy"
)

func TestImageBuilder_BuildCreatesArtifactReference(t *testing.T) {
	imageBuilder := builder.NewImageBuilder()

	got, err := imageBuilder.Build(deploy.DeploymentRequest{
		Service:     "payments-api",
		Version:     "v1.4.2",
		Environment: "production",
	})
	if err != nil {
		t.Fatalf("Build() returned error: %v", err)
	}

	want := deploy.Artifact{
		Service: "payments-api",
		Version: "v1.4.2",
		Image:   "payments-api:v1.4.2",
	}
	if got != want {
		t.Fatalf("got artifact %+v, want %+v", got, want)
	}
}
