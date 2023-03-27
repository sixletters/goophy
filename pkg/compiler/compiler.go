package compiler

import (
	"cs4215/goophy/pkg/ast"
	"fmt"
	"strconv"
)

/* idk how to fix this builtin mapping need change all the functions from JS to their GO equivalent maybe we do our own built-in mappings for what we want specifically
var builtin_mapping = map[string] func() {
	"display":   println(),
	"get_time":  time.Now(),
	"stringify": stringify(),
	"error": func(x interface{}) interface{} {
		PC = len(instrs) - 1
		return x
	},
	"prompt":       prompt(),
	"is_number":    is_number(),
	"is_string":    is_string(),
	"is_boolean":   is_boolean(),
	"is_undefined": is_undefined(),
	"parse_int":    parse_int(),
	"char_at":      char_at(),
	"math_abs":     math_abs(),
	"math_acos":    math_acos(),
	"math_acosh":   math_acosh(),
	"math_asin":    math_asin(),
	"math_asinh":   math_asinh(),
	"math_atan":    math_atan(),
	"math_atanh":   math_atanh(),
	"math_atan2":   math_atan2(),
	"math_ceil":    math_ceil(),
	"math_cbrt":    math_cbrt(),
	"math_expm1":   math_expm1(),
	"math_clz32":   math_clz32(),
	"math_cos":     math_cos(),
	"math_cosh":    math_cosh(),
	"math_exp":     math_exp(),
	"math_floor":   math_floor(),
	"math_fround":  math_fround(),
	"math_hypot":   math_hypot(),
	"math_imul":    math_imul(),
	"math_log":     math_log(),
	"math_log1p":   math_log1p(),
	"math_log2":    math_log2(),
	"math_log10":   math_log10(),
	"math_max":     math_max(),
	"math_min":     math_min(),
	"math_pow":     math_pow(),
	"math_random":  math_random(),
	"math_round":   math_round(),
	"math_sign":    math_sign(),
	"math_sin":     math_sin(),
	"math_sinh":    math_sinh(),
	"math_sqrt":    math_sqrt(),
	"math_tanh":    math_tanh(),
	"math_trunc":   math_trunc(),
	"pair":         pair(),
	"is_pair":      is_pair(),
	"head":         head(),
	"tail":         tail(),
	"is_null":      is_null(),
	"set_head":     set_head(),
	"set_tail":     set_tail(),
	"array_length": array_length(),
	"is_array":     is_array(),
	"list":         list(),
	"is_list":      is_list(),
	"display_list": display_list(),
	// from list libarary
	"equal":          equal(),
	"length":         length(),
	"list_to_string": list_to_string(),
	"reverse":        reverse(),
	"append":         append(),
	"member":         member(),
	"remove":         remove(),
	"remove_all":     remove_all(),
	"enum_list":      enum_list(),
	"list_ref":       list_ref(),
	// misc
	"draw_data": draw_data(),
	"parse":     parse(),
	"tokenize":  tokenize(),
}*/

var wc = 0

func scan(statements []ast.Statement) []string {
	var result []string
	for _, statement := range statements {
		statement, ok := statement.(*ast.LetStatement)
		if ok {
			result = append(result, statement.Name.Value)
		}
	}
	return result
}

// not done yet for expression statement
func Compile_statement(statement ast.Statement, instrs []Instruction) []Instruction {
	token := statement.GetToken()
	fmt.Println(token.Type)
	fmt.Println(token.Literal)
	switch token.Type {
	case "LET":
		assignStatement := statement.(*ast.LetStatement)
		newInstrs := Compile_expression(assignStatement.Value, instrs)
		instrs = append(instrs, newInstrs...)
		wc += 1
		assignInstruction := ASSIGNInstruction{Tag: "ASSIGN", Sym: assignStatement.Name.Value}
		instrs = append(instrs, assignInstruction)
	case "RETURN":
		returnStatement := statement.(*ast.ReturnStatement)
		newInstrs := Compile_expression(returnStatement.ReturnValue, instrs)
		instrs = append(instrs, newInstrs...)
		// i need to change call into tailcall if its in the return statement
		/*returnStatement, ok := returnStatement.ReturnValue.(*ast.CallStatement)
		// I havent even do a normal call expresison yet so wait
		if ok {
		} else {*/
		wc += 1
		resetInstruction := RESETInstruction{Tag: "RESET"}
		instrs = append(instrs, resetInstruction)
	case "{":
		//block is broken, i not sure how to fix the compile statement step in the loop for statements
		blkStatement := statement.(*ast.BlockStatement)
		locals := scan(blkStatement.Statements)
		wc += 1
		enterScopeInstruction := ENTERSCOPEInstruction{Tag: "ENTER_SCOPE", Syms: locals}
		instrs = append(instrs, enterScopeInstruction)
		for _, statement := range blkStatement.Statements {
			fmt.Print(statement)
			newInstrs := Compile_statement(statement, []Instruction{})
			instrs = append(instrs, newInstrs...)
		}
		wc += 1
		exitScopeInstruction := EXITSCOPEInstruction{Tag: "EXIT_SCOPE"}
		instrs = append(instrs, exitScopeInstruction)
	default:
		expressionStatement := statement.(*ast.ExpressionStatement)
		newInstrs := Compile_expression(expressionStatement.Expression, instrs)
		instrs = append(instrs, newInstrs...)
	}
	return instrs
}

