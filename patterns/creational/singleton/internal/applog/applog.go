package applog

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Logger struct {
	mu  sync.Mutex
	out io.Writer
}

var (
	once     sync.Once
	instance *Logger
)

func Get() *Logger {
	created := false

	once.Do(func() {
		created = true
		fmt.Println("Creating single instance now.")
		instance = &Logger{out: os.Stdout}
	})

	if !created {
		fmt.Println("Single instance already created.")
	}

	return instance
}

func (l *Logger) Info(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	ts := time.Now().UTC().Format(time.RFC3339Nano)
	_, _ = fmt.Fprintf(l.out, "ts=%s level=INFO msg=%q\n", ts, msg)
}
