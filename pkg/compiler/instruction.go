package compiler

type Instruction interface {
	GetTag() string
	// instructionNode()
}

// LDC ->Boolean
type LDCBInstruction struct {
	tag string
	Val bool
}

func (ldcb LDCBInstruction) GetTag() string {
	return ldcb.tag
}

func (ldcb LDCBInstruction) getValue() bool {
	return ldcb.Val
}

// LDC ->Integer TODO: Expand to handle other types like float64
type LDCNInstruction struct {
	tag string
	Val int
}

func (ldcn LDCNInstruction) GetTag() string {
	return ldcn.tag
}

func (ldcn LDCNInstruction) getValue() int {
	return ldcn.Val
}

// LD Symbolic
type LDSInstruction struct {
	tag string
	Sym string
}

func (lds LDSInstruction) GetTag() string {
	return lds.tag
}

func (lds LDSInstruction) GetSym() string {
	return lds.Sym
}

type DONEInstruction struct {
	tag string
}

func (done DONEInstruction) GetTag() string {
	return done.tag
}

type UNOPS string

const (
	negative UNOPS = "-unary"
	not      UNOPS = "!"
)

type UNOPInstruction struct {
	tag string
	Sym UNOPS
}

func (unop UNOPInstruction) GetTag() string {
	return unop.tag
}

func (unop UNOPInstruction) getSym() UNOPS {
	return unop.Sym
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

type BINOPInstruction struct {
	tag string
	Sym BINOPS
}

func (binop BINOPInstruction) GetTag() string {
	return binop.tag
}

func (binop BINOPInstruction) GetSym() BINOPS {
	return binop.Sym
}

type POPInstruction struct {
	tag string
}

func (pop POPInstruction) GetTag() string {
	return pop.tag
}

type JOFInstruction struct {
	tag  string
	addr int
}

func (jof JOFInstruction) GetTag() string {
	return jof.tag
}
func (jof JOFInstruction) GetAddr() int {
	return jof.addr
}

// GOTO Absolute
type GOTOInstruction struct {
	tag  string
	addr int
}

func (gotoa GOTOInstruction) GetTag() string {
	return gotoa.tag
}

func (gotoa GOTOInstruction) GetAddr() int {
	return gotoa.addr
}

// Enter Scope
type ENTERSCOPEInstruction struct {
	tag  string
	syms []string
}

func (enterScope ENTERSCOPEInstruction) GetTag() string {
	return enterScope.tag
}

func (enterScope ENTERSCOPEInstruction) GetSyms() []string {
	return enterScope.syms
}

// Exit Scope
type EXITSCOPEInstruction struct {
	tag string
}

func (exitScope EXITSCOPEInstruction) GetTag() string {
	return exitScope.tag
}

type ASSIGNInstruction struct {
	tag string
	sym string
}

func (assign ASSIGNInstruction) GetTag() string {
	return assign.tag
}

func (assign ASSIGNInstruction) GetSym() string {
	return assign.sym
}

// LoaD Function
type LDFInstruction struct {
	tag  string
	prms []string
	addr int
}

func (ldf LDFInstruction) GetTag() string {
	return ldf.tag
}

func (ldf LDFInstruction) GetPrms() []string {
	return ldf.prms
}

func (ldf LDFInstruction) GetAddr() int {
	return ldf.addr
}

// CALL
type CALLInstruction struct {
	tag   string
	arity int
}

func (call CALLInstruction) GetTag() string {
	return call.tag
}

func (call CALLInstruction) GetArity() int {
	return call.arity
}

// TAIL_CALL special case where we dont push onto RTS
type TAILCALLInstruction struct {
	tag   string
	arity int
}

func (tailCall TAILCALLInstruction) GetTag() string {
	return tailCall.tag
}

func (tailCall TAILCALLInstruction) GetArity() int {
	return tailCall.arity
}

// RESET
type RESETInstruction struct {
	tag string
}

func (reset RESETInstruction) GetTag() string {
	return reset.tag
}
