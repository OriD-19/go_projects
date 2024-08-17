package format

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FormatName(s string) string {
	// Delete the hyphens and capitalize each word
	return cases.Title(language.English, cases.Compact).String(strings.ReplaceAll(s, "-", " "))
}
