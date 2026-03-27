package parser

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
	"fmt"
	"learningLanguage/ast"
	"learningLanguage/lexer"
	"learningLanguage/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.CREATE:
		stmt := p.parseCreateStatement()
		if stmt == nil {
			return nil
		}
		return stmt

	case token.SET:
		stmt := p.parseSetStatement()
		if stmt == nil {
			return nil
		}
		return stmt
	default:
		return nil
	}

}

func (p *Parser) checkNextToken(tokType token.TokenType) bool {
	if p.peekToken.Type == tokType {
		p.nextToken()
		return true
	} else {
		p.peekError(tokType)
		return false
	}
}

func (p *Parser) peekError(tokType token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s, received %s.", tokType, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseCreateStatement() *ast.CreateStatement {
	//CREATE DATATYPE(INT/FLOAT/ETC.) IDENT SEMICOLON
	statement := &ast.CreateStatement{Token: p.curToken}

	//DATATYPE
	if !p.checkNextToken(token.INT) { //TODO: add other datatypes
		return nil
	}

	//IDENT
	if !p.checkNextToken(token.IDENT) {
		return nil
	}
	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	//SEMICOLON
	if !p.checkNextToken(token.SEMICOLON) {
		return nil
	}

	return statement
}

func (p *Parser) parseSetStatement() *ast.SetStatement {
	//SET IDENT EQUALS EXPRESSION SEMICOLON
	statement := &ast.SetStatement{Token: p.curToken}

	//IDENT
	if !p.checkNextToken(token.IDENT) {
		return nil
	}
	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	//TODO: skipping all totems til end of statement, need to implement expressions
	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return statement
}