// WIP
func Compile_expression(expression ast.Expression, instrs []Instruction) []Instruction {
	token := expression.GetToken()
	switch token.Type {
	case "ILLEGAL":
		panic("ILLEGAL EXPRESSION ENCOUNTERED")
	case "EOF":
		wc += 1
		doneInstruction := DONEInstruction{Tag: "DONE"}
		instrs = append(instrs, doneInstruction)
	case "IDENT":
		wc += 1
		ldsInstruction := LDSInstruction{Tag: "LDS", Sym: token.Literal}
		instrs = append(instrs, ldsInstruction)
	case "TRUE", "FALSE":
		val, _ := strconv.ParseBool(token.Literal)
		wc += 1
		ldcbInstruction := LDCBInstruction{Tag: "LDCB", Val: val}
		instrs = append(instrs, ldcbInstruction)
	case "INT":
		val, _ := strconv.Atoi(token.Literal)
		wc += 1
		ldcnInstruction := LDCNInstruction{Tag: "LDCN", Val: val}
		instrs = append(instrs, ldcnInstruction)
	case "+", "*", "/", "<=", ">", "==", "!=": /*tokens have not included modulo*/
		expr := expression.(*ast.InfixExpression)
		newInstrs := Compile_expression(expr.Left, []Instruction{})
		instrs = append(instrs, newInstrs...)
		newerInstrs := Compile_expression(expr.Right, []Instruction{})
		instrs = append(instrs, newerInstrs...)
		wc += 1
		binopInstruction := BINOPInstruction{Tag: "BINOP", Sym: BINOPS(token.Literal)}
		instrs = append(instrs, binopInstruction)
	case "-":
		// both negative number and binop minus is handled here
		expr, ok := expression.(*ast.PrefixExpression)
		if ok {
			newerInstrs := Compile_expression(expr.Right, []Instruction{})
			instrs = append(instrs, newerInstrs...)
			wc += 1
			unopInstruction := UNOPInstruction{Tag: "UNOP", Sym: "-unary"}
			instrs = append(instrs, unopInstruction)
		} else {
			expr := expression.(*ast.InfixExpression)
			newInstrs := Compile_expression(expr.Left, []Instruction{})
			instrs = append(instrs, newInstrs...)
			newerInstrs := Compile_expression(expr.Right, []Instruction{})
			instrs = append(instrs, newerInstrs...)
			wc += 1
			binopInstruction := BINOPInstruction{Tag: "BINOP", Sym: BINOPS(token.Literal)}
			instrs = append(instrs, binopInstruction)
		}
	case "FUNCTION":
		functionLiteral := expression.(*ast.FunctionLiteral)
		wc += 1
		var parametersToString []string
		for i, parameter := range functionLiteral.Parameters {
			parametersToString[i] = parameter.Value
		}
		ldfInstruction := LDFInstruction{Tag: "LDF", Prms: parametersToString, Addr: wc + 1}
		instrs = append(instrs, ldfInstruction)
		bodyInstrs := Compile_statement(functionLiteral.Body, []Instruction{})
		wc += 1
		gotoInstruction := GOTOInstruction{Tag: "GOTO", Addr: wc}
		instrs = append(instrs, gotoInstruction)
		instrs = append(instrs, bodyInstrs...)
		wc += 1
		assignInstruction := ASSIGNInstruction{Tag: "ASSIGN", Sym: token.Literal}
		instrs = append(instrs, assignInstruction)
	case "(":
		callExpression := expression.(*ast.CallExpression)
		functionInstrs := Compile_expression(callExpression.Function, []Instruction{})
		instrs = append(instrs, functionInstrs...)
		var length = 0
		for _, arg := range callExpression.Arguments {
			argumentInstrs := Compile_expression(arg, []Instruction{})
			instrs = append(instrs, argumentInstrs...)
			length += 1
		}
		wc += 1
		callInstruction := CALLInstruction{Tag: "CALL", Arity: length}
		instrs = append(instrs, callInstruction)
	case "IF":
		ifExpression := expression.(*ast.IfExpression)
		condInstrs := Compile_expression(ifExpression.Condition, []Instruction{})
		instrs = append(instrs, condInstrs...)
		ifBlockInstrs := Compile_statement(ifExpression.IfBlock, []Instruction{})
		jofInstruction := JOFInstruction{Tag: "JOF", Addr: wc}
		elseBlockInstrs := Compile_statement(ifExpression.ElseBlock, []Instruction{})
		gotoInstruction := GOTOInstruction{Tag: "GOTO", Addr: wc}
		instrs = append(instrs, jofInstruction)
		instrs = append(instrs, ifBlockInstrs...)
		instrs = append(instrs, gotoInstruction)
		instrs = append(instrs, elseBlockInstrs...)
	}
	return instrs
}

func Compile(program ast.Program) []Instruction {
	instrs := make([]Instruction, 0)
	locals := scan(program.Statements)
	wc += 1
	enterScopeInstruction := ENTERSCOPEInstruction{Tag: "ENTER_SCOPE", Syms: locals}
	instrs = append(instrs, enterScopeInstruction)
	for i, statement := range program.Statements {
		instrs = append(instrs, Compile_statement(statement, []Instruction{})...)
		if i != len(program.Statements)-1 {
			wc += 1
			popInstruction := POPInstruction{Tag: "POP"}
			instrs = append(instrs, popInstruction)
		}
	}
	wc += 1
	exitScopeInstruction := EXITSCOPEInstruction{Tag: "EXIT_SCOPE"}
	instrs = append(instrs, exitScopeInstruction)
	wc += 1
	doneInstruction := DONEInstruction{Tag: "DONE"}
	instrs = append(instrs, doneInstruction)
	return instrs
}
