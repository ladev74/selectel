package main

import (
	"log"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/ladev74/linter/internal/analyzer"
	"github.com/ladev74/linter/internal/config"
)

// TODO: в README написать почему не вынес в конфиг поддерживаемые логгеры (в тз жестко прописаны какие логгеры нужно поддерживать)
// TODO: указать почему в тестах нет zap

func main() {
	cfg, err := config.New("/home/ladev/projects/linter/config/example.yaml")
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	singlechecker.Main(analyzer.New(&cfg.Analyzer))
}
