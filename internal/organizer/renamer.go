package organizer

import (
	"fmt"

	"ordynfile/internal/models"
)

func BuildFilename(metadata models.Metadata, ext string) string {
	// TODO: rules for TV, Anime, Movies

	// Placeholder
	return fmt.Sprintf("%s%s", metadata.Title, ext)
}
