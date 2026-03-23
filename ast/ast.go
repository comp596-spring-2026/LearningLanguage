package ast

import "learningLanguage/token"

//Interfaces
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//Variable Creation Statement
type CreateStatement struct {
	Token token.Token
	Name  *Identifier
}

func (ls *CreateStatement) statementNode()       {}
func (ls *CreateStatement) TokenLiteral() string { return ls.Token.Literal }

//Variable Set Statement
type SetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *SetStatement) statementNode()       {}
func (ls *SetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
