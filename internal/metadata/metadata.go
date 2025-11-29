package metadata

import "ordynfile/internal/models"

// Resolve determines if cache hit or API needed
func Resolve(file models.MediaFile) (models.Metadata, error) {
	// TODO: implement cache lookup

	// TODO: call TMDB or AniList

	// TEMPORARY placeholder
	return models.Metadata{
		Title: file.ParsedTitle,
		Type:  file.Type,
	}, nil
}
