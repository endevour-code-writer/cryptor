package stringCryptor

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

const (
	PigLatinTail = "ay"
	Vowels = "aeiou"
)

func Crypt(args []string) {
	if 0 == len(args) {
		fmt.Println("No String to crypt")
		os.Exit(0)
	}

	flag := args[:1][0]
	sourceToCrypt := args[1:]

	if 0 == len(sourceToCrypt) {
		fmt.Println("No String to crypt")
		os.Exit(0)
	}

	strToCrypt := strings.Join(args[1:], " ")

	switch flag {
	case "toPigLatin":
		fmt.Println(toPigLatin(strToCrypt))
	case "encodeVowels":
		fmt.Println(encodeVowels(strToCrypt))
	default:
		fmt.Println("Too far away.")
	}
}

func toPigLatin(str string) string  {
	var encoded []string
	words := strings.Split(str, " ")

	for _, word := range words {
		pigLatinated := wordToPigLatin(word)
		encoded = append(encoded, pigLatinated)
	}

	return strings.Join(encoded, " ")

}

func wordToPigLatin(word string) string {
	vowels := Vowels + strings.ToUpper(Vowels)
	firstLetter := word[0:1]

	if strings.Contains(vowels, firstLetter) {
		return word + PigLatinTail
	}

	for i, v := range word {
		if isVowel(string(v)) {
			word = word[i:] + word[:i] + PigLatinTail
			return word
		}
	}

	return word
}

func encodeVowels(str string) string {
	for i, char := range str {
		toString := strings.ToLower(string(char))

		if isVowel(toString) {
			index := strconv.Itoa(strings.Index(Vowels, toString) + 1)
			str = str[:i] + index + str[i+1:]
		}
	}

	return str
}

func isVowel(char string) bool  {
	vowels := Vowels + strings.ToUpper(Vowels)

	return strings.Contains(vowels, char)
}
