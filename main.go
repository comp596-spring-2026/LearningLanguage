package main

import (
	"learningLanguage/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
