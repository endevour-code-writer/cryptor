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

	switch flag {
	case ToPigLatinFlag:
		fmt.Println(toPigLatin(strToCrypt))
	case EncodeVowelsToIntegers:
		fmt.Println(encodeVowelsToIntegers(strToCrypt))
	case DecodeIntegersToVowels:
		fmt.Println(decodeIntegersToVowels(strToCrypt))
	default:
		fmt.Printf("Please, use one of the flags: %v, %v, %v", ToPigLatinFlag, EncodeVowelsToIntegers, DecodeIntegersToVowels)
	}
}
func isVowel(char string) bool  {
	vowels := Vowels + strings.ToUpper(Vowels)

	return strings.Contains(vowels, char)
}
