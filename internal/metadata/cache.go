package metadata

import "ordynfile/internal/models"

// TODO: hash(title+year+type)
func GenerateHashKey(file models.MediaFile) string {
	return file.ParsedTitle // placeholder
}

func CacheLookup(key string) (*models.CacheEntry, bool) {
	// TODO: implement
	return nil, false
}

func CacheStore(entry models.CacheEntry) error {
	// TODO: implement
	return nil
}
