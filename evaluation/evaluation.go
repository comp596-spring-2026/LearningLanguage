package evaluation

import (
	"bytes"
	"fmt"
	"learningLanguage/ast"
	"strconv"
)

var variableMap = make(map[string]int64)
var errors []string

func EvaluateProgram(program *ast.Program) (string, []string) {
	var output bytes.Buffer
	errors = []string{}
	for _, statement := range program.Statements {
		output.WriteString(evaluateStatement(statement))
	}

	return output.String(), errors
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
		value := evaluateExpression(exprStmt.Expression)
		output = strconv.FormatInt(value, 10)
	}

	return output
}

func evaluateCreateStatement(statement *ast.CreateStatement) string {
	variableMap[statement.Name.Value] = 0
	return ""
}

func evaluateSetStatement(statement *ast.SetStatement) string {
	_, ok := variableMap[statement.Name.Value]
	if !ok {
		errors = append(errors, fmt.Sprintf("Variable %s has not been created.", statement.Name.Value))
		return ""
	}
	variableMap[statement.Name.Value] = evaluateExpression(statement.Value)
	return ""
}

func evaluateExpression(expression ast.Expression) int64 {
	var value int64

	intLit, ok := expression.(*ast.IntegerLiteral)
	if ok {
		value = evaluateIntLit(intLit)
	}

	identifierExp, ok := expression.(*ast.Identifier)
	if ok {
		value = evaluateIdentifier(identifierExp)
	}

	prefixExp, ok := expression.(*ast.PrefixExpression)
	if ok {
		value = evaluatePrefixExp(prefixExp)
	}

	infixExp, ok := expression.(*ast.InfixExpression)
	if ok {
		value = evaluateInfixExp(infixExp)
	}

	return value
}

func evaluateIntLit(expression *ast.IntegerLiteral) int64 {
	return expression.Value
}

func evaluateIdentifier(identifier *ast.Identifier) int64 {
	value, ok := variableMap[identifier.Value]
	if !ok {
		errors = append(errors, fmt.Sprintf("Variable %s does not exist.", identifier.Value))
		return 0
	}
	return value
}

func evaluatePrefixExp(expression *ast.PrefixExpression) int64 {
	value := evaluateExpression(expression.Right)
	if expression.Operator == "-" {
		value = -1 * value
	}

	return value
}

func evaluateInfixExp(expression *ast.InfixExpression) int64 {
	leftValue := evaluateExpression(expression.Left)
	rightValue := evaluateExpression(expression.Right)

	switch expression.Operator {
	case "+":
		return leftValue + rightValue
	case "-":
		return leftValue - rightValue
	case "*":
		return leftValue * rightValue
	case "/":
		return leftValue / rightValue
	default:
		return 0
	}
}
