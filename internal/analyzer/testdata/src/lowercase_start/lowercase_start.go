package lowercase_start

import (
	"log/slog"
)

func lowercaseStart() {
	logger := &slog.Logger{}

	logger.Info("Starting server") // want "^log message must start with lowercase letter$"

	logger.Info("starting server") // ok
	logger.Info("sTarting server") // ok

	slog.Info("Starting server") // want "^log message must start with lowercase letter$"

	slog.Info("starting server") // ok
	slog.Info("sTarting server") // ok
}
