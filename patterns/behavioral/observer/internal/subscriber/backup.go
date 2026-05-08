package subscriber

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/pkg/filewatcher"
)

type BackupService struct {
	id      string
	backups []string
}

func NewBackupService(id string) *BackupService {
	return &BackupService{id: id}
}

func (b *BackupService) ID() string {
	return b.id
}

func (b *BackupService) Update(event filewatcher.Event) {
	if event.Operation != "write" {
		return
	}

	b.backups = append(b.backups, fmt.Sprintf("backup created for %s at revision %d", event.Path, event.Revision))
}

func (b *BackupService) Backups() []string {
	return append([]string(nil), b.backups...)
}
