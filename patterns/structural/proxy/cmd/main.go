package main

import (
	"fmt"
	"log"

	"github.com/martishin/go-design-patterns/patterns/structural/proxy/internal/cachedregistry"
	"github.com/martishin/go-design-patterns/patterns/structural/proxy/internal/remoteregistry"
)

func main() {
	remoteRegistry := remoteregistry.NewRemoteArtifactRegistry("registry.acme.io")
	cachedRegistry, err := cachedregistry.NewCachedArtifactRegistry(remoteRegistry)
	if err != nil {
		log.Fatal(err)
	}

	firstArtifact, err := cachedRegistry.Get("payments-api", "v1.4.2")
	if err != nil {
		log.Fatal(err)
	}

	secondArtifact, err := cachedRegistry.Get("payments-api", "v1.4.2")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("First lookup image: %s\n", firstArtifact.Image)
	fmt.Printf("Second lookup image: %s\n", secondArtifact.Image)
	fmt.Printf("Remote registry calls: %d\n", remoteRegistry.Calls())
	fmt.Printf("Cached artifacts: %d\n", cachedRegistry.CachedArtifacts())
}
