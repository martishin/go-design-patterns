package applog

import (
	"github.com/martishin/go-design-patterns/patterns/creational/singleton/internal/applog"
)

// Logger is a small app-level logger.
type Logger interface {
	Info(msg string)
}

// Get returns the lazily initialized singleton logger.
func Get() Logger {
	return applog.Get()
}
