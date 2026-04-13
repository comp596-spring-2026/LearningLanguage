package main

import (
	"fmt"
	"io"
	"learningLanguage/evaluation"
	"learningLanguage/lexer"
	"learningLanguage/parser"
	"learningLanguage/repl"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		repl.StartREPL()
	} else {
		inFile := os.Stdin
		outFile := os.Stdout
		index := 1
		for index < len(os.Args) {
			switch os.Args[index] {
			case "-i":
				index++
				inFile, _ = os.Open(os.Args[index])
			case "-o":
				index++
				outFile, _ = os.OpenFile(os.Args[index], os.O_WRONLY|os.O_CREATE, 0600)
			default:
				index++
			}
		}
		executeProgram(inFile, outFile)
		inFile.Close()
		outFile.Close()
	}
	// line := "if (1 != 1) begin; 123; end; else begin; 321; end;"
	// debug(line)
}

func executeProgram(in io.Reader, out io.Writer) {
	text, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}

	lexer := lexer.New(string(text))
	parser := parser.New(lexer)
	program := parser.ParseProgram()
	output, errors := evaluation.EvaluateProgram(program)

	if len(errors) == 0 {
		fmt.Fprint(out, output)
		fmt.Println(output)
	} else {
		for _, err := range errors {
			fmt.Fprint(out, err)
			fmt.Println(err)
		}
	}
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
