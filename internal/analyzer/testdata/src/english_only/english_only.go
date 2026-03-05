package english_only

import (
	"log/slog"
)

func englishOnly() {
	logger := &slog.Logger{}

	logger.Info("запуск сервера") // want "^log message must be in English$"
	logger.Info("servыr started") // want "^log message must be in English$"
	logger.Info("server started") // ok

	slog.Info("запуск сервера") // want "^log message must be in English$"
	slog.Info("servыr started") // want "^log message must be in English$"
	slog.Info("server started") // ok
}
