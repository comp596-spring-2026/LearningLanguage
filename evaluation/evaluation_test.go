package evaluation

import (
	"learningLanguage/lexer"
	"learningLanguage/parser"
	"strings"
	"testing"
)

func TestCreateSetEval(test *testing.T) {
	input := `create int x;set x = 64;x;`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	output, errors := EvaluateProgram(program)
	if len(errors) != 0 {
		test.Fatalf("Errors were found: %v", errors)
	}

	output = strings.TrimSpace(output)

	if output != "64" {
		test.Fatalf("Incorrect variable value, expected 64, got %s", output)
	}
}

func TestIfEval(test *testing.T) {
	input := `if (1 > 2) begin; true; end; else begin; false; end;`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	output, errors := EvaluateProgram(program)
	if len(errors) != 0 {
		test.Fatalf("Errors were found: %v", errors)
	}

	output = strings.TrimSpace(output)

	if output != "false" {
		test.Fatalf("Incorrect variable value, expected false, got %s", output)
	}
}

func TestStructEval(test *testing.T) {
	input := `struct myStruct (int x, bool y);
				set myStruct.x = 123;
				set myStruct.y = false;
				myStruct.x;
				myStruct.y;`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	output, errors := EvaluateProgram(program)
	if len(errors) != 0 {
		test.Fatalf("Errors were found: %v", errors)
	}

	output = strings.TrimSpace(output)

	if output != "123\nfalse" {
		test.Fatalf("Incorrect variable value, expected \n123\nfalse, got \n%s", output)
	}
}

func TestPrefixEval(test *testing.T) {
	input := `-123; !true;`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	output, errors := EvaluateProgram(program)
	if len(errors) != 0 {
		test.Fatalf("Errors were found: %v", errors)
	}

	output = strings.TrimSpace(output)

	if output != "-123\nfalse" {
		test.Fatalf("Incorrect variable value, expected \n-123\nfalse, got \n%s", output)
	}
}

func TestInfixEval(test *testing.T) {
	input := `1+1;2-2;10/5;8*8;1>0;1>=1;1==1;1!=1;1<2;1<=1;`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	output, errors := EvaluateProgram(program)
	if len(errors) != 0 {
		test.Fatalf("Errors were found: %v", errors)
	}

	output = strings.TrimSpace(output)

	if output != "2\n0\n2\n64\ntrue\ntrue\ntrue\nfalse\ntrue\ntrue" {
		test.Fatalf("Incorrect variable value, expected \n2\n0\n2\n64\ntrue\ntrue\ntrue\nfalse\ntrue\ntrue\ngot \n%s", output)
	}
}
