package repl

// READ EVALUATE PRINT LOOP

import (
	"bufio"
	"fmt"
	"learningLanguage/evaluation"
	"learningLanguage/lexer"
	"learningLanguage/parser"
	"learningLanguage/token"
	"os"
)

func StartREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)
		parser := parser.New(lexer)
		program := parser.ParseProgram()
		output, errors := evaluation.EvaluateProgram(program)

		if len(parser.Errors()) > 0 {
			for _, error := range parser.Errors() {
				fmt.Printf("ERROR: %s\n", error)
			}
		} else if len(errors) > 0 {
			for _, error := range errors {
				fmt.Printf("ERROR: %s\n", error)
			}
		} else {
			fmt.Printf("%s\n", output)
		}
	}
}

func StartRLPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)
		tok := lexer.NextToken()

		for tok.Type != token.EOF {
			fmt.Printf("%+v\n", tok)
			tok = lexer.NextToken()
		}
	}
}

func StartRPPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)
		parser := parser.New(lexer)
		program := parser.ParseProgram()

		if len(parser.Errors()) > 0 {
			for _, error := range parser.Errors() {
				fmt.Printf("ERROR: %s\n", error)
			}
		} else {
			fmt.Printf("%s\n", program.String())
		}
	}
}
