package organizer

import (
	"os"
)

func MoveFile(src, dst string, dry bool) error {
	if dry {
		return nil
	}
	return os.Rename(src, dst)
}
