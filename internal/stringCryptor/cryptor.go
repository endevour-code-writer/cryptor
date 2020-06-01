package stringCryptor

import (
	"fmt"
	"os"
	"strings"
)

const (
	Vowels = "aeiou"
	ToPigLatinFlag = "toPigLatin"
	EncodeVowelsToIntegers = "encodeVowelsToIntegers"
	DecodeIntegersToVowels = "decodeIntegersToVowels"
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
	transformBasedOnFlag(flag, strToCrypt)
}

func transformBasedOnFlag(flag string, translateTarget string) {
	var transformed string

	switch flag {
	case ToPigLatinFlag:
		transformed = toPigLatin(translateTarget)
	case EncodeVowelsToIntegers:
		transformed = encodeVowelsToIntegers(translateTarget)
	case DecodeIntegersToVowels:
		transformed = decodeIntegersToVowels(translateTarget)
	default:
		transformed = "Please, use one of the flags: " +  ToPigLatinFlag + "," + EncodeVowelsToIntegers + "," + DecodeIntegersToVowels
	}

	fmt.Println(transformed)
}

func isVowel(char string) bool  {
	vowels := Vowels + strings.ToUpper(Vowels)

	return strings.Contains(vowels, char)
}
