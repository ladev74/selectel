package rules

type Config struct {
	LowercaseStart            LowercaseStart            `yaml:"lowercase_start"`
	EnglishOnly               EnglishOnly               `yaml:"english_only"`
	DisallowSpecialCharacters DisallowSpecialCharacters `yaml:"disallow_special_characters"`
	DisallowSensitiveData     DisallowSensitiveData     `yaml:"disallow_sensitive_data"`
}

type LowercaseStart struct {
	Enabled bool `yaml:"enabled"`
}
type EnglishOnly struct {
	Enabled bool `yaml:"enabled"`
}

type DisallowSpecialCharacters struct {
	Enabled bool `yaml:"enabled"`
}

type DisallowSensitiveData struct {
	Enabled  bool     `yaml:"enabled"`
	Patterns []string `yaml:"patterns"`
}
