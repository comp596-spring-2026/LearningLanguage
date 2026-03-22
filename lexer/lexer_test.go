package lexer

/*
Language subset currently being worked on:
create int a;
create int b;
set b = 3;
set a = 5;

if a > b begin;
c = a - b; end;
else begin;
c = b - a; end;

Tokens:
CREATE
SET
INT
IDENTIFIER
EQUAL
NUMBER
SEMICOLON
IF
GREATER
LESS
GE
LE
EQ
BEGIN
END
MINUS
*/

import (
	"learningLanguage/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := "create int a; set a = 3;"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CREATE, "create"},
		{token.INT, "int"},
		{token.IDENT, "a"},
		{token.SEMICOLON, ";"},
		{token.SET, "set"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.NUMBER, "3"},
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
