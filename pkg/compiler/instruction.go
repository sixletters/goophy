package compiler

type Instruction interface {
	getTag() string
	// instructionNode()
}

type Instructions struct {
	instrs []Instruction
}

// LDC ->Boolean
type LDCBInstruction struct {
	tag string
	val bool
}

func (ldcb *LDCBInstruction) getTag() string {
	return ldcb.tag
} 

func (ldcb *LDCBInstruction) getValue() bool {
	return ldcb.val
} 

// LDC ->Integer TODO: Expand to handle other types like float64
type LDCNInstruction struct {
	tag string
	val int
}
func (ldcn *LDCNInstruction) getTag() string {
	return ldcn.tag
}

func (ldcn *LDCNInstruction) getValue() int {
	return ldcn.val
} 

type LDSInstruction struct {
	tag string
	sym string
}

func (lds *LDSInstruction) getTag() string {
	return lds.tag
}

func (lds *LDSInstruction) getSym() string {
	return lds.sym
} 

type DONEInstruction struct {
	tag string
}

func (done *DONEInstruction) getTag() string {
	return done.tag
}

type UNOPS string

const(
	negative  UNOPS = "-unary"
	not UNOPS  = "!"
)

type UNOPSInstruction struct {
	tag string
	sym UNOPS
}

func (unop *UNOPSInstruction) getTag() string {
	return unop.tag
}

func (unop *UNOPSInstruction) getSym() UNOPS {
	return unop.sym
}
type BINOPS string
const (
	add      BINOPS = "+"
	multiply BINOPS = "*"
	minus    BINOPS = "-"
	divide   BINOPS = "/"
	modulo   BINOPS = "%"
	lt       BINOPS = "<"
	le       BINOPS = "<="
	ge       BINOPS = ">="
	gt       BINOPS = ">"
	eq       BINOPS = "==="
	neq      BINOPS = "!=="
)

type BINOPSInstruction struct {
	tag string
	sym BINOPS
}

func (binop *BINOPSInstruction) getTag() string {
	return binop.tag
}

func (binop *BINOPSInstruction) getSym() BINOPS {
	return binop.sym
}

type POPInstruction struct {
	tag string
}
func (pop *POPInstruction) getTag() string {
	return pop.tag
}