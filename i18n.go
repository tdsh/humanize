package humanize

import (
	"fmt"
	"os"

	"github.com/nicksnyder/go-i18n/i18n"
)

var trans i18n.TranslateFunc

// SetLocale sets l to current locale. l must be specified in
// "lang code (ISO 639)"-"country code (ISO 3166-1)". (en-us, ja-jp,
// for example.) Available translations are located at translations
// directory as JSON file. If unsupported language is selected,
// English (en-us) is used.
func SetLocale(l string) {
	tFile := fmt.Sprintf("translations/%s.all.json", l)
	_, err := os.Stat(tFile)
	if err == nil {
		i18n.MustLoadTranslationFile(tFile)
	} else {
		i18n.MustLoadTranslationFile("translations/en-us.all.json")
	}
	trans, _ = i18n.Tfunc(l, "en-us")
}

// getLocale initializes locale. It must be run before using translation.
func getLocale() {
	if trans == nil {
		i18n.MustLoadTranslationFile("translations/en-us.all.json")
		trans, _ = i18n.Tfunc("en-us")
	}
}
