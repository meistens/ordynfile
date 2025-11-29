package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	LibraryRoot string `json:"library_root"`

	MoviesDir string `json:"movies_dir"`
	TVDir     string `json:"tv_dir"`
	AnimeDir  string `json:"anime_dir"`
	MusicDir  string `json:"music_dir"`

	EnableNFO     bool `json:"enable_nfo"`
	EnableArtwork bool `json:"enable_artwork"`
	EnableCache   bool `json:"enable_cache"`

	TMDBKey    string `json:"tmdb_key"`
	AniListKey string `json:"anilist_key"`
	MALKey     string `json:"mal_key"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	return &cfg, err
}
