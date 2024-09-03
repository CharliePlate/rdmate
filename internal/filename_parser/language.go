package filename_parser

import (
	"strings"

	"github.com/dlclark/regexp2"
)

const (
	LanguageArabic     string = "Arabic"
	LanguageBengali    string = "Bengali"
	LanguageBrazilian  string = "Brazilian"
	LanguageBulgarian  string = "Bulgarian"
	LanguageCantonese  string = "Cantonese"
	LanguageCatalan    string = "Catalan"
	LanguageChinese    string = "Chinese"
	LanguageCzech      string = "Czech"
	LanguageDanish     string = "Danish"
	LanguageDutch      string = "Dutch"
	LanguageEnglish    string = "English"
	LanguageEstonian   string = "Estonian"
	LanguageFinnish    string = "Finnish"
	LanguageFlemish    string = "Flemish"
	LanguageFrench     string = "French"
	LanguageGerman     string = "German"
	LanguageGreek      string = "Greek"
	LanguageHebrew     string = "Hebrew"
	LanguageHindi      string = "Hindi"
	LanguageHungarian  string = "Hungarian"
	LanguageIcelandic  string = "Icelandic"
	LanguageItalian    string = "Italian"
	LanguageJapanese   string = "Japanese"
	LanguageKorean     string = "Korean"
	LanguageLatvian    string = "Latvian"
	LanguageLithuanian string = "Lithuanian"
	LanguageMandarin   string = "Mandarin"
	LanguageNorwegian  string = "Norwegian"
	LanguageNordic     string = "Nordic"
	LanguagePolish     string = "Polish"
	LanguagePortuguese string = "Portuguese"
	LanguagePersian    string = "Persian"
	LanguageRomanian   string = "Romanian"
	LanguageRussian    string = "Russian"
	LanguageSerbian    string = "Serbian"
	LanguageSlovak     string = "Slovak"
	LanguageSpanish    string = "Spanish"
	LanguageSwedish    string = "Swedish"
	LanguageTamil      string = "Tamil"
	LanguageThai       string = "Thai"
	LanguageTurkish    string = "Turkish"
	LanguageUkrainian  string = "Ukrainian"
	LanguageVietnamese string = "Vietnamese"
)

var languagePatterns = map[string]*regexp2.Regexp{
	LanguageArabic:     regexp2.MustCompile(`\b(arabic)\b`, 0),
	LanguageBengali:    regexp2.MustCompile(`\b(Bengali)\b`, 0),
	LanguageBrazilian:  regexp2.MustCompile(`\b(Brazilian)\b`, 0),
	LanguageBulgarian:  regexp2.MustCompile(`\b(Bulgarian)\b`, 0),
	LanguageCantonese:  regexp2.MustCompile(`\b(cantonese)\b`, 0),
	LanguageCatalan:    regexp2.MustCompile(`\b(catalan)\b`, 0),
	LanguageChinese:    regexp2.MustCompile(`\b(chi|chinese)\b`, 0),
	LanguageCzech:      regexp2.MustCompile(`\b(czech)\b`, 0),
	LanguageDanish:     regexp2.MustCompile(`\b(DK|DAN|danish)\b`, 0),
	LanguageDutch:      regexp2.MustCompile(`\b(nl|dutch)\b`, 0),
	LanguageEnglish:    regexp2.MustCompile(`\b(english|eng|EN|FI)\b`, 0),
	LanguageEstonian:   regexp2.MustCompile(`\b(estonian)\b`, 0),
	LanguageFinnish:    regexp2.MustCompile(`\b(finnish)\b`, 0),
	LanguageFlemish:    regexp2.MustCompile(`\b(flemish)\b`, 0),
	LanguageFrench:     regexp2.MustCompile(`\b(FR|FRENCH|VOSTFR|VO|VFF|VFQ|VF2|TRUEFRENCH|SUBFRENCH)\b`, 0),
	LanguageGerman:     regexp2.MustCompile(`\b(german)\b`, 0),
	LanguageGreek:      regexp2.MustCompile(`\b(greek)\b`, 0),
	LanguageHebrew:     regexp2.MustCompile(`\b(hebrew)\b`, 0),
	LanguageHindi:      regexp2.MustCompile(`\b(HIN|Hindi)\b`, 0),
	LanguageHungarian:  regexp2.MustCompile(`\b(HUNDUB|HUN|hungarian)\b`, 0),
	LanguageIcelandic:  regexp2.MustCompile(`\b(ice|Icelandic)\b`, 0),
	LanguageItalian:    regexp2.MustCompile(`\b(ita|italian)\b`, 0),
	LanguageJapanese:   regexp2.MustCompile(`\b(japanese)\b`, 0),
	LanguageKorean:     regexp2.MustCompile(`\b(korean)\b`, 0),
	LanguageLatvian:    regexp2.MustCompile(`\b(Latvian)\b`, 0),
	LanguageLithuanian: regexp2.MustCompile(`\b(Lithuanian)\b`, 0),
	LanguageMandarin:   regexp2.MustCompile(`\b(mandarin)\b`, 0),
	LanguageNorwegian:  regexp2.MustCompile(`\b(norwegian|NO)\b`, 0),
	LanguageNordic:     regexp2.MustCompile(`\b(nordic|NORDICSUBS)\b`, 0),
	LanguagePolish:     regexp2.MustCompile(`\b(PL|PLDUB|POLISH)\b`, 0),
	LanguagePortuguese: regexp2.MustCompile(`\b(portuguese)\b`, 0),
	LanguagePersian:    regexp2.MustCompile(`\b(Persian)\b`, 0),
	LanguageRomanian:   regexp2.MustCompile(`\b(RO|Romanian|rodubbed)\b`, 0),
	LanguageRussian:    regexp2.MustCompile(`\b(russian|rus)\b`, 0),
	LanguageSerbian:    regexp2.MustCompile(`\b(Serbian)\b`, 0),
	LanguageSlovak:     regexp2.MustCompile(`\b(SK|Slovak)\b`, 0),
	LanguageSpanish:    regexp2.MustCompile(`\b(spanish)\b`, 0),
	LanguageSwedish:    regexp2.MustCompile(`\b(SE|SWE|swedish)\b`, 0),
	LanguageTamil:      regexp2.MustCompile(`\b(TAM|Tamil)\b`, 0),
	LanguageThai:       regexp2.MustCompile(`\b(thai)\b`, 0),
	LanguageTurkish:    regexp2.MustCompile(`\b(turkish)\b`, 0),
	LanguageUkrainian:  regexp2.MustCompile(`\b(ukrainian|ukr)\b`, 0),
	LanguageVietnamese: regexp2.MustCompile(`\b(vietnamese)\b`, 0),
}

// parseLanguages extracts the languages based on the filename pattern
func parseLanguages(title string) []string {
	titleLower := strings.ToLower(title)
	var detectedLanguages []string

	for lang, pattern := range languagePatterns {
		if match, _ := pattern.MatchString(titleLower); match {
			detectedLanguages = append(detectedLanguages, lang)
		}
	}

	// Default to English if no languages detected
	if len(detectedLanguages) == 0 {
		detectedLanguages = append(detectedLanguages, LanguageEnglish)
	}

	uniqueLanguages := make(map[string]struct{})
	for _, lang := range detectedLanguages {
		uniqueLanguages[lang] = struct{}{}
	}

	var result []string
	for lang := range uniqueLanguages {
		result = append(result, lang)
	}

	return result
}
