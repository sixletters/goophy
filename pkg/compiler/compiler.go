package compiler

import (
	"cs4215/goophy/pkg/ast"
	"strconv"
)

type Compiler struct {
	wc int
}

func NewCompiler() *Compiler {
	return &Compiler{
		wc: 0,
	}
}

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
func (c *Compiler) Compile_statement(statement ast.Statement, instrs []Instruction) []Instruction {
	token := statement.GetToken()
	// fmt.Println(token.Type)
	// fmt.Println(token.Literal)
	switch token.Type {
	case "LET":
		assignStatement := statement.(*ast.LetStatement)
		newInstrs := c.Compile_expression(assignStatement.Value, instrs)
		instrs = append(instrs, newInstrs...)
		c.wc += 1
		assignInstruction := ASSIGNInstruction{Tag: "ASSIGN", Sym: assignStatement.Name.Value}
		instrs = append(instrs, assignInstruction)
	case "GO":
		goStatement := statement.(*ast.GoStatement)
		newInstrs := c.Compile_expression(goStatement.FunctionCall, instrs)
		instrs = append(instrs, GOInstruction{Tag: "GO"})
		instrs = append(instrs, newInstrs...)
		c.wc += 1
	case "RETURN":
		returnStatement := statement.(*ast.ReturnStatement)
		newInstrs := c.Compile_expression(returnStatement.ReturnValue, instrs)
		instrs = append(instrs, newInstrs...)
		// i need to change call into tailcall if its in the return statement
		/*returnStatement, ok := returnStatement.ReturnValue.(*ast.CallStatement)
		// I havent even do a normal call expresison yet so wait
		if ok {
		} else {*/
		c.wc += 1
		resetInstruction := RESETInstruction{Tag: "RESET"}
		instrs = append(instrs, resetInstruction)
	case "{":
		//block is broken, i not sure how to fix the compile statement step in the loop for statements
		blkStatement := statement.(*ast.BlockStatement)
		locals := scan(blkStatement.Statements)
		c.wc += 1
		enterScopeInstruction := ENTERSCOPEInstruction{Tag: "ENTER_SCOPE", Syms: locals}
		instrs = append(instrs, enterScopeInstruction)
		for _, statement := range blkStatement.Statements {
			// fmt.Println(len(blkStatement.Statements))
			newInstrs := c.Compile_statement(statement, []Instruction{})
			instrs = append(instrs, newInstrs...)
		}
		c.wc += 1
		exitScopeInstruction := EXITSCOPEInstruction{Tag: "EXIT_SCOPE"}
		instrs = append(instrs, exitScopeInstruction)
	default:
		expressionStatement := statement.(*ast.ExpressionStatement)
		newInstrs := c.Compile_expression(expressionStatement.Expression, instrs)
		instrs = append(instrs, newInstrs...)
	}
	return instrs
}

