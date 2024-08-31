package parser

import (
	"interpreter-go/ast"
	"interpreter-go/lexer"
	"interpreter-go/token"
	"testing"
)

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return add(10, 5);
	`

	lex := lexer.New(input)
	prser := New(lex)

	program := prser.parseProgram()
	checkParserErrors(t, prser)
	if program == nil {
		t.Fatalf("Program is nil.")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Found: %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Stmt not an *ast.ReturnStatement, actually is %T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 10833;
	`

	lex := lexer.New(input)
	prser := New(lex)

	program := prser.parseProgram()
	checkParserErrors(t, prser)
	if program == nil {
		t.Fatalf("Program is nil.")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements did not have 3 statements. Got %d statements.", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tl := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tl.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, expectedId string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. Found: %q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. Found: %T", stmt)
		return false
	}

	if letStmt.Name.Value != expectedId {
		t.Errorf("letStmt.Name.Value not '%s'. Found: %s", expectedId, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != expectedId {
		t.Errorf("letStmt.Name not %s. Found: %s", expectedId, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errs := p.Errors()
	if len(errs) == 0 {
		return
	}

	t.Errorf("Parser had %d errors:", len(errs))
	for _, msg := range errs {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. Found: %q", program.String())
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`

	l := lexer.New(input)
	p := New(l)
	program := p.parseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program doesn't have expected statements. Found: %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. Found type: %T",
			program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. Found type: %T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. Found: %s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. Found: %s", "foobar", ident.TokenLiteral())
	}
}
