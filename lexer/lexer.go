package lexer

import (
	"learningLanguage/token"
	"slices"
)

var keywords = []string{"int", "variable", "constant", "if", "else", "begin", "end"}

type Lexer struct {
	input string
	head  int
	ch    byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.head >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.head]
	}
	l.head += 1
}

func (l *Lexer) goBack() {
	if l.head > 0 {
		l.head--
	}
}

func (l *Lexer) nextToken() token.Token {
	var tok token.Token

	if isLetter(l.ch) {
		tok = readString(l)
	} else {
		literal := string(l.ch)
		switch l.ch {
		case '=':
			temp := string(l.ch)
			l.readChar()
			if l.ch == '=' {
				temp += string(l.ch)
				tok = newToken(token.EQ, temp)
			} else {
				l.goBack()
				tok = newToken(token.ASSIGN, literal)
			}
		case '>':
			temp := string(l.ch)
			l.readChar()
			if l.ch == '=' {
				temp += string(l.ch)
				tok = newToken(token.GE, temp)
			} else {
				l.goBack()
				tok = newToken(token.GT, literal)
			}
		case '<':
			temp := string(l.ch)
			l.readChar()
			if l.ch == '=' {
				temp += string(l.ch)
				tok = newToken(token.LE, temp)
			} else {
				l.goBack()
				tok = newToken(token.LT, literal)
			}
		case ';':
			tok = newToken(token.SEMICOLON, literal)
		case '+':
			tok = newToken(token.PLUS, literal)
		case '-':
			tok = newToken(token.MINUS, literal)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func readString(l *Lexer) token.Token {
	var tok token.Token
	str := string(l.ch)
	l.readChar()
	for isLetter(l.ch) {
		str += string(l.ch)
		l.readChar()
	}
	isKeyword := checkKeyword(str)
	if isKeyword {
		tok = createKeyword(str)
	} else {
		tok = newToken(token.IDENT, str)
	}
	return tok
}

func checkKeyword(str string) bool {
	return slices.Contains(keywords, str)
}

func createKeyword(str string) token.Token {
	var tok token.Token
	switch str {
	case "variable":
		tok = newToken(token.VARIABLE, "VARIABLE")
	case "constant":
		tok = newToken(token.CONSTANT, "CONSTANT")
	case "if":
		tok = newToken(token.IF, "IF")
	case "else":
		tok = newToken(token.ELSE, "ELSE")
	case "begin":
		tok = newToken(token.BEGIN, "BEGIN")
	case "end":
		tok = newToken(token.END, "END")
	case "int":
		tok = newToken(token.INT, "INT")
	}
	return tok
}
