package compiler

// Following are tags for the instructions on the instruction array
// type string string

// type string struct {
// 	tag string
// 	// Literal string
// }

const (
	//SVMLa
	DONE string = "DONE" // last instruction

	// LDC string = "LDC" // LDC i/b
	LDCB string = "LDCB"
	LDCN string = "LDCN"
	/*
		LDCN  Symbol = "LDCN"  // i is a number
		LDCB  Symbol = "LDCB"  // b is a boolean*/

	UNOP  string = "UNOP"  // unary operations
	BINOP string = "BINOP" // binary operations
	POP   string = "POP"
	/*
		PLUS  Symbol = "PLUS"  // addition
		MINUS Symbol = "MINUS" // subtraction
		TIMES Symbol = "TIMES" // multiplication
		AND   Symbol = "AND"   // conjunction
		OR    Symbol = "OR"    // disjunction
		NOT   Symbol = "NOT"   // negation
		LT    Symbol = "LT"    // less-than
		GT    Symbol = "GT"    // greater-than
		EQ    Symbol = "EQ"    // equality

		//SVMLb
		DIV string = "DIV" // divide (FOR NOW IGNORE divide by 0 COS SVML just evaluates to Infinity)
	*/
	//SVMLc
	JOF  string = "JOF"  // Jump On False absolute (alternatively we can do relative)
	GOTO string = "GOTO" // GOTO absolute

	//SVMLd
	ENTER_SCOPE string = "ENTER_SCOPE" // enters blk
	EXIT_SCOPE  string = "EXIT_SCOPE"  //exit blk
	LD          string = "LD"          // LoaD Symbolic
	ASSIGN      string = "ASSIGN"      // let/assignment/const
	LDF         string = "LDF"         // LoaD Function Symbolic
	CALL        string = "CALL"        // finds operands of application in reverse followed by operator, pushes to RTS a frame containing it's addr/env/params
	TAIL_CALL   string = "TAIL_CALL"   // does same as call except doesnt push to RTS
	RESET       string = "RESET"       // resets to caller
)
