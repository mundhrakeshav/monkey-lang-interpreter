package ast

import "github.com/mundhrakeshav/monkey-lang-interpreter/token"

type Node interface {
	// returns literal value associate with node
	TokenLiteral() string
}

// All statement nodes implement this
type Statement interface {
	Node
	// Dummy method to differentiate StatementNode and ExpressionNode
	statementNode()
}

// All expression nodes implement this
type Expression interface {
	Node
	// Dummy method to differentiate StatementNode and ExpressionNode
	expressionNode()
}

// ---------------------------------------------
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// ---------------------------------------------

type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier // Identifier to bind to
	Value Expression  // expression that produces a value
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
