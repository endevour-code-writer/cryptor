package stringCryptor

import "strings"

const (
	PigLatinTail = "ay"
	Vowels = "aeiou"
)

func ToPigLatin(str string) string  {
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
		if isVowel(v) {
			word = word[i:] + word[:i] + PigLatinTail
			return word
		}
	}

	return word
}


func isVowel(char rune) bool  {
	vowels := Vowels + strings.ToUpper(Vowels)

	return strings.Contains(vowels, string(char))
}
