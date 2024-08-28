package parser

import (
	"interpreter-go/ast"
	"interpreter-go/lexer"
	"interpreter-go/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseProgram() *ast.Program {
	return nil
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	// To initialize p.curToken and p.peekToken
	p.nextToken()
	p.nextToken()
	return p
}
