package compiler

import (
	"cs4215/goophy/pkg/ast"
	"reflect"
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

/*
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	LT       TokenType = "<"
	GT       TokenType = ">"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"

	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="
*/

func scan(statements []Statement) []string {
	var result []string
	for i, statement := range statements {
		if (reflect.TypeOf(statement) == LetStatement) {
			result = append(result, statement.Name.Value)
		}
	}
	return result;
}

//not done yet for expression statement
func compile_statement(statement ast.Statement, instrsArray Instructions) Instructions {
	instrs := instrsArray.getInstrs()
	token := statement.node
	switch token.TokenLiteral() {
	case "LET":
		letStatement := statement(ast.LetStatement)
		compile_expression(letStatement.Value, instrsArray)
		letInstruction := LETInstruction{tag: "LET", sym: letStatement.Name.Value}
		instrs = append(instrs, letInstruction)
	case "RETURN":
		returnStatement := statement(ast.ReturnStatement)
		compile_expression(returnStatement.ReturnValue, instrsArray)
		// i need to change call into tailcall if its in the return statement
		if (reflect.TypeOf(returnStatement.ReturnValue) == CallExpression) {
			// I havent even do a normal call expresison yet so wait
		} else {
			resetInstruction := RESETInstruction{tag: "RESET"}
			instrs = append(instrs, resetInstruction)
		}
	case "LBRACE" :
		blkStatement := statement(ast.BlockStatement)
		locals := scan(blkStatement.Statements)
		enterScopeInstruction := ENTERSCOPEInstruction{tag: "ENTER_SCOPE", syms: locals}
		instrs = append(instrs, enterScopeInstruction)
		for i, statement := range blkStatement.Statements {
			instrs = compile_statement(statement, instrs)
		}		
		exitScopeInstruction := EXITSCOPEInstruction{tag: "EXIT_SCOPE", syms: locals}
		instrs = append(instrs, exitScopeInstruction)
	}
	return instrsArray
}

//WIP
func compile_expression(expression ast.Expression, instrsArray Instructions) Instructions {
	instrs := instrsArray.getInstrs()
	token := statement.node
	switch token.TokenLiteral() {
	case "ILLEGAL":
		panic("ILLEGAL EXPRESSION ENCOUNTERED")
	case "EOF":
		doneInstruction := DONEInstruction{tag: "DONE"}
		instrs = append(instrs, doneInstruction)
	case "IDENT":
		ldsInstruction := LDSInstruction{tag: "LDS", sym: token.String()}
		instrs = append(instrs, ldsInstruction)
	case "INT":
		ldcnInstruction := LDCNInstruction{tag: "LDCN", val: token.String()}
		instrs = append(instrs, ldcnInstruction)
	case "LET":
		letInstruction := LETInstruction{tag: "LET", sym: token.String()}
		instrs = append(instrs, letInstruction)
	case "PLUS", "MINUS", "ASTERISK", "SLASH", "LT", "GT", "EQ", "NOT_EQ": /*tokens have not included modulo*/
		compile_expression(, instrsArray)
		binopInstruction := BINOPInstruction{tag: "BINOP", sym: token.String()}
		instrs = append(instrs, binopInstruction)
	case "BANG": /* UNARY MINUS not ready yet*/
		unopInstruction := UNOPInstruction{tag: "UNOP", sym: token.String()}
		instrs = append(instrs, unopInstruction)
	case "TRUE", "FALSE":
		ldcbInstruction := LDCBInstruction{tag: "LDCB", val: token.String()}
		instrs = append(instrs, ldcbInstruction)
	case "RETURN":
		resetInstruction := RESETInstruction{tag: "RESET"}
		instrs = append(instrs, resetInstruction)

		/* BELOW IS A WORK IN PROGRESS
		case "name":
			instrs[wc] = &instruction{tag: "LD", sym: comp.sym, pos: compileTimeEnvironmentPosition(ce, comp.sym)}
			wc++
		case "unop":
			compile(comp.frst, ce)
			instrs[wc] = &instruction{tag: "UNOP", sym: comp.sym}
			wc++
		case "binop":
			compile(comp.frst, ce)
			compile(comp.scnd, ce)
			instrs[wc] = &instruction{tag: "BINOP", sym: comp.sym}
			wc++
		case "log":
			var pred, cons, alt *astNode
			if comp.sym == "&&" {
				pred, cons, alt = comp.frst, &astNode{tag: "lit", val: true}, comp.scnd
			} else {
				pred, cons, alt = comp.frst, comp.scnd, &astNode{tag: "lit", val: false}
			}
			compile(&astNode{tag: "cond_expr", pred: pred, cons: cons, alt: alt}, ce)
		case "cond":
			compile(comp.pred, ce)
			jumpOnFalseInstruction := &instruction{tag: "JOF"}
			instrs[wc] = jumpOnFalseInstruction
			wc++
			compile(comp.cons, ce)
			gotoInstruction := &instruction{tag: "GOTO"}
			instrs[wc] = gotoInstruction
			wc++
			alternativeAddress := wc
			compile(comp.alt, ce)
			gotoInstruction.addr = wc
			jumpOnFalseInstruction.addr = alternativeAddress
		case "while":
			loopStart := wc
			compile(comp.pred, ce)
			jumpOnFalseInstruction := instruction{tag: "JOF"}
			instrs[wc] = jumpOnFalseInstruction
			wc++
			compile(comp.body, ce)
			instrs[wc] = instruction{tag: "POP"}
			wc++
			instrs[wc] = instruction{tag: "GOTO", addr: loopStart}
			wc++
			jumpOnFalseInstruction.addr = wc
			instrs[wc] = instruction{tag: "LDC", val: Undefined}
			wc++
		case "app":
			compile(comp.fun, ce)
			for _, arg := range comp.args {
				compile(arg, ce)
			}
			instrs[wc] = instruction{tag: "CALL", arity: len(comp.args)}
			wc++
		case "assmt":
			compile(comp.expr, ce)
			instrs[wc] = instruction{tag: "ASSIGN", pos: compileTimeEnvironmentPosition(ce, comp.sym)}
			wc++
		case "lam":
			instrs[wc] = instruction{tag: "LDF", arity: comp.arity, addr: wc + 1}
			gotoInstruction := instruction{tag: "GOTO"}
			instrs[wc] = gotoInstruction
			wc++
			compile(comp.body, compileTimeEnvironmentExtend(comp.prms, ce))
			instrs[wc] = instruction{tag: "LDC", val: Undefined}
			wc++
			instrs[wc] = instruction{tag: "RESET"}
			wc++
			gotoInstruction.addr = wc
		case "seq":
			compileSequence(comp.stmts, ce)
		case "blk":
			locals := scan(comp.body)
			instrs[wc] = instruction{tag: "ENTER_SCOPE", num: len(locals)}
			wc++
			compile(comp.body, compileTimeEnvironmentExtend(locals, ce))
			instrs[wc] = instruction{tag: "EXIT_SCOPE"}
			wc++
		case "let", "const":
			compile(comp.expr, ce)
			instrs[wc] = instruction{tag: "ASSIGN", pos: compileTimeEnvironmentPosition(ce, comp.sym)}
			wc++
		case "ret":
			compile(comp.expr, ce)
			if comp.expr.tag == "app" {
				instrs[wc-1].tag = "TAIL_CALL"
			} else {
				instrs[wc] = instruction{tag: "RESET"}
				wc++
			}
		case "fun":
			compile(expression{tag: "const", sym: comp.sym, expr: expression{tag: "lam", prms: comp.prms, body: comp.body}}, ce)*/
	}
	return instrsArray
}
