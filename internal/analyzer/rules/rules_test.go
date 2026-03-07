package rules

import "testing"

func TestCheckRules(t *testing.T) {
	cfg := Config{
		LowercaseStart:            LowercaseStart{true},
		EnglishOnly:               EnglishOnly{true},
		DisallowSpecialCharacters: DisallowSpecialCharacters{true},
		DisallowSensitiveData:     DisallowSensitiveData{true, []string{"pass", "api_key"}},
	}

	tests := []struct {
		name string
		msg  string
		want string
	}{
		{
			name: "valid",
			msg:  "hello world",
			want: "",
		},
		{
			name: "lowercase start",
			msg:  "Hello world",
			want: StartsWithLowerErrMsg,
		},
		{
			name: "english only",
			msg:  "привет мир",
			want: EnglishOnlyErrMsg,
		},
		{
			name: "disallow special characters",
			msg:  "hello world!",
			want: DisallowSpecialCharactersErrMsg,
		},
		{
			name: "disallow sensitive data",
			msg:  "pass",
			want: DisallowSensitiveDataErrMsg,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckRules(cfg, tt.msg)
			if got != tt.want {
				t.Errorf("CheckRules() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowercaseStart(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{
			name: "first letter lowercase",
			msg:  "hello world",
			want: true,
		},
		{
			name: "first letter uppercase",
			msg:  "Hello world",
			want: false,
		},
		{
			name: "second letter uppercase",
			msg:  "hEllo world",
			want: true,
		},
		{
			name: "last letter uppercase",
			msg:  "hello worlD",
			want: true,
		},
		{
			name: "symbols first second uppercase",
			msg:  "!Abc",
			want: false,
		},
		{
			name: "symbols first third uppercase",
			msg:  "!aBc",
			want: true,
		},
		{
			name: "only special characters",
			msg:  "!",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLowercaseStart(tt.msg)
			if got != tt.want {
				t.Errorf("lowercaseStart(%s) got: %v, want: %v", tt.msg, got, tt.want)
			}
		})
	}
}

func TestEnglishOnly(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{
			name: "english only",
			msg:  "hello world",
			want: true,
		},
		{
			name: "not english",
			msg:  "привет мир",
			want: false,
		},
		{
			name: "first letter not english",
			msg:  "хello world",
			want: false,
		},
		{
			name: "only first letter english",
			msg:  "gривет мир",
			want: false,
		},
		{
			name: "second letter not english",
			msg:  "hыllo world",
			want: false,
		},
		{
			name: "only second letter english",
			msg:  "прiвет мир",
			want: false,
		},
		{
			name: "last letter not english",
			msg:  "hello worlд",
			want: false,
		},
		{
			name: "only last letter english",
			msg:  "привет миr",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEnglishOnly(tt.msg)
			if got != tt.want {
				t.Errorf("englishOnly(%s) got: %v, want: %v", tt.msg, got, tt.want)
			}
		})
	}
}

func TestDisallowSpecialCharacters(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{
			name: "not special characters",
			msg:  "hello",
			want: true,
		},
		{
			name: "special characters !",
			msg:  "hello!",
			want: false,
		},
		{
			name: "special characters .",
			msg:  "hello.",
			want: false,
		},
		{
			name: "special characters ,",
			msg:  "hello,",
			want: false,
		},
		{
			name: "special characters <",
			msg:  "hello<",
			want: false,
		},
		{
			name: "special characters :",
			msg:  "hello:",
			want: false,
		},
		{
			name: "emoji",
			msg:  "🦆",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasNoDisallowSpecialCharacters(tt.msg)
			if got != tt.want {
				t.Errorf("disallowSpecialCharacters(%s) got: %v, want: %v", tt.msg, got, tt.want)
			}
		})
	}
}

func TestDisallowSensitiveData(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		patters []string
		want    bool
	}{
		{
			name:    "sensitive data: password",
			msg:     "my password",
			patters: []string{"password"},
			want:    false,
		},
		{
			name:    "sensitive data: password with :",
			msg:     "my password:",
			patters: []string{"password"},
			want:    false,
		},
		{
			name:    "sensitive data: pass",
			msg:     "my password",
			patters: []string{"pass"},
			want:    true,
		},
		{
			name:    "sensitive data: api_key",
			msg:     "api_key",
			patters: []string{"api_key"},
			want:    false,
		},
		{
			name:    "sensitive data: api_key",
			msg:     "apikey",
			patters: []string{"api_key"},
			want:    true,
		},
		{
			name:    "sensitive data: api_key and pass",
			msg:     "api_key and pass",
			patters: []string{"api_key", "pass"},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasNoDisallowSensitiveData(tt.msg, tt.patters)
			if got != tt.want {
				t.Errorf("disallowSensitiveData(%s) got: %v, want: %v, patters: %v", tt.msg, got, tt.want, tt.patters)
			}
		})
	}
}
