package scan

import (
	"ordynfile/internal/models"
	"os"
	"path/filepath"
)

func Scan(path string) ([]models.MediaFile, error) {
	var files []models.MediaFile

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(p)
		if isMediaExtension(ext) {
			files = append(files, models.MediaFile{
				FullPath:  p,
				Directory: filepath.Dir(p),
				Filename:  info.Name(),
				Extension: ext,
				Size:      info.Size(),
			})
		}
		return nil
	})

	return files, err
}

func isMediaExtension(ext string) bool {
	switch ext {
	case ".mp4", ".mkv", ".avi", ".mov":
		return true
	default:
		return false
	}
}
