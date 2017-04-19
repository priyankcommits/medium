package helpers

import "strings"

func StripWhiteSpaces(whiteSpacedString string) string {
	return strings.Join(strings.Fields(whiteSpacedString), "")
}
