package repl

import (
	"bufio"
	"fmt"
	"io"
	"learningLanguage/lexer"
	"learningLanguage/parser"
)

func Start(in io.Reader, out io.Writer) {
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

		if len(parser.Errors()) == 0 {
			fmt.Fprintf(out, "%s\n", program.String())
		} else {
			for _, error := range parser.Errors() {
				fmt.Fprintf(out, "ERROR: %s\n", error)
			}
		}
	}
}
