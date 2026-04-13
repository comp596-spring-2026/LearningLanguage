package token

/*
Language feature currently being worked on:
struct myStruct(
	int a,
	bool b
) [a: 123, b: true];

OR

struct myStruct(int a, bool b);
myStruct.a = 123;
myStruct.b = true;

Tokens:
STRUCT
COMMA
COLON
DOT
LBRACKET
RBRACKET
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
	STRUCT = "STRUCT"

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
	COLON     = "COLON"
	DOT       = "DOT"
	COMMA     = "COMMA"
	LBRACKET  = "LBRACKET"
	RBRACKET  = "RBRACKET"

	SET    = "SET"
	CREATE = "CREATE"
	IF     = "IF"
	ELSE   = "ELSE"
	BEGIN  = "BEGIN"
	END    = "END"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
)
