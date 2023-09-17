package ast

import "github.com/mundhrakeshav/monkey-lang-interpreter/token"

type Node interface {
	TokenLiteral() string
}

// All statement nodes implement this
type Statement interface {
	Node
	statementNode() string
}

// All expression nodes implement this
type Expression interface {
	Node
	expressionNode() string
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

type Identifier struct {
	Token token.Token
	Value string
}

//lint:ignore U1000 Ignore Unused
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatement struct {
	Token token.Token
	Name  *Identifier // Identifier of binding
	Value Expression  // Value for the expression that produces a value
}

//lint:ignore U1000 Ignore Unused
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
