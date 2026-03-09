package lexer

import (
	"learningLanguage/token"
	"testing"
)

func TestLexer(t *testing.T) {
	// input := "=+-;==<<=>>=int variable constant if else begin end asdf"
	input := "variable int a = 3;" //constant int b = 3;a = 5;if a > b begin;c = a - b; end;else begin;c = b - a; end;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// 	{token.ASSIGN, "="},
		// 	{token.PLUS, "+"},
		// 	{token.MINUS, "-"},
		// 	{token.SEMICOLON, ";"},
		// 	{token.EQ, "=="},
		// 	{token.LT, "<"},
		// 	{token.LE, "<="},
		// 	{token.GT, ">"},
		// 	{token.GE, ">="},
		// 	{token.INT, "INT"},
		// 	{token.VARIABLE, "VARIABLE"},
		// 	{token.CONSTANT, "CONSTANT"},
		// 	{token.IF, "IF"},
		// 	{token.ELSE, "ELSE"},
		// 	{token.BEGIN, "BEGIN"},
		// 	{token.END, "END"},
		// 	{token.IDENT, "asdf"},

		{token.VARIABLE, "VARIABLE"},
		{token.INTTYPE, "INTTYPE"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
