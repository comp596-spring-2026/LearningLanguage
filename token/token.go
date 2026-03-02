package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	EQ = "=="
	GT = ">"
	GE = ">="
	LT = "<"
	LE = "<="

	LPAREN    = "("
	RPAEREN   = ")"
	SEMICOLON = ";"

	VARIABLE = "VARIABLE"
	CONSTANT = "CONSTANT"
	IF       = "IF"
	ELSE     = "ELSE"
	BEGIN    = "BEGIN"
	END      = "END"
)
