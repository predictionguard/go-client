package client

import "fmt"

// Set of known languages.
var languages = make(map[string]Language)

// Set of possible languages.
var (
	LangEnglish = newLanguage("english")
	LanSpanish  = newLanguage("spanish")
)

// Language represents a language in the system.
type Language struct {
	name string
}

func newLanguage(language string) Language {
	l := Language{language}
	languages[language] = l
	return l
}

// ParseLanguage parses the string value and returns a language if one exists.
func ParseLanguage(value string) (Language, error) {
	lang, exists := languages[value]
	if !exists {
		return Language{}, fmt.Errorf("invalid language %q", value)
	}

	return lang, nil
}

// MustParseLanguage parses the string value and returns a language if one
// exists. If an error occurs the function panics.
func MustParseLanguage(value string) Language {
	lang, err := ParseLanguage(value)
	if err != nil {
		panic(err)
	}

	return lang
}

// Name returns the name of the role.
func (l Language) Name() string {
	return l.name
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (l *Language) UnmarshalText(data []byte) error {
	lang, err := ParseLanguage(string(data))
	if err != nil {
		return err
	}

	l.name = lang.name
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (l Language) MarshalText() ([]byte, error) {
	return []byte(l.name), nil
}

// Equal provides support for the go-cmp package and testing.
func (l Language) Equal(l2 Language) bool {
	return l.name == l2.name
}
