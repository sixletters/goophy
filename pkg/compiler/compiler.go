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
	switch token.Type {
	case "LET":
		assignStatement := statement.(*ast.LetStatement)
		newInstrs := Compile_expression(assignStatement.Value, instrs)
		instrs = append(instrs, newInstrs...)
		assignInstruction := ASSIGNInstruction{Tag: "ASSIGN", sym: assignStatement.Name.Value}
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
		resetInstruction := RESETInstruction{Tag: "RESET"}
		instrs = append(instrs, resetInstruction)
	/*case "LBRACE":
	blkStatement := statement.(*ast.BlockStatement)
	locals := scan(blkStatement.Statements)
	enterScopeInstruction := ENTERSCOPEInstruction{Tag: "ENTER_SCOPE", syms: locals}
	instrs = append(instrs, enterScopeInstruction)
	for i, statement := range blkStatement.Statements {
		newInstrs := Compile_statement(statement, instrs)
		instrs = append(instrs, newInstrs...)
	}
	exitScopeInstruction := EXITSCOPEInstruction{Tag: "EXIT_SCOPE"}
	instrs = append(instrs, exitScopeInstruction)*/
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
	fmt.Println(token.Type)
	switch token.Type {
	case "ILLEGAL":
		panic("ILLEGAL EXPRESSION ENCOUNTERED")
	case "EOF":
		doneInstruction := DONEInstruction{Tag: "DONE"}
		instrs = append(instrs, doneInstruction)
	case "IDENT":
		ldsInstruction := LDSInstruction{Tag: "LDS", Sym: token.Literal}
		instrs = append(instrs, ldsInstruction)
	case "TRUE", "FALSE":
		val, _ := strconv.ParseBool(token.Literal)
		ldcbInstruction := LDCBInstruction{Tag: "LDCB", Val: val}
		instrs = append(instrs, ldcbInstruction)
	case "INT":
		val, _ := strconv.Atoi(token.Literal)
		ldcnInstruction := LDCNInstruction{Tag: "LDCN", Val: val}
		instrs = append(instrs, ldcnInstruction)
	case "+", "-", "*", "/", "<=", ">", "==", "!=": /*tokens have not included modulo*/
		expr := expression.(*ast.InfixExpression)
		newInstrs := Compile_expression(expr.Left, []Instruction{})
		instrs = append(instrs, newInstrs...)
		newerInstrs := Compile_expression(expr.Right, []Instruction{})
		instrs = append(instrs, newerInstrs...)
		binopInstruction := BINOPInstruction{Tag: "BINOP", Sym: BINOPS(token.Literal)}
		instrs = append(instrs, binopInstruction)
	case "BANG":
	}
	return instrs
}

func Compile(program ast.Program) []Instruction {
	instrs := make([]Instruction, 0)
	locals := scan(program.Statements)
	enterScopeInstruction := ENTERSCOPEInstruction{Tag: "ENTER_SCOPE", syms: locals}
	instrs = append(instrs, enterScopeInstruction)
	for i, statement := range program.Statements {
		instrs = append(instrs, Compile_statement(statement, []Instruction{})...)
		if i != len(program.Statements)-1 {
			popInstruction := POPInstruction{Tag: "POP"}
			instrs = append(instrs, popInstruction)
		}
	}
	exitScopeInstruction := EXITSCOPEInstruction{Tag: "EXIT_SCOPE"}
	instrs = append(instrs, exitScopeInstruction)
	doneInstruction := DONEInstruction{Tag: "DONE"}
	instrs = append(instrs, doneInstruction)
	return instrs
}
