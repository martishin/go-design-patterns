package storage

type Uploader interface {
	Upload(path string, data []byte) error
}

type BackupService struct {
	uploader Uploader
}

func (bs *BackupService) Backup(path string, data []byte) error {
	return bs.uploader.Upload(path, data)
}

func NewBackupService(uploader Uploader) *BackupService {
	return &BackupService{
		uploader: uploader,
	}
}
