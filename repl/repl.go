package repl

// READ EVALUATE PRINT LOOP

import (
	"bufio"
	"fmt"
	"io"
	"learningLanguage/evaluation"
	"learningLanguage/lexer"
	"learningLanguage/parser"
	"learningLanguage/token"
)

func StartREPL(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, ">> ")
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
				fmt.Fprintf(out, "ERROR: %s\n", error)
			}
		} else if len(errors) > 0 {
			for _, error := range errors {
				fmt.Fprintf(out, "ERROR: %s\n", error)
			}
		} else {
			fmt.Fprintf(out, "%s\n", output)
		}
	}
}

func StartRLPL(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, ">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)
		tok := lexer.NextToken()

		for tok.Type != token.EOF {
			fmt.Fprintf(out, "%+v\n", tok)
			tok = lexer.NextToken()
		}
	}
}

func StartRPPL(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, ">> ")
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
				fmt.Fprintf(out, "ERROR: %s\n", error)
			}
		} else {
			fmt.Fprintf(out, "%s\n", program.String())
		}
	}
}
