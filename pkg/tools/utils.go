package tools

import (
	"strings"
)

func RemoveSpaces(input string) string {
	return strings.ReplaceAll(input, " ", "")
}
