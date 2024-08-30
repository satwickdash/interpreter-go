package parser

import (
	"fmt"
	"interpreter-go/ast"
	"interpreter-go/lexer"
	"interpreter-go/token"
)

type Parser struct {
	l    *lexer.Lexer
	errs []string

	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) expectPeekIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.expectPeekIs(t) {
		p.nextToken()
		return true
	} else {
		p.addTokenError(t)
		return false
	}
}

func (p *Parser) addTokenError(tok token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s, found %q", tok, p.peekToken.Type)
	p.errs = append(p.errs, msg)
}

func (p *Parser) Errors() []string {
	return p.errs
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:    l,
		errs: []string{},
	}

	// To initialize p.curToken and p.peekToken
	p.nextToken()
	p.nextToken()
	return p
}
