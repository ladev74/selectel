package disallow_special_characters

import "log/slog"

func disallowSpecialCharacters() {
	logger := &slog.Logger{}

	logger.Info("server started!🚀")   // want "^log message must not contain special characters or emojis$"
	logger.Info("starting server!!!") // want "^log message must not contain special characters or emojis$"
	logger.Info("starting server...") // want "^log message must not contain special characters or emojis$"
	logger.Info("starting server")    // ok

	slog.Info("server started!🚀")   // want "^log message must not contain special characters or emojis$"
	slog.Info("starting server!!!") // want "^log message must not contain special characters or emojis$"
	slog.Info("starting server...") // want "^log message must not contain special characters or emojis$"
	slog.Info("starting server")    // ok
}
