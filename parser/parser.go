package parser

import (
	// "github.com/mundhrakeshav/monkey-lang-interpreter/ast"
	"github.com/mundhrakeshav/monkey-lang-interpreter/ast"
	"github.com/mundhrakeshav/monkey-lang-interpreter/lexer"
	"github.com/mundhrakeshav/monkey-lang-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token // Points to current token
	peekToken token.Token // Points to next token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	/*Set currToken and peekToken*/
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}
