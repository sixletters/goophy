package compiler

type Instruction interface {
	GetTag() string
}

// LDC ->Boolean
type LDCBInstruction struct {
	Tag string
	Val bool
}

func (ldcb LDCBInstruction) GetTag() string {
	return ldcb.Tag
}

func (ldcb LDCBInstruction) GetValue() bool {
	return ldcb.Val
}

// LDC ->Integer TODO: Expand to handle other types like float64
type LDCNInstruction struct {
	Tag string
	Val int
}

func (ldcn LDCNInstruction) GetTag() string {
	return ldcn.Tag
}

func (ldcn LDCNInstruction) GetValue() int {
	return ldcn.Val
}

// LDC ->Integer TODO: Expand to handle other types like float64
type LDIInstruction struct {
	Tag string
	Val string
}

func (ldci LDIInstruction) GetTag() string {
	return ldci.Tag
}

func (ldci LDIInstruction) GetValue() string {
	return ldci.Val
}

// LD Symbolic
type LDSInstruction struct {
	Tag string
	Sym string
}

func (lds LDSInstruction) GetTag() string {
	return lds.Tag
}

func (lds LDSInstruction) GetSym() string {
	return lds.Sym
}

type DONEInstruction struct {
	Tag string
}

func (done DONEInstruction) GetTag() string {
	return done.Tag
}

type UNOPS string

const (
	negative UNOPS = "-unary"
	not      UNOPS = "!"
)

type UNOPInstruction struct {
	Tag string
	Sym UNOPS
}

func (unop UNOPInstruction) GetTag() string {
	return unop.Tag
}

func (unop UNOPInstruction) GetSym() UNOPS {
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
	Tag string
	Sym BINOPS
}

func (binop BINOPInstruction) GetTag() string {
	return binop.Tag
}

func (binop BINOPInstruction) GetSym() BINOPS {
	return binop.Sym
}

type POPInstruction struct {
	Tag string
}

func (pop POPInstruction) GetTag() string {
	return pop.Tag
}

type JOFInstruction struct {
	Tag  string
	Addr int
}

func (jof JOFInstruction) GetTag() string {
	return jof.Tag
}
func (jof JOFInstruction) GetAddr() int {
	return jof.Addr
}

// GOTO Absolute
type GOTOInstruction struct {
	Tag  string
	Addr int
}

func (gotoa GOTOInstruction) GetTag() string {
	return gotoa.Tag
}

func (gotoa GOTOInstruction) GetAddr() int {
	return gotoa.Addr
}

// Enter Scope
type ENTERSCOPEInstruction struct {
	Tag  string
	Syms []string
}

func (enterScope ENTERSCOPEInstruction) GetTag() string {
	return enterScope.Tag
}

func (enterScope ENTERSCOPEInstruction) GetSyms() []string {
	return enterScope.Syms
}

// Exit Scope
type EXITSCOPEInstruction struct {
	Tag string
}

func (exitScope EXITSCOPEInstruction) GetTag() string {
	return exitScope.Tag
}

type ASSIGNInstruction struct {
	Tag string
	Sym string
}

func (assign ASSIGNInstruction) GetTag() string {
	return assign.Tag
}

func (assign ASSIGNInstruction) GetSym() string {
	return assign.Sym
}

// LoaD Function
type LDFInstruction struct {
	Tag  string
	Prms []string
	Addr int
}

func (ldf LDFInstruction) GetTag() string {
	return ldf.Tag
}

func (ldf LDFInstruction) GetPrms() []string {
	return ldf.Prms
}

func (ldf LDFInstruction) GetAddr() int {
	return ldf.Addr
}

// CALL
type CALLInstruction struct {
	Tag   string
	Arity int
}

func (call CALLInstruction) GetTag() string {
	return call.Tag
}

func (call CALLInstruction) GetArity() int {
	return call.Arity
}

// TAIL_CALL special case where we dont push onto RTS
type TAILCALLInstruction struct {
	Tag   string
	Arity int
}

func (tailCall TAILCALLInstruction) GetTag() string {
	return tailCall.Tag
}

func (tailCall TAILCALLInstruction) GetArity() int {
	return tailCall.Arity
}

// RESET
type RESETInstruction struct {
	Tag string
}

func (reset RESETInstruction) GetTag() string {
	return reset.Tag
}

type GOInstruction struct {
	Tag string
}

func (g GOInstruction) GetTag() string {
	return g.Tag
}
