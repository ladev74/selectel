package rules

import (
	"strings"
	"unicode"
)

func CheckRules(cfg Config, msg string) string {
	msg = strings.TrimSpace(msg)

	if cfg.LowercaseStart.Enabled && !lowercaseStart(msg) {
		return StartsWithLowerErrMsg
	}

	if cfg.EnglishOnly.Enabled && !englishOnly(msg) {
		return EnglishOnlyErrMsg
	}

	if cfg.DisallowSensitiveData.Enabled && !disallowSensitiveData(msg, cfg.DisallowSensitiveData.Patterns) {
		return DisallowSensitiveDataErrMsg
	}

	if cfg.DisallowSpecialCharacters.Enabled && !disallowSpecialCharacters(msg) {
		return DisallowSpecialCharactersErrMsg
	}

	return ""
}

func lowercaseStart(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return unicode.IsLower(r)
		}
	}

	return true
}

func englishOnly(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			return false
		}
	}

	return true
}

func disallowSpecialCharacters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			continue
		}

		return false
	}

	return true
}

func disallowSensitiveData(msg string, patterns []string) bool {
	words := splitToWords(msg)

	for _, word := range words {
		for _, pattern := range patterns {
			lowerPattern := strings.ToLower(pattern)
			lowerWord := strings.ToLower(word)

			if lowerWord == lowerPattern {
				return false
			}
		}
	}

	return true
}

func splitToWords(s string) []string {
	var words []string
	var current []rune

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			current = append(current, unicode.ToLower(r))
		} else {
			if len(current) > 0 {
				words = append(words, string(current))
				current = []rune{}
			}
		}
	}

	if len(current) > 0 {
		words = append(words, string(current))
	}

	return words
}
