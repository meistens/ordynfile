package main

import (
	"flag"
	"fmt"
	"os"

	"ordynfile/internal/config"
	"ordynfile/internal/pipeline"
)

func main() {
	// CLI flags
	path := flag.String("path", ".", "Path to scan for media files")
	dryRun := flag.Bool("dry", false, "Run without making changes")
	confPath := flag.String("config", "config.json", "Path to config file")

	flag.Parse()

	// Load config
	cfg, err := config.Load(*confPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	// Create pipeline runner
	err = pipeline.Run(*path, cfg, *dryRun)
	if err != nil {
		fmt.Println("Pipeline error:", err)
		os.Exit(1)
	}

	fmt.Println("Completed successfully.")
}
