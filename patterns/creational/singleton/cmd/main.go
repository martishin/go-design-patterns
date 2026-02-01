package main

import (
	"github.com/martishin/go-design-patterns/patterns/creational/singleton/pkg/applog"
)

func main() {
	log := applog.Get()
	log.Info("service starting")

	applog.Get().Info("handling request")
}
