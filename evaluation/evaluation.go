package evaluation

import (
	"bytes"
	"fmt"
	"learningLanguage/ast"
	"strconv"
	"strings"
)

const (
	INTTYPE = iota
	BOOLTYPE
	STRINGTYPE
	FLOATTYPE
)

type Data struct {
	dataType    int
	intValue    int64
	boolValue   bool
	stringValue string
	floatValue  float64
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

	structStmt, ok := statement.(*ast.StructStatement)
	if ok {
		output = evaluateStructStatement(structStmt)
	}

	exprStmt, ok := statement.(*ast.ExpressionStatement)
	if ok {
		value := evaluateExpression(exprStmt.Expression)
		switch value.dataType {
		case INTTYPE:
			output = strconv.FormatInt(value.intValue, 10) + "\n"
		case BOOLTYPE:
			output = strconv.FormatBool(value.boolValue) + "\n"
		case FLOATTYPE:
			output = strconv.FormatFloat(value.floatValue, 'f', -1, 64) + "\n"
		case STRINGTYPE:
			output = strings.Trim(value.stringValue, "\"") + "\n"
		}
	}

	return output
}

func evaluateCreateStatement(statement *ast.CreateStatement) string {
	variableMap[statement.Name.Value] = Data{dataType: variableTypes[statement.Name.DataType]}
	return ""
}

func evaluateSetStatement(statement *ast.SetStatement) string {
	var name string
	if statement.Name.Attribute != "" {
		name = fmt.Sprintf("%s.%s", statement.Name.Value, statement.Name.Attribute)
	} else {
		name = statement.Name.Value
	}
	_, ok := variableMap[name]
	if !ok {
		errors = append(errors, fmt.Sprintf("Variable %s has not been created.", statement.Name.Value))
		return ""
	}
	variableMap[name] = evaluateExpression(statement.Value)
	return ""
}

func evaluateStructStatement(statement *ast.StructStatement) string {
	for _, attribute := range statement.Attributes {
		attributeName := fmt.Sprintf("%s.%s", statement.StructIdent.Value, attribute.Value)
		variableMap[attributeName] = evaluateExpression(statement.Values[attribute.Value])
	}
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

	floatLiteral, ok := expression.(*ast.FloatLiteral)
	if ok {
		value = evaluateFloatLit(floatLiteral)
	}

	strLit, ok := expression.(*ast.StringLiteral)
	if ok {
		value = evaluateStringLit(strLit)
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

func evaluateFloatLit(expression *ast.FloatLiteral) Data {
	return Data{dataType: FLOATTYPE, floatValue: expression.Value}
}

func evaluateStringLit(expression *ast.StringLiteral) Data {
	return Data{dataType: STRINGTYPE, stringValue: expression.Value}
}

func evaluateIdentifier(identifier *ast.Identifier) Data {
	var name string
	if identifier.Attribute != "" {
		name = fmt.Sprintf("%s.%s", identifier.Value, identifier.Attribute)
	} else {
		name = identifier.Value
	}
	value, ok := variableMap[name]
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
		retValue.boolValue = leftValue.intValue > rightValue.intValue
	case ">=":
		retValue.dataType = BOOLTYPE
		if leftValue.dataType != INTTYPE && rightValue.dataType != INTTYPE {
			errors = append(errors, "Cannot perform perform quanitative comparisons with non-integers.")
			return Data{}
		}
		retValue.boolValue = leftValue.intValue >= rightValue.intValue
	case "<":
		retValue.dataType = BOOLTYPE
		if leftValue.dataType != INTTYPE && rightValue.dataType != INTTYPE {
			errors = append(errors, "Cannot perform perform quanitative comparisons with non-integers.")
			return Data{}
		}
		retValue.boolValue = leftValue.intValue < rightValue.intValue
	case "<=":
		retValue.dataType = BOOLTYPE
		if leftValue.dataType != INTTYPE && rightValue.dataType != INTTYPE {
			errors = append(errors, "Cannot perform perform quanitative comparisons with non-integers.")
			return Data{}
		}
		retValue.boolValue = leftValue.intValue <= rightValue.intValue
	}
	return retValue
}
