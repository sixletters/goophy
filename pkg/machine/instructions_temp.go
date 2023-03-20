package machine

type Instruction interface {
	getTag() string
	// instructionNode()
}

type Instructions struct {
	instrs []Instruction
}

func (instrs Instructions) getInstrs() []Instruction {
	return instrs.instrs
}

func (instrs Instructions) addInstrs(instruction Instruction, wc int) {
	instrs.instrs = append(instrs.instrs, instruction)
}

// LDC ->Boolean
type LDCBInstruction struct {
	tag string
	val bool
}

func (ldcb LDCBInstruction) getTag() string {
	return ldcb.tag
}

func (ldcb LDCBInstruction) getValue() bool {
	return ldcb.val
}

// LDC ->Integer TODO: Expand to handle other types like float64
type LDCNInstruction struct {
	tag string
	val int
}

func (ldcn LDCNInstruction) getTag() string {
	return ldcn.tag
}

func (ldcn LDCNInstruction) getValue() int {
	return ldcn.val
}

// LD Symbolic
type LDSInstruction struct {
	tag string
	sym string
}

func (lds LDSInstruction) getTag() string {
	return lds.tag
}

func (lds LDSInstruction) getSym() string {
	return lds.sym
}

type DONEInstruction struct {
	tag string
}

func (done DONEInstruction) getTag() string {
	return done.tag
}

type UNOPS string

const (
	negative UNOPS = "-unary"
	not      UNOPS = "!"
)

type UNOPInstruction struct {
	tag string
	sym UNOPS
}

func (unop UNOPInstruction) getTag() string {
	return unop.tag
}

func (unop UNOPInstruction) getSym() UNOPS {
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

type BINOPInstruction struct {
	tag string
	sym BINOPS
}

func (binop BINOPInstruction) getTag() string {
	return binop.tag
}

func (binop BINOPInstruction) getSym() BINOPS {
	return binop.sym
}

type POPInstruction struct {
	tag string
}

func (pop POPInstruction) getTag() string {
	return pop.tag
}

type JOFInstruction struct {
	tag  string
	addr int
}

func (jof JOFInstruction) getTag() string {
	return jof.tag
}
func (jof JOFInstruction) getAddr() int {
	return jof.addr
}

// GOTO Absolute
type GOTOInstruction struct {
	tag  string
	addr int
}

func (gotoa GOTOInstruction) getTag() string {
	return gotoa.tag
}

func (gotoa GOTOInstruction) getAddr() int {
	return gotoa.addr
}

// Enter Scope
type ENTERSCOPEInstruction struct {
	tag  string
	syms []string
}

func (enterScope ENTERSCOPEInstruction) getTag() string {
	return enterScope.tag
}

func (enterScope ENTERSCOPEInstruction) getSyms() []string {
	return enterScope.syms
}

// Exit Scope
type EXITSCOPEInstruction struct {
	tag string
}

func (exitScope EXITSCOPEInstruction) getTag() string {
	return exitScope.tag
}

// LoaD Function
type LDFInstruction struct {
	tag  string
	prms []string
	addr int
}

func (ldf LDFInstruction) getTag() string {
	return ldf.tag
}

func (ldf LDFInstruction) getPrms() []string {
	return ldf.prms
}

func (ldf LDFInstruction) getAddr() int {
	return ldf.addr
}

// CALL
type CALLInstruction struct {
	tag   string
	arity int
}

func (call CALLInstruction) getTag() string {
	return call.tag
}

func (call CALLInstruction) getArity() int {
	return call.arity
}

// TAIL_CALL special case where we dont push onto RTS
type TAILCALLInstruction struct {
	tag   string
	arity int
}

func (tailCall TAILCALLInstruction) getTag() string {
	return tailCall.tag
}

func (tailCall TAILCALLInstruction) getArity() int {
	return tailCall.arity
}

// RESET
type RESETInstruction struct {
	tag string
}

func (reset RESETInstruction) getTag() string {
	return reset.tag
}
