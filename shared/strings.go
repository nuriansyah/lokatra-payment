package shared

import "strings"

func SanitizeSpace(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	return str
}

func PaddingLeft(str, paddedTxt string, length int) string {
	var padded string = strings.Repeat(paddedTxt, length-len(str)) + str
	return padded
}
