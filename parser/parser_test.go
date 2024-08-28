package parser

import (
	"interpreter-go/ast"
	"interpreter-go/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 10833;
	`

	lex := lexer.New(input)
	prser := New(lex)

	program := prser.parseProgram()
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
