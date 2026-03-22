package token

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

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	NUMBER = "NUMBER"

	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"
	MINUS  = "MINUS"

	EQ = "EQUALTO"
	GT = "GREATER"
	GE = "GREQUAL"
	LT = "LESS"
	LE = "LEQUAL"

	LPAREN    = "LPAREN"
	RPAEREN   = "RPAREN"
	SEMICOLON = "SEMICOLON"

	SET    = "SET"
	CREATE = "CREATE"
	IF     = "IF"
	ELSE   = "ELSE"
	BEGIN  = "BEGIN"
	END    = "END"
)
