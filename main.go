package main

import (
	"fmt"
	"learningLanguage/evaluation"
	"learningLanguage/lexer"
	"learningLanguage/parser"
	"learningLanguage/repl"
	"os"
)

func main() {
	switch os.Args[1] {
	case "lex":
		repl.StartRLPL(os.Stdin, os.Stdout)
	case "parse":
		repl.StartRPPL(os.Stdin, os.Stdout)
	case "eval":
		repl.StartREPL(os.Stdin, os.Stdout)
	}
	// line := "create bool x; set x = true; x == false;"
	// debug(line)
}

func debug(line string) {
	lexer := lexer.New(line)
	parser := parser.New(lexer)
	program := parser.ParseProgram()
	output, errors := evaluation.EvaluateProgram(program)

	if len(parser.Errors()) > 0 {
		for _, error := range parser.Errors() {
			fmt.Fprintf(os.Stdout, "ERROR: %s\n", error)
		}
	} else if len(errors) > 0 {
		for _, error := range errors {
			fmt.Fprintf(os.Stdout, "ERROR: %s\n", error)
		}
	} else {
		fmt.Fprintf(os.Stdout, "%s\n", output)
	}
}
