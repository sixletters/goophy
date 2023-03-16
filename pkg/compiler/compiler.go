package compiler

import (
	"pkg/ast"
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

// instrs: instruction array
var instrs []Instruction

// wc: write counter
var wc = 0

// The base Node interface
type Node interface {
	TagLiteral() string
	String() string
}

type Instruction interface {
	Node
	instructionNode()
}

type Instructions struct {
	instrs []Instruction
}

type LDCInstruction struct {
	Instruction
	Tag Tag // the LDC tag
}

func (ldc *LDCInstruction) instructionNode()   {}
func (ldc *LDCInstruction) TagLiteral() string { return ldc.Tag.Literal }
func (ldc *LDCInstruction) String() string     { return ldc.Tag.Literal }

type LDInstruction struct {
	Instruction
	Tag Tag // the LD Symbolic tag
}

func (ld *LDInstruction) instructionNode()   {}
func (ld *LDInstruction) TagLiteral() string { return ld.Tag.Literal }
func (ld *LDInstruction) String() string     { return ld.Tag.Literal }

type DONEInstruction struct {
	Tag Tag // the DONE tag
}

func (done *DONEInstruction) instructionNode()   {}
func (done *DONEInstruction) TagLiteral() string { return done.Tag.Literal }
func (done *DONEInstruction) String() string     { return done.Tag.Literal }

// switch cases
func compile_comp(token Token, statement Statement) {
	switch token.Literal {
	case "INT":
		ldcInstruction := LDCInstruction{Tag: NewTag("LDC", statement.Node.String())}
		instrs[wc] = ldcInstruction
		wc++
	case "IDENT":
		ldInstruction := LDInstruction{Tag: NewTag("LD", statement.Node.String())}
		instrs[wc] = ldInstruction
		wc++
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
}

// finish with a done instruction
func compile(statement Statement) {
	compile_comp(statement.tokenType, statement)
	doneInstruction := DONEInstruction{Tag: NewTag("DONE", statement.Node.String())}
	instrs[wc] = doneInstruction
}

func compileProgram(program ast.Program) {
	wc := 0
	instrs := make([]Instructions, 1000)
	compile(program.Statements[wc])
}
