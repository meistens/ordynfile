package parser

import "ordynfile/internal/models"

// Parse fills ParsedTitle, ParsedYear, ParsedSeason, ParsedEpisode, Type
func Parse(file *models.MediaFile) {
	// TODO: implement actual regex + logic
	file.ParsedTitle = file.Filename // placeholder
	file.Type = models.Movie         // placeholder
}
