package token

/*
Language subset currently being worked on:
create int a;
create int b;
set b = 3;
set a = 5;
set a = a * 1;
set b = b / 5;
set a = a - 1
set b = b + 1

if (a > b) begin;
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
NEQ
NOT
BEGIN
END
PLUS
MINUS
DIVIDE
MULTIPLY
LPAREN
RPAREN
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
	BOOL   = "BOOL"
	NUMBER = "NUMBER"

	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	DIVIDE   = "DIVIDE"
	MULTIPLY = "MULTIPLY"

	EQ  = "EQUALTO"
	GT  = "GREATER"
	GE  = "GREQUAL"
	LT  = "LESS"
	LE  = "LEQUAL"
	NEQ = "NOTEQUAL"
	NOT = "NOT"

	LPAREN    = "LPAREN"
	RPAEREN   = "RPAREN"
	SEMICOLON = "SEMICOLON"

	SET    = "SET"
	CREATE = "CREATE"
	IF     = "IF"
	ELSE   = "ELSE"
	BEGIN  = "BEGIN"
	END    = "END"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
)
