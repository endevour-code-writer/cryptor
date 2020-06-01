package stringCryptor

import "strings"

const PigLatinTail = "ay"

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
