package client

import "fmt"

// Set of known languages.
var languages = make(map[string]Language)

// Set of possible languages.
var (
	LangAfrikanns   = newLanguage("afr")
	LangAmharic     = newLanguage("amh")
	LangArabic      = newLanguage("ara")
	LangArmenian    = newLanguage("hye")
	LangAzerbaijani = newLanguage("aze")
	LangBasque      = newLanguage("eus")
	LangBelarusian  = newLanguage("bel")
	LangBengali     = newLanguage("ben")
	LangBosnian     = newLanguage("bos")
	LangCatalan     = newLanguage("cat")
	LangChechen     = newLanguage("che")
	LangCherokee    = newLanguage("chr")
	LangChinese     = newLanguage("zho")
	LangCroatian    = newLanguage("hrv")
	LangCzech       = newLanguage("ces")
	LangDanish      = newLanguage("dan")
	LangDutch       = newLanguage("nld")
	LangEnglish     = newLanguage("eng")
	LangEstonian    = newLanguage("est")
	LangFijian      = newLanguage("fij")
	LangFilipino    = newLanguage("fil")
	LangFinnish     = newLanguage("fin")
	LangFrench      = newLanguage("fra")
	LangGalician    = newLanguage("glg")
	LangGeorgian    = newLanguage("kat")
	LangGerman      = newLanguage("deu")
	LangGreek       = newLanguage("ell")
	LangGujarati    = newLanguage("guj")
	LangHaitian     = newLanguage("hat")
	LangHebrew      = newLanguage("heb")
	LangHindi       = newLanguage("hin")
	LangHungarian   = newLanguage("hun")
	LangIcelandic   = newLanguage("isl")
	LangIndonesian  = newLanguage("ind")
	LangIrish       = newLanguage("gle")
	LangItalian     = newLanguage("ita")
	LangJapanese    = newLanguage("jpn")
	LangKannada     = newLanguage("kan")
	LangKazakh      = newLanguage("kaz")
	LangKorean      = newLanguage("kor")
	LangLatvian     = newLanguage("lav")
	LangLithuanian  = newLanguage("lit")
	LangMacedonian  = newLanguage("mkd")
	LangMalay1      = newLanguage("msa")
	LangMalay2      = newLanguage("zlm")
	LangMalayalam   = newLanguage("mal")
	LangMaltese     = newLanguage("mlt")
	LangMarathi     = newLanguage("mar")
	LangNepali      = newLanguage("nep")
	LangNorwegian   = newLanguage("nor")
	LangPersian     = newLanguage("fas")
	LangPolish      = newLanguage("pol")
	LangPortuguese  = newLanguage("por")
	LangRomanian    = newLanguage("ron")
	LangRussian     = newLanguage("rus")
	LangSamoan      = newLanguage("smo")
	LangSerbian     = newLanguage("srp")
	LangSlovak      = newLanguage("slk")
	LangSlovenian   = newLanguage("slv")
	LangSlavonic    = newLanguage("chu")
	LangSpanish     = newLanguage("spa")
	LangSwahili     = newLanguage("swh")
	LangSwedish     = newLanguage("swe")
	LangTamil       = newLanguage("tam")
	LangTelugu      = newLanguage("tel")
	LangThai        = newLanguage("tha")
	LangTurkish     = newLanguage("tur")
	LangUkrainian   = newLanguage("ukr")
	LangUrdu        = newLanguage("urd")
	LangWelsh       = newLanguage("cym")
	LangVietnamese  = newLanguage("vie")
)

// Language represents a language in the system.
type Language struct {
	code string
}

func newLanguage(code string) Language {
	l := Language{code}
	languages[code] = l
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

// Code returns the ISO-639 code of the language.
func (l Language) Code() string {
	return l.code
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (l *Language) UnmarshalText(data []byte) error {
	lang, err := ParseLanguage(string(data))
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
