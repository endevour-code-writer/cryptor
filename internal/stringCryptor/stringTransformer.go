package stringCryptor

import (
	"strconv"
	"strings"
)

func encodeVowelsToIntegers(str string) string {
	for i, char := range str {
		toString := strings.ToLower(string(char))

		if isVowel(toString) {
			index := strconv.Itoa(strings.Index(Vowels, toString) + 1)
			str = str[:i] + index + str[i+1:]
		}
	}

	return str
}

func decodeIntegersToVowels(str string) string {
	for i, char := range str {
		if charToInt, err := strconv.Atoi(string(char)); err == nil {
			if charToInt != 0  && len(Vowels) >= charToInt {
				vowel := string(Vowels[charToInt-1])
				str = str[:i] + vowel + str[i+1:]
			}
		}
	}

	return str
}
