package evaluation

import (
	"bytes"
	"fmt"
	"learningLanguage/ast"
	"strconv"
)

const (
	INTTYPE = iota
	BOOLTYPE
	STRINGTYPE
)

type Data struct {
	dataType    int
	intValue    int64
	boolValue   bool
	stringValue string
}

var variableTypes = map[string]int{
	"int":  INTTYPE,
	"bool": BOOLTYPE,
}
var variableMap = make(map[string]Data)
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

	ifStmt, ok := statement.(*ast.IfStatement)
	if ok {
		output = evaluateIfStatement(ifStmt)
	}

	exprStmt, ok := statement.(*ast.ExpressionStatement)
	if ok {
		value := evaluateExpression(exprStmt.Expression)
		switch value.dataType {
		case INTTYPE:
			output = strconv.FormatInt(value.intValue, 10)
		case BOOLTYPE:
			output = strconv.FormatBool(value.boolValue)
		}
	}

	return output
}

func evaluateCreateStatement(statement *ast.CreateStatement) string {
	variableMap[statement.Name.Value] = Data{dataType: variableTypes[statement.DataType]}
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

func evaluateIfStatement(statement *ast.IfStatement) string {
	conditionData := evaluateExpression(statement.Condition)
	if conditionData.dataType != BOOLTYPE {
		errors = append(errors, "Cannot use non-boolean expression in if statement condition.")
		return ""
	} else {
		if conditionData.boolValue {
			return evaluateStatement(statement.IfTrue)
		} else {
			return evaluateStatement(statement.Else)
		}
	}
}

func evaluateExpression(expression ast.Expression) Data {
	var value Data

	intLit, ok := expression.(*ast.IntegerLiteral)
	if ok {
		value = evaluateIntLit(intLit)
	}

	boolLit, ok := expression.(*ast.BooleanLiteral)
	if ok {
		value = evaluateBoolLit(boolLit)
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

func evaluateIntLit(expression *ast.IntegerLiteral) Data {
	return Data{dataType: INTTYPE, intValue: expression.Value}
}

func evaluateBoolLit(expression *ast.BooleanLiteral) Data {
	return Data{dataType: BOOLTYPE, boolValue: expression.Value}

}

func evaluateIdentifier(identifier *ast.Identifier) Data {
	value, ok := variableMap[identifier.Value]
	if !ok {
		errors = append(errors, fmt.Sprintf("Variable %s does not exist.", identifier.Value))
		return Data{}
	}

	return value
}

func evaluatePrefixExp(expression *ast.PrefixExpression) Data {
	value := evaluateExpression(expression.Right)
	switch expression.Operator {
	case "-":
		value.intValue = -1 * value.intValue
	case "!":
		value.boolValue = !value.boolValue
	}

	return value
}

func evaluateInfixExp(expression *ast.InfixExpression) Data {
	leftValue := evaluateExpression(expression.Left)
	rightValue := evaluateExpression(expression.Right)
	if leftValue.dataType != rightValue.dataType {
		errors = append(errors, "Mismatching datatypes error")
		return Data{}
	}
	var retValue Data

	switch expression.Operator {
	case "+":
		retValue.dataType = INTTYPE
		retValue.intValue = leftValue.intValue + rightValue.intValue
	case "-":
		retValue.dataType = INTTYPE
		retValue.intValue = leftValue.intValue - rightValue.intValue
	case "*":
		retValue.dataType = INTTYPE
		retValue.intValue = leftValue.intValue * rightValue.intValue
	case "/":
		retValue.dataType = INTTYPE
		retValue.intValue = leftValue.intValue / rightValue.intValue
	case "==":
		retValue.dataType = BOOLTYPE
		switch leftValue.dataType {
		case BOOLTYPE:
			retValue.boolValue = leftValue.boolValue == rightValue.boolValue
		case INTTYPE:
			retValue.boolValue = leftValue.intValue == rightValue.intValue
		}
	case "!=":
		retValue.dataType = BOOLTYPE
		switch leftValue.dataType {
		case BOOLTYPE:
			retValue.boolValue = leftValue.boolValue != rightValue.boolValue
		case INTTYPE:
			retValue.boolValue = leftValue.intValue != rightValue.intValue
		}
	case ">":
		retValue.dataType = BOOLTYPE
		if leftValue.dataType != INTTYPE && rightValue.dataType != INTTYPE {
			errors = append(errors, "Cannot perform perform quanitative comparisons with non-integers.")
			return Data{}
		}
		retValue.boolValue = leftValue.intValue > retValue.intValue
	case ">=":
		retValue.dataType = BOOLTYPE
		if leftValue.dataType != INTTYPE && rightValue.dataType != INTTYPE {
			errors = append(errors, "Cannot perform perform quanitative comparisons with non-integers.")
			return Data{}
		}
		retValue.boolValue = leftValue.intValue >= retValue.intValue
	case "<":
		retValue.dataType = BOOLTYPE
		if leftValue.dataType != INTTYPE && rightValue.dataType != INTTYPE {
			errors = append(errors, "Cannot perform perform quanitative comparisons with non-integers.")
			return Data{}
		}
		retValue.boolValue = leftValue.intValue < retValue.intValue
	case "<=":
		retValue.dataType = BOOLTYPE
		if leftValue.dataType != INTTYPE && rightValue.dataType != INTTYPE {
			errors = append(errors, "Cannot perform perform quanitative comparisons with non-integers.")
			return Data{}
		}
		retValue.boolValue = leftValue.intValue <= retValue.intValue
	}
	return retValue
}
