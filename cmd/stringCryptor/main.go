package main

import (
	"fmt"
	"github.com/endevour-code-writer/cryptor/internal/stringCryptor"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No String")
	} else {
		str := strings.Join(args, " ")
		fmt.Println(stringCryptor.ToPigLatin(str))
	}
}
