package disallow_sensitive_data

import "log/slog"

func disallowSensitiveData() {
	logger := &slog.Logger{}

	logger.Info("my pass")    // want "^log message must not contain sensitive data$"
	logger.Info("my api_key") // want "^log message must not contain sensitive data$"

	slog.Info("my pass")    // want "^log message must not contain sensitive data$"
	slog.Info("my api_key") // want "^log message must not contain sensitive data$"
}
