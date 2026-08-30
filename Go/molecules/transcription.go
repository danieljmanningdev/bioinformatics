package molecules

import "strings"

func Transcribe(t string) string {
	return strings.ReplaceAll(t, "T", "U")
}
