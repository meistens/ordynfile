package pipeline

import (
	"ordynfile/internal/config"
	"ordynfile/internal/metadata"
	"ordynfile/internal/organizer"
	"ordynfile/internal/parser"
	"ordynfile/internal/scan"
	"ordynfile/internal/utils"
)

func Run(path string, cfg *config.Config, dry bool) error {
	// Step 1: Scan
	files, err := scan.Scan(path)
	if err != nil {
		return err
	}
	utils.Info("Found %d files", len(files))

	for _, f := range files {
		// Step 2: Parse
		parser.Parse(&f)

		// Step 3: Metadata
		md, err := metadata.Resolve(f)
		if err != nil {
			utils.Error("Metadata failed: %v", err)
			continue
		}

		// Step 4: Renaming
		newName := organizer.BuildFilename(md, f.Extension)

		// Step 5: Move
		destPath := cfg.LibraryRoot + "/" + newName
		err = organizer.MoveFile(f.FullPath, destPath, dry)
		if err != nil {
			utils.Error("Move failed: %v", err)
		}
	}

	return nil
}