// WIP
func (c *Compiler) Compile_expression(expression ast.Expression, instrs []Instruction) []Instruction {
	token := expression.GetToken()
	// fmt.Println(token.Type)
	// fmt.Println(token.Literal)
	switch token.Type {
	case "ILLEGAL":
		panic("ILLEGAL EXPRESSION ENCOUNTERED")
	case "EOF":
		c.wc += 1
		doneInstruction := DONEInstruction{Tag: "DONE"}
		instrs = append(instrs, doneInstruction)
	case "IDENT":
		c.wc += 1
		ldsInstruction := LDSInstruction{Tag: "LDS", Sym: token.Literal}
		instrs = append(instrs, ldsInstruction)
	case "TRUE", "FALSE":
		val, _ := strconv.ParseBool(token.Literal)
		c.wc += 1
		ldcbInstruction := LDCBInstruction{Tag: "LDCB", Val: val}
		instrs = append(instrs, ldcbInstruction)
	case "INT":
		val, _ := strconv.Atoi(token.Literal)
		c.wc += 1
		ldcnInstruction := LDCNInstruction{Tag: "LDCN", Val: val}
		instrs = append(instrs, ldcnInstruction)
	case "+", "*", "/", "<=", ">", "==", "!=": /*tokens have not included modulo*/
		expr := expression.(*ast.InfixExpression)
		newInstrs := c.Compile_expression(expr.Left, []Instruction{})
		instrs = append(instrs, newInstrs...)
		newerInstrs := c.Compile_expression(expr.Right, []Instruction{})
		instrs = append(instrs, newerInstrs...)
		c.wc += 1
		binopInstruction := BINOPInstruction{Tag: "BINOP", Sym: BINOPS(token.Literal)}
		instrs = append(instrs, binopInstruction)
	case "-":
		// both negative number and binop minus is handled here
		expr, ok := expression.(*ast.PrefixExpression)
		if ok {
			newerInstrs := c.Compile_expression(expr.Right, []Instruction{})
			instrs = append(instrs, newerInstrs...)
			c.wc += 1
			unopInstruction := UNOPInstruction{Tag: "UNOP", Sym: "-unary"}
			instrs = append(instrs, unopInstruction)
		} else {
			expr := expression.(*ast.InfixExpression)
			newInstrs := c.Compile_expression(expr.Left, []Instruction{})
			instrs = append(instrs, newInstrs...)
			newerInstrs := c.Compile_expression(expr.Right, []Instruction{})
			instrs = append(instrs, newerInstrs...)
			c.wc += 1
			binopInstruction := BINOPInstruction{Tag: "BINOP", Sym: BINOPS(token.Literal)}
			instrs = append(instrs, binopInstruction)
		}
	case "FUNCTION":
		functionLiteral := expression.(*ast.FunctionLiteral)
		c.wc += 1
		parametersToString := []string{}
		for _, parameter := range functionLiteral.Parameters {
			parametersToString = append(parametersToString, parameter.Value)
		}
		ldfInstruction := LDFInstruction{Tag: "LDF", Prms: parametersToString, Addr: c.wc + 1}
		instrs = append(instrs, ldfInstruction)
		bodyInstrs := c.Compile_statement(functionLiteral.Body, []Instruction{})
		// WE SKIP an LDC instruction here because only in JS lambda returns undefined
		c.wc += 1
		resetInstruction := RESETInstruction{Tag: "RESET"}
		c.wc += 1
		gotoInstruction := GOTOInstruction{Tag: "GOTO", Addr: c.wc}
		instrs = append(instrs, gotoInstruction)
		instrs = append(instrs, bodyInstrs...)
		instrs = append(instrs, resetInstruction)
	case "(":
		callExpression := expression.(*ast.CallExpression)
		functionInstrs := c.Compile_expression(callExpression.Function, []Instruction{})
		instrs = append(instrs, functionInstrs...)
		var length = 0
		for _, arg := range callExpression.Arguments {
			argumentInstrs := c.Compile_expression(arg, []Instruction{})
			instrs = append(instrs, argumentInstrs...)
			length += 1
		}
		c.wc += 1
		callInstruction := CALLInstruction{Tag: "CALL", Arity: length}
		instrs = append(instrs, callInstruction)
	case "IF":
		ifExpression := expression.(*ast.IfExpression)
		condInstrs := c.Compile_expression(ifExpression.Condition, []Instruction{})
		instrs = append(instrs, condInstrs...)
		ifBlockInstrs := c.Compile_statement(ifExpression.IfBlock, []Instruction{})
		c.wc += 1
		jofInstruction := JOFInstruction{Tag: "JOF", Addr: c.wc}
		var elseBlockInstrs []Instruction
		var gotoInstruction GOTOInstruction
		if ifExpression.ElseBlock != nil {
			elseBlockInstrs = c.Compile_statement(ifExpression.ElseBlock, []Instruction{})
			c.wc += 1
			gotoInstruction = GOTOInstruction{Tag: "GOTO", Addr: c.wc}
		}
		instrs = append(instrs, jofInstruction)
		instrs = append(instrs, ifBlockInstrs...)
		if ifExpression.ElseBlock != nil {
			instrs = append(instrs, gotoInstruction)
			instrs = append(instrs, elseBlockInstrs...)
		}
	}
	return instrs
}

func (c *Compiler) Compile(program ast.Program) []Instruction {
	instrs := make([]Instruction, 0)
	locals := scan(program.Statements)
	c.wc += 1
	enterScopeInstruction := ENTERSCOPEInstruction{Tag: "ENTER_SCOPE", Syms: locals}
	instrs = append(instrs, enterScopeInstruction)
	for i, statement := range program.Statements {
		instrs = append(instrs, c.Compile_statement(statement, []Instruction{})...)
		if i != len(program.Statements)-1 {
			c.wc += 1
			popInstruction := POPInstruction{Tag: "POP"}
			instrs = append(instrs, popInstruction)
		}
	}
	c.wc += 1
	exitScopeInstruction := EXITSCOPEInstruction{Tag: "EXIT_SCOPE"}
	instrs = append(instrs, exitScopeInstruction)
	c.wc += 1
	doneInstruction := DONEInstruction{Tag: "DONE"}
	instrs = append(instrs, doneInstruction)
	return instrs
}
