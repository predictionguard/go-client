package client

import "fmt"

type languageSet struct {
	Afrikanns  Language
	Amharic    Language
	Arabic     Language
	Armenian   Language
	Azerbaijan Language
	Basque     Language
	Belarusian Language
	Bengali    Language
	Bosnian    Language
	Catalan    Language
	Chechen    Language
	Cherokee   Language
	Chinese    Language
	Croatian   Language
	Czech      Language
	Danish     Language
	Dutch      Language
	English    Language
	Estonian   Language
	Fijian     Language
	Filipino   Language
	Finnish    Language
	French     Language
	Galician   Language
	Georgian   Language
	German     Language
	Greek      Language
	Gujarati   Language
	Haitian    Language
	Hebrew     Language
	Hindi      Language
	Hungarian  Language
	Icelandic  Language
	Indonesian Language
	Irish      Language
	Italian    Language
	Japanese   Language
	Kannada    Language
	Kazakh     Language
	Korean     Language
	Latvian    Language
	Lithuanian Language
	Macedonian Language
	Malay1     Language
	Malay2     Language
	Malayalam  Language
	Maltese    Language
	Marathi    Language
	Nepali     Language
	Norwegian  Language
	Persian    Language
	Polish     Language
	Portuguese Language
	Romanian   Language
	Russian    Language
	Samoan     Language
	Serbian    Language
	Slovak     Language
	Slovenian  Language
	Slavonic   Language
	Spanish    Language
	Swahili    Language
	Swedish    Language
	Tamil      Language
	Telugu     Language
	Thai       Language
	Turkish    Language
	Ukrainian  Language
	Urdu       Language
	Welsh      Language
	Vietnamese Language
}

// Languages represents the set of languages that can be used.
var Languages = languageSet{
	Afrikanns:  newLanguage("afr"),
	Amharic:    newLanguage("amh"),
	Arabic:     newLanguage("ara"),
	Armenian:   newLanguage("hye"),
	Azerbaijan: newLanguage("aze"),
	Basque:     newLanguage("eus"),
	Belarusian: newLanguage("bel"),
	Bengali:    newLanguage("ben"),
	Bosnian:    newLanguage("bos"),
	Catalan:    newLanguage("cat"),
	Chechen:    newLanguage("che"),
	Cherokee:   newLanguage("chr"),
	Chinese:    newLanguage("zho"),
	Croatian:   newLanguage("hrv"),
	Czech:      newLanguage("ces"),
	Danish:     newLanguage("dan"),
	Dutch:      newLanguage("nld"),
	English:    newLanguage("eng"),
	Estonian:   newLanguage("est"),
	Fijian:     newLanguage("fij"),
	Filipino:   newLanguage("fil"),
	Finnish:    newLanguage("fin"),
	French:     newLanguage("fra"),
	Galician:   newLanguage("glg"),
	Georgian:   newLanguage("kat"),
	German:     newLanguage("deu"),
	Greek:      newLanguage("ell"),
	Gujarati:   newLanguage("guj"),
	Haitian:    newLanguage("hat"),
	Hebrew:     newLanguage("heb"),
	Hindi:      newLanguage("hin"),
	Hungarian:  newLanguage("hun"),
	Icelandic:  newLanguage("isl"),
	Indonesian: newLanguage("ind"),
	Irish:      newLanguage("gle"),
	Italian:    newLanguage("ita"),
	Japanese:   newLanguage("jpn"),
	Kannada:    newLanguage("kan"),
	Kazakh:     newLanguage("kaz"),
	Korean:     newLanguage("kor"),
	Latvian:    newLanguage("lav"),
	Lithuanian: newLanguage("lit"),
	Macedonian: newLanguage("mkd"),
	Malay1:     newLanguage("msa"),
	Malay2:     newLanguage("zlm"),
	Malayalam:  newLanguage("mal"),
	Maltese:    newLanguage("mlt"),
	Marathi:    newLanguage("mar"),
	Nepali:     newLanguage("nep"),
	Norwegian:  newLanguage("nor"),
	Persian:    newLanguage("fas"),
	Polish:     newLanguage("pol"),
	Portuguese: newLanguage("por"),
	Romanian:   newLanguage("ron"),
	Russian:    newLanguage("rus"),
	Samoan:     newLanguage("smo"),
	Serbian:    newLanguage("srp"),
	Slovak:     newLanguage("slk"),
	Slovenian:  newLanguage("slv"),
	Slavonic:   newLanguage("chu"),
	Spanish:    newLanguage("spa"),
	Swahili:    newLanguage("swh"),
	Swedish:    newLanguage("swe"),
	Tamil:      newLanguage("tam"),
	Telugu:     newLanguage("tel"),
	Thai:       newLanguage("tha"),
	Turkish:    newLanguage("tur"),
	Ukrainian:  newLanguage("ukr"),
	Urdu:       newLanguage("urd"),
	Welsh:      newLanguage("cym"),
	Vietnamese: newLanguage("vie"),
}

// Parse parses the string value and returns a language if one exists.
func (languageSet) Parse(value string) (Language, error) {
	lang, exists := languages[value]
	if !exists {
		return Language{}, fmt.Errorf("invalid language %q", value)
	}

	return lang, nil
}

// MustParse parses the string value and returns a language if one
// exists. If an error occurs the function panics.
func (languageSet) MustParse(value string) Language {
	lang, err := Languages.Parse(value)
	if err != nil {
		panic(err)
	}

	return lang
}

// =============================================================================

// Set of known languages.
var languages = make(map[string]Language)

// Language represents a language in the system.
type Language struct {
	code string
}

func newLanguage(code string) Language {
	l := Language{code}
	languages[code] = l
	return l
}

// String returns the ISO-639 code of the language.
func (l Language) String() string {
	return l.code
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (l *Language) UnmarshalText(data []byte) error {
	lang, err := Languages.Parse(string(data))
	if err != nil {
		return err
	}

	l.code = lang.code
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (l Language) MarshalText() ([]byte, error) {
	return []byte(l.code), nil
}

// Equal provides support for the go-cmp package and testing.
func (l Language) Equal(l2 Language) bool {
	return l.code == l2.code
}
