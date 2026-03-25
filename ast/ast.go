package ast

import (
	"bytes"
	"learningLanguage/token"
)

// Interfaces
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, statement := range p.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}

// Variable Creation Statement
type CreateStatement struct {
	Token token.Token
	Name  *Identifier
}

func (cs *CreateStatement) statementNode()       {}
func (cs *CreateStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *CreateStatement) String() string {
	var out bytes.Buffer

	out.WriteString("Identifier Name: ")
	out.WriteString(cs.Name.String())
	out.WriteString(".")

	return out.String()
}

// Variable Set Statement
type SetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ss *SetStatement) statementNode()       {}
func (ss *SetStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SetStatement) String() string {
	var out bytes.Buffer

	out.WriteString("Identifier Name: ")
	out.WriteString(ss.Name.String())
	out.WriteString(". Expression: ")
	out.WriteString(ss.Value.String())
	out.WriteString(".")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	return es.Expression.String()
}
