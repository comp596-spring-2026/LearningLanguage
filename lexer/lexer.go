package lexer

/*
Language subset currently being worked on:
create int a;
create int b;
set b = 3;
set a = 5;

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
BEGIN
END
PLUS
MINUS
DIVIDE
MULTIPLY
LPAREN
RPAREN
*/

import (
	"learningLanguage/token"
	"slices"
)

var keywords = []string{"int", "create", "set", "if", "else", "begin", "end"}

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

func (l *Lexer) ignoreWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.ignoreWhitespace()

	if isLetter(l.ch) {
		tok = readString(l)
	} else if isNumber(l.ch) {
		tok = readNumber(l)
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
		default:
			tok.Literal = "ERROR"
			tok.Type = token.ILLEGAL
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func isNumber(char byte) bool {
	return '0' <= char && char <= '9'
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func readNumber(l *Lexer) token.Token {
	str := string(l.ch)
	l.readChar()
	for isNumber(l.ch) {
		str += string(l.ch)
		l.readChar()
	}
	l.goBack()
	return newToken(token.NUMBER, str)
}

func readString(l *Lexer) token.Token {
	var tok token.Token
	str := string(l.ch)
	l.readChar()
	for isLetter(l.ch) {
		str += string(l.ch)
		l.readChar()
	}
	l.goBack()
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
	case "create":
		tok = newToken(token.CREATE, str)
	case "set":
		tok = newToken(token.SET, str)
	case "if":
		tok = newToken(token.IF, str)
	case "else":
		tok = newToken(token.ELSE, str)
	case "begin":
		tok = newToken(token.BEGIN, str)
	case "end":
		tok = newToken(token.END, str)
	case "int":
		tok = newToken(token.INT, str)
	}
	return tok
}
