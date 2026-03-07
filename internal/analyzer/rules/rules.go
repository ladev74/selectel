package rules

import (
	"strings"
	"unicode"
)

func CheckRules(cfg Config, msg string) string {
	msg = strings.TrimSpace(msg)

	if cfg.LowercaseStart.Enabled && !IsLowercaseStart(msg) {
		return StartsWithLowerErrMsg
	}

	if cfg.EnglishOnly.Enabled && !IsEnglishOnly(msg) {
		return EnglishOnlyErrMsg
	}

	if cfg.DisallowSensitiveData.Enabled && !HasNoDisallowSensitiveData(msg, cfg.DisallowSensitiveData.Patterns) {
		return DisallowSensitiveDataErrMsg
	}

	if cfg.DisallowSpecialCharacters.Enabled && !HasNoDisallowSpecialCharacters(msg) {
		return DisallowSpecialCharactersErrMsg
	}

	return ""
}

func IsLowercaseStart(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return unicode.IsLower(r)
		}
	}

	return true
}

func IsEnglishOnly(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			return false
		}
	}

	return true
}

func HasNoDisallowSpecialCharacters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			continue
		}

		return false
	}

	return true
}

func HasNoDisallowSensitiveData(msg string, patterns []string) bool {
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
