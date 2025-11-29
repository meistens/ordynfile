package models

type MediaType int

const (
	Movie MediaType = iota
	TV
	Anime
)

type MediaFile struct {
	FullPath  string
	Directory string
	Filename  string
	Extension string
	Size      int64

	ParsedTitle   string
	ParsedYear    int
	ParsedSeason  int
	ParsedEpisode int

	Type MediaType
}

type Metadata struct {
	ID            string
	Title         string
	OriginalTitle string
	Overview      string
	Year          int
	ReleaseDate   string
	PosterURL     string
	FanartURL     string
	Genres        []string
	Rating        float64

	SeasonNumber  int
	EpisodeNumber int
	EpisodeTitle  string

	Type MediaType
}

type CacheEntry struct {
	HashKey  string
	Metadata Metadata
}
