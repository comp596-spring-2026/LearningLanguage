package parser

import (
	"fmt"
	"learningLanguage/ast"
	"learningLanguage/lexer"
	"testing"
)

func checkParserErrors(test *testing.T, p *Parser) {
	numErrors := len(p.errors)
	if numErrors == 0 {
		return
	}

	test.Errorf("Parser has %d errors.", numErrors)
	for _, msg := range p.errors {
		test.Errorf("Parser error: %q", msg)
	}
	test.FailNow()
}

func TestCreateStatements(test *testing.T) {
	input := `
	create int x; 
	create int y; 
	create int z;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(test, p)
	if program == nil {
		test.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		test.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"z"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testCreateStatement(test, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testCreateStatement(test *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "create" {
		test.Errorf("s.TokenLiteral not 'let'. got=%q", statement.TokenLiteral())
		return false
	}

	createStmt, ok := statement.(*ast.CreateStatement)
	if !ok {
		test.Errorf("s not *ast.LetStatement. got=%T", statement)
		return false
	}

	if createStmt.Name.Value != name {
		test.Errorf("letStmt.Name.Value not '%s'. got=%s", name, createStmt.Name.Value)
		return false
	}

	if createStmt.Name.TokenLiteral() != name {
		test.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, createStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestSetStatements(test *testing.T) {
	input := `
	set x = 3; 
	set y = 2; 
	set z = 223;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(test, p)
	if program == nil {
		test.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		test.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
		expectedIntLit     int64
	}{
		{"x", 3},
		{"y", 2},
		{"z", 223},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testSetStatement(test, stmt, tt.expectedIdentifier, tt.expectedIntLit) {
			return
		}
	}
}

func testSetStatement(test *testing.T, statement ast.Statement, name string, value int64) bool {
	if statement.TokenLiteral() != "set" {
		test.Errorf("s.TokenLiteral not 'set'. got=%q", statement.TokenLiteral())
		return false
	}

	setStmt, ok := statement.(*ast.SetStatement)
	if !ok {
		test.Errorf("s not *ast.SetStatement. got=%T", statement)
		return false
	}

	if setStmt.Name.Value != name {
		test.Errorf("setStmt.Name.Value not '%s'. got=%s", name, setStmt.Name.Value)
		return false
	}

	if setStmt.Name.TokenLiteral() != name {
		test.Errorf("setStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, setStmt.Name.TokenLiteral())
		return false
	}

	if !testIntegerLiteral(test, setStmt.Value, value) {
		return false
	}

	return true
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "2;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 2 {
		t.Errorf("literal.Value not %d. got=%d", 2, literal.Value)
	}
	if literal.TokenLiteral() != "2" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "2",
			literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"-22;", "-", 22},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				2, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}

	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"2+2;", 2, "+", 2},
		{"2-2;", 2, "-", 2},
		{"2*2;", 2, "*", 2},
		{"2/2;", 2, "/", 2},
	}

	for _, test := range infixTests {
		lex := lexer.New(test.input)
		parser := New(lex)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				2, len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		expression, ok := statement.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", statement.Expression)
		}

		if !testIntegerLiteral(t, expression.Left, test.leftValue) {
			return
		}

		if expression.Operator != test.operator {
			t.Fatalf("Operator is not %s. Got %s", test.operator, expression.Operator)
		}

		if !testIntegerLiteral(t, expression.Right, test.rightValue) {
			return
		}
	}
}
