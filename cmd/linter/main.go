package main

import (
	"log"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/ladev74/linter/internal/analyzer"
	"github.com/ladev74/linter/internal/config"
)

func main() {
	cfg, err := config.New("./config/example.yaml")
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	singlechecker.Main(analyzer.New(&cfg.Analyzer))
}
