package util

import (
	"encoding/json"
	"regexp"
	"strings"
)

func ParseMapToJson(mp map[string]string) string {
	str, _ := json.Marshal(mp)
	return string(str)
}

func TrimAllSpacesInString(str string) string {
	return strings.Replace(str, " ", "", -1)
}

func RevomeSpecialChars(str string) string {
	regx := regexp.MustCompile(`[^ A-Za-z0-9]`)
	return regx.ReplaceAllString(str, "")
}

func LetOnlyNumbers(str string) string {
	regx := regexp.MustCompile(`[^ 0-9]`)
	return regx.ReplaceAllString(str, "")
}

// ---------------------------------< validate func >--------------------------------- \\

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}