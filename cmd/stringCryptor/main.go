package main

import (
	"github.com/endevour-code-writer/cryptor/internal/stringCryptor"
	"os"
)

func main() {
	args := os.Args[1:]
	stringCryptor.Crypt(args)
}
