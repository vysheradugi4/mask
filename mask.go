package mask

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// CreateMask function get current code (string) and keyword (string), then
// create mask from this information like \d\dKey where "Key" is keyword
// in necessary register case
func CreateMask(code string, key string) (string, error) {
	// \d			digit char
	// [a-z]		lowercase character (NOT SUPPORTED)
	// [A-Z]		uppercase character (NOT SUPPORTED)
	// key			keyword in lowercase
	// Key			keyword with first character in uppercase
	// KEY			keyword with all characters in uppercase
	// kEy			keyword with camelcase
	//
	// example: code WINTER10 in mask KEY\d\d

	// Check if code contain keyword

	var mask string // for ready mask
	var err error   // just in case
	lencode := len(code)
	lenkey := len(key)

	switch {

	// Length of code is null
	case lencode == 0:
		return "", errors.New("Code string should be not empty")

	// Length of code less key
	case lencode < lenkey:
		return "", errors.New("Code string should be not shorter than the key")

	case !containsFold(code, key):
		return "", errors.New("Code must contain the key")

	case lencode == lenkey:
		mask, err = changeToKey(code, key)
		if err != nil {
			return "", err
		}

		return mask, nil

	case lencode > lenkey:
		mask, err = changeToKey(code, key)
		if err != nil {
			return "", err
		}

		mask = changeChar(mask)
	}
	return mask, nil
}

// changeToKey is function find key in code and change it to keyword "key"
// in necessary register case.
func changeToKey(code, key string) (string, error) {

	if !containsFold(code, key) {
		return "", errors.New("Code must contain key")
	}

	// Find key into code with necessary case
	reg := regexp.MustCompile(`(?i)` + key)
	key = reg.FindString(code)

	// For flag value lowercase by default, valid values of flag: lowercase,
	// UPPERCASE, camelCase, Capitalize
	const lowercase string = "lowercase"
	const UPPERCASE string = "UPPERCASE"
	const Capitalize string = "Capitalize"
	const camelCase string = "camelCase"
	flag := lowercase

	// Define flag
	for index, char := range key {
		switch {
		case unicode.IsUpper(char) && index == 0:
			flag = Capitalize

		case unicode.IsUpper(char) && index == 1 && flag == Capitalize:
			flag = UPPERCASE

		case unicode.IsUpper(char) && index > 1 && flag == lowercase:
			flag = camelCase

		case unicode.IsUpper(char) && index > 1 && flag == UPPERCASE:
			flag = UPPERCASE

		case unicode.IsLower(char) && index > 1 && flag == UPPERCASE:
			flag = camelCase

		case unicode.IsUpper(char) && index > 1 && flag == Capitalize:
			flag = camelCase
		}
	}

	var mask string

	// Change keyword
	switch {
	case flag == UPPERCASE:
		mask = strings.Replace(code, key, "KEY", -1)

	case flag == Capitalize:
		mask = strings.Replace(code, key, "Key", -1)

	case flag == camelCase:
		mask = strings.Replace(code, key, "kEy", -1)

	default:
		mask = strings.Replace(code, key, "key", -1)
	}
	return mask, nil
}

// changeChar is function for substitution all numbers in string
// to special expression "\d" and all characters to [a-z] (case sensitive)
func changeChar(s string) (out string) {
	flag := 3 // counter for "key" word

	for index, char := range s {

		// Skipping "key" word
		if flag < 3 && flag > 0 {
			out += string(char)
			flag--
			continue
		}
		if strings.ToLower(string(char)) == "k" {
			if strings.ToLower(string([]rune(s)[index+1])) == "e" {
				if strings.ToLower(string([]rune(s)[index+2])) == "y" {
					out += string(char)
					flag--
					continue
				}
			}
		}

		// Replace digit characters to \d
		if unicode.IsDigit(char) {
			out += "\\d"
			continue
		}

		// Replace characters to [a-z] or [A-Z] depends on case of char
		if unicode.IsLower(char) {
			out += "[a-z]"
		} else {
			out += "[A-Z]"
		}
	}
	return out
}

// containsFold checks the first string contains the second string
func containsFold(s, substr string) bool {
	if s == "" || substr == "" {
		return false
	}
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)

	return strings.Contains(s, substr)
}

// GenerateCodesFromMask is super generator of codes with use base mask and
// keyword. Iterations is max count of characters for substitution
func GenerateCodesFromMask(mask, keyword string, iterations int) []string {

	// Check how many characters contain mask for substitution
	reg, err := regexp.Compile("\\[A-Z\\]|\\[a-z\\]|\\\\d")
	if err != nil {
		return nil
	}

	if len(reg.FindAllStringIndex(mask, -1)) > iterations {
		return nil
	}

	// var code0, code1, code2 string
	var code0 string
	var result []string

	// Replace key to keyword
	switch {
	case strings.Contains(mask, "key"):
		code0 = strings.Replace(mask, "key", strings.ToLower(keyword), -1)
	case strings.Contains(mask, "Key"):
		code0 = strings.Replace(mask, "Key", strings.Title(keyword), -1)
	case strings.Contains(mask, "KEY"):
		code0 = strings.Replace(mask, "KEY", strings.ToUpper(keyword), -1)
	}

	substitution(code0, &result)

	return result
}

// substitution is a recursive function for change characters
func substitution(code string, result *[]string) {

	for i := 0; i < 10; i++ {
		if !strings.Contains(code, "\\d") {

			*result = append(*result, code)
			break
		}

		substitution(strings.Replace(code, "\\d", strconv.Itoa(i), 1), result)
	}
}
