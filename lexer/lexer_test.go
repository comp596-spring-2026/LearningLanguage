package lexer

import (
	"learningLanguage/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := "=+-;==<<=>>=int variable constant if else begin end asdf"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.SEMICOLON, ";"},
		{token.EQ, "=="},
		{token.LT, "<"},
		{token.LE, "<="},
		{token.GT, ">"},
		{token.GE, ">="},
		{token.INT, "INT"},
		{token.VARIABLE, "VARIABLE"},
		{token.CONSTANT, "CONSTANT"},
		{token.IF, "IF"},
		{token.ELSE, "ELSE"},
		{token.BEGIN, "BEGIN"},
		{token.END, "END"},
		{token.IDENT, "asdf"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.nextToken()
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
