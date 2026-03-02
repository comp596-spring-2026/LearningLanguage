package lexer

import (
	"learningLanguage/token"
)

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

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
