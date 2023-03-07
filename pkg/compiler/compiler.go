package compiler
/*
// scanning out the declarations from (possibly nested)
// sequences of statements, ignoring blocks
func scan(comp *astNode) []string {
    switch comp.tag {
    case "seq":
        var decls []string
        for _, stmt := range comp.stmts {
            decls = append(decls, scan(stmt)...)
        }
        return decls
    case "let", "const", "fun":
        return []string{comp.sym}
    }
    return []string{}
}

func compile_sequence(seq []*astNode, ce compileEnvironment) {
    if len(seq) == 0 {
        instrs[wc] = &instruction{tag: "LDC", val: nil}
        wc++
        return
    }
    first := true
    for _, comp := range seq {
        if !first {
            instrs[wc] = &instruction{tag: "POP"}
            wc++
        } else {
            first = false
        }
        compile(comp, ce)
    }
}
// switch cases
func compile(comp expression, ce *CompileTimeEnvironment) {
    switch comp.tag {
	case "lit": 
		val := comp.val
		switch val.(type) {
		case nil:
			val = Null
		case bool:
			val = boolToValue(val.(bool))
		case string:
			val = addString(val.(string))
		}
		instrs[wc] = &instruction{tag: "LDC", val: val}
		wc++
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
        compile(expression{tag: "const", sym: comp.sym, expr: expression{tag: "lam", prms: comp.prms, body: comp.body}}, ce)
    }
}

func compile(comp interface{}, ce environment) {
    compile_comp[comp.(map[string]interface{})["tag"].(string)](comp, ce)
}

func compileProgram(program string) {
    wc := 0
    instrs := make([]instruction, 1000)

    compile(program, global_compile_environment, &wc, &instrs)
    instrs[wc] = instruction{tag: "DONE"}
}

func parseCompileRun(program string) interface{} {
    parseJSON := parseToJSON(program)
    compileProgram(parseJSON)
    return run()
}

func run() interface{} {
    OS := make([]interface{}, 0)
    PC := 0
    E := global_environment
    RTS := make([]interface{}, 0)

    for instrs[PC].tag != "DONE" {
        instr := instrs[PC]
        microcode[instr.tag](instr, &OS, &PC, &E, &RTS)
    }
    return peek(OS, 0)
}
*/


/*
scheduler sees the ids of each thread and decides to context switch
scheduler is inside the VM
threads are functions and there exists a keyword in the array instruction that decides to create new thread (needs to get parsed and compiled)
*/