package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT   = "IDENT"
	INTTYPE = "INTTYPE"
	INT     = "INT"

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

	VARIABLE = "VARIABLE"
	CONSTANT = "CONSTANT"
	IF       = "IF"
	ELSE     = "ELSE"
	BEGIN    = "BEGIN"
	END      = "END"
)
