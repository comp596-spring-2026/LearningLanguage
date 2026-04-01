package evaluation

import (
	"bytes"
	"learningLanguage/ast"
)

var variableMap = make(map[string]int64)

func evaluateProgram(program ast.Program) string {
	var output bytes.Buffer
	for _, statement := range program.Statements {
		output.WriteString(evaluateStatement(statement))
	}

	return output.String()
}

func evaluateStatement(statement ast.Statement) string {
	var output string
	createStmt, ok := statement.(*ast.CreateStatement)
	if ok {
		output = evaluateCreateStatement(createStmt)
	}

	setStmt, ok := statement.(*ast.SetStatement)
	if ok {
		output = evaluateSetStatement(setStmt)
	}

	exprStmt, ok := statement.(*ast.ExpressionStatement)
	if ok {
		output = evaluateExpressionStatement(exprStmt)
	}

	return output
}

func evaluateCreateStatement(statement *ast.CreateStatement) string {
	return ""
}

func evaluateSetStatement(statement *ast.SetStatement) string {
	return ""
}

func evaluateExpressionStatement(statement *ast.ExpressionStatement) string {
	return ""
}
