package compiler

// Following are tags for the instructions on the instruction array
type TagType string

type Tag struct {
	Type    TagType
	Literal string
}

func NewTag(tagType TagType, ch byte) Tag {
	return Tag{
		Type:    tagType,
		Literal: string(ch),
	}
}

func NewTagWithStr(tagType TagType, ch string) Tag {
	return Tag{
		Type:    tagType,
		Literal: ch,
	}
}

const (
	//SVMLa
	DONE TagType = "DONE" // last instruction

	LDC TagType = "LDC" // LDC i/b
	/*
		LDCN  Symbol = "LDCN"  // i is a number
		LDCB  Symbol = "LDCB"  // b is a boolean*/

	UNOP  TagType = "UNOP"  // unary operations
	BINOP TagType = "BINOP" // binary operations
	POP   TagType = "POP"
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
		DIV TagType = "DIV" // divide (FOR NOW IGNORE divide by 0 COS SVML just evaluates to Infinity)
	*/
	//SVMLc
	JOF  TagType = "JOF"  // Jump On False absolute (alternatively we can do relative)
	GOTO TagType = "GOTO" // GOTO absolute

	//SVMLd
	ENTER_SCOPE TagType = "ENTER_SCOPE" // enters blk
	EXIT_SCOPE  TagType = "EXIT_SCOPE"  //exit blk
	LD          TagType = "LD"          // LoaD Symbolic
	ASSIGN      TagType = "ASSIGN"      // let/assignment/const
	LDF         TagType = "LDF"         // LoaD Function Symbolic
	CALL        TagType = "CALL"        // finds operands of application in reverse followed by operator, pushes to RTS a frame containing it's addr/env/params
	TAIL_CALL   TagType = "TAIL_CALL"   // does same as call except doesnt push to RTS
	RESET       TagType = "RESET"       // resets to caller
)
