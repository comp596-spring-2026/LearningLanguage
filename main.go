package main

import (
	"fmt"
	"learningLanguage/lexer"
	"learningLanguage/parser"
	"learningLanguage/repl"
	"os"
)

func main() {
	files := os.Args[2]
	print(files)
	switch os.Args[1] {
	case "lex":
		repl.StartRLPL(os.Stdin, os.Stdout)
	case "parse":
		repl.StartRPPL(os.Stdin, os.Stdout)
	case "eval":
		repl.StartREPL(os.Stdin, os.Stdout)
	}
	// line := "set a = 3 * 3 + 2;"
	// debug(line)
}

func debug(line string) {
	out := os.Stdout
	lexer := lexer.New(line)
	parser := parser.New(lexer)
	program := parser.ParseProgram()

	if len(parser.Errors()) == 0 {
		fmt.Fprintf(out, "%s\n", program.String())
	} else {
		for _, error := range parser.Errors() {
			fmt.Fprintf(out, "ERROR: %s\n", error)
		}
	}
}
