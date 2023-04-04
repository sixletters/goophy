// Idealized Virtual Machine

package machine

// "errors"
// "encoding/json" // To convert array in string form back into array
// "cs4215/goophy/pkg/environment"
import (
	"cs4215/goophy/pkg/compiler"
	"cs4215/goophy/pkg/environment"
	"cs4215/goophy/pkg/scheduler"
	"cs4215/goophy/pkg/util"
	"fmt"
	"reflect"
)

// Parser
// Microcode
// Scheduler

// const global_environment = heap
// RTS, OS, PC, E
// var RTS Stack
// var OS Stack
// var E *Environment
// var PC int

type Machine struct {
	Rrs         *scheduler.RoundRobinScheduler
	OS          util.Stack
	RTS         util.Stack
	E           *environment.Environment
	spawnThread bool
	PC          int
	microcode   map[string]func(instr compiler.Instruction)
	binop       map[string]func(x, y interface{}) interface{}
	unop        map[string]func(interface{}) interface{}
}

// Closures to wait for compiler
type closure struct {
	tag  string
	prms []string //an array of symbols
	addr int
	env  *environment.Environment
}
type builtinType struct {
	tag   string
	sym   string
	arity int
}

func NewMachine() *Machine {
	global_environment := environment.NewEnvironment() //TODO: Populate global environment with all built-ins
	for key, fn := range builtin_mapping {
		fnType := reflect.TypeOf(fn)
		value := builtinType{
			tag:   "BUILTIN",
			sym:   key,
			arity: fnType.NumIn(),
		}
		global_environment.Set_declare(key, value)
	}
	// fmt.Println(global_environment.Get("print"))
	// apply_builtin("print", "1+2")
	global_environment.Extend() //Redundant?
	return &Machine{
		Rrs:         scheduler.NewRoundRobinScheduler(),
		OS:          util.Stack{},
		PC:          0,
		E:           global_environment,
		RTS:         util.Stack{},
		spawnThread: false,
	}
}

func (m *Machine) Init() *Machine {
	mainThread := scheduler.Thread{
		Os:  m.OS,
		Env: m.E,
		Rts: m.RTS,
		Pc:  m.PC,
	}
	m.Rrs.NewThread(mainThread)
	m.unop = map[string]func(interface{}) interface{}{
		"-unary": func(x interface{}) interface{} {
			switch x := x.(type) {
			case int:
				return -x
			case float64:
				return -x
			default:
				return nil // return an error message or value here
			}
		},
		"!": func(x interface{}) interface{} {
			if ok := isBool(x); ok {
				return !x.(bool)
			}
			return nil // return an error message or value here
		},
	}
	m.binop = map[string]func(x, y interface{}) interface{}{
		"=": func(x, y interface{}) interface{} {
			if _, ok := y.(string); !ok {
				panic("Symbol is not a string type!")
			}

			if _, ok := m.E.Get(y.(string)); !ok {
				panic("Symbol not found!")
			}
			m.E.Set_assign(y.(string), x)
			return y
		},
		"+": func(x, y interface{}) interface{} {
			if x == nil || y == nil {
				return nil
			}
			if isNumber(x) && isNumber(y) {
				switch x.(type) {
				case int:
					return x.(int) + y.(int)
				case float64:
					return x.(float64) + y.(float64)
				}
			} else if isString(x) && isString(y) {
				return x.(string) + y.(string)
			}
			return "error: + expects two numbers or two strings, got:" +
				reflect.TypeOf(x).String() + " and " +
				reflect.TypeOf(y).String()
		},

		"*": func(x, y interface{}) interface{} {
			// return x.(float64) * y.(float64)
			switch x.(type) {
			case int:
				return x.(int) * y.(int)
			case float64:
				return x.(float64) * y.(float64)
			default:
				panic("invalid type")
			}
		},
		"-": func(x, y interface{}) interface{} {
			switch x.(type) {
			case int:
				return x.(int) - y.(int)
			case float64:
				return x.(float64) - y.(float64)
			default:
				panic("invalid type")
			}
		},
		"/": func(x, y interface{}) interface{} {
			switch x.(type) {
			case int:
				return x.(int) / y.(int)
			case float64:
				return x.(float64) / y.(float64)
			default:
				panic("invalid type")
			}
		},
		"%": func(x, y interface{}) interface{} {
			switch x.(type) {
			case int:
				return x.(int) % y.(int)
			default:
				panic("invalid type")
			}
		},
		"<": func(y, x interface{}) interface{} {
			switch x.(type) {
			case int:
				return x.(int) < y.(int)
			case float64:
				return x.(float64) < y.(float64)
			default:
				panic("invalid type")
			}
		},
		"<=": func(y, x interface{}) interface{} {
			switch x.(type) {
			case int:
				return x.(int) <= y.(int)
			case float64:
				return x.(float64) <= y.(float64)
			default:
				panic("invalid type")
			}
		},
		">=": func(y, x interface{}) interface{} {
			switch x.(type) {
			case int:
				return x.(int) >= y.(int)
			case float64:
				return x.(float64) >= y.(float64)
			default:
				panic("invalid type")
			}
		},
		">": func(y, x interface{}) interface{} {
			switch x.(type) {
			case int:
				fmt.Println(x.(int) > y.(int))
				return x.(int) > y.(int)
			case float64:
				return x.(float64) > y.(float64)
			default:
				panic("invalid type")
			}
		},
		"==": func(x, y interface{}) interface{} {
			return x == y
		},
		"!=": func(x, y interface{}) interface{} {
			return x != y
		},
	}
	m.microcode = map[string]func(instr compiler.Instruction){
		"LDCN": func(instr compiler.Instruction) {
			ldcnInstr, ok := instr.(compiler.LDCNInstruction)
			if !ok {
				panic("instr is not of type LDCNInstruction")
			}
			m.PC++
			m.OS.Push(ldcnInstr.Val)
		},
		"LDCB": func(instr compiler.Instruction) {
			ldcbInstr, ok := instr.(compiler.LDCBInstruction)
			if !ok {
				panic("instr is not of type LDCBInstruction")
			}
			m.PC++
			m.OS.Push(ldcbInstr.Val)
		},
		"LDI": func(instr compiler.Instruction) {
			ldiInstr, ok := instr.(compiler.LDIInstruction)
			if !ok {
				panic("instr is not of type LDIInstruction")
			}
			m.PC++
			m.OS.Push(ldiInstr.Val)
		},
		"UNOP": func(instr compiler.Instruction) {
			unopInstr, ok := instr.(compiler.UNOPInstruction)
			if !ok {
				panic("instr is not of type UNOPInstruction")
			}
			m.PC++
			m.OS.Push(m.unop[string(unopInstr.Sym)](m.OS.Pop()))
		},
		"BINOP": func(instr compiler.Instruction) {
			binopInstr, ok := instr.(compiler.BINOPInstruction)
			if !ok {
				panic("instr is not of type BINOPInstruction")
			}
			m.PC++
			m.OS.Push(m.binop[string(binopInstr.Sym)](m.OS.Pop(), m.OS.Pop()))
		},
		"POP": func(instr compiler.Instruction) {
			m.PC++
			m.OS.Pop()
		},
		"LDS": func(instr compiler.Instruction) {
			ldsInstr, ok := instr.(compiler.LDSInstruction)
			if !ok {
				panic("instr is not of type LDSInstruction")
			}
			m.PC++
			val, ok := m.E.Get(ldsInstr.GetSym())
			if !ok || val == nil {
				panic(fmt.Sprintf("symbol %s not found", ldsInstr.GetSym()))
			}
			m.OS.Push(val) //Note this pushes interface{} type into OS
		},
		// throw error if cannot find value
		"ASSIGN": func(instr compiler.Instruction) {
			assignInstr, ok := instr.(compiler.ASSIGNInstruction)
			if !ok {
				panic("instr is not of type ASSIGNInstruction")
			}
			m.PC++
			m.E.Set_declare(assignInstr.GetSym(), m.OS.Peek())
			// assign_value(assignInstr.GetSym(), OS.Peek(), &E)
		},
		"ENTER_SCOPE": func(instr compiler.Instruction) {
			m.PC++
			enterscopeInstr, ok := instr.(compiler.ENTERSCOPEInstruction)
			if !ok {
				panic("instr is not of type BINOPInstruction")
			}
			m.RTS.Push(stackFrame{tag: "BLOCK_FRAME", E: m.E}) //TODO: Include RTS back for OS, PC
			locals := enterscopeInstr.GetSyms()
			m.E = m.E.Extend()
			for _, i := range locals {
				m.E.Set_declare(i, environment.Unassigned)
			}
			// unassigneds := make([]interface{}, len(locals)) //TODO: Change the type to String?
			// for i := range unassigneds {
			// 	unassigneds[i] = unassigned
			// }
			// E.Extend()
			// E.Set()
		},
		"EXIT_SCOPE": func(instr compiler.Instruction) {
			m.PC++
			sf, ok := m.RTS.Pop().(stackFrame)
			if !ok {
				panic("Frame on RTS is not of type stackFrame")
			}
			m.E = sf.E
		},
		"LDF": func(instr compiler.Instruction) {
			m.PC++
			ldfInstr, ok := instr.(compiler.LDFInstruction)
			if !ok {
				panic("instr is not of type LDFInstruction")
			}
			closure_var := closure{
				tag:  ldfInstr.GetTag(),
				prms: ldfInstr.GetPrms(),
				addr: ldfInstr.GetAddr(),
				env:  m.E,
			}
			m.OS.Push(closure_var)
		},
		"JOF": func(instr compiler.Instruction) {
			cond := m.OS.Pop()
			jofInstr, ok := instr.(compiler.JOFInstruction)
			if !ok {
				panic("instr is not of type JOFInstruction")
			}
			if isTruthy(cond) {
				// fmt.Println("x is truthy")
				m.PC = m.PC + 1
			} else {
				// fmt.Println("x is falsy")
				m.PC = jofInstr.GetAddr()
			}
		},
		"GOTO": func(instr compiler.Instruction) {
			gotoInstr, ok := instr.(compiler.GOTOInstruction)
			if !ok {
				panic("instr is not of type GOTOInstruction")
			}
			// fmt.Println(instr)
			m.PC = gotoInstr.GetAddr()
		},
		"CALL": func(instr compiler.Instruction) {
			callInstr, ok := instr.(compiler.CALLInstruction)
			if !ok {
				panic("instr is not of type CALLInstruction")
			}
			arity := callInstr.GetArity()
			args := make([]interface{}, arity)
			for i := arity - 1; i >= 0; i-- {
				args[i] = m.OS.Pop()
			}
			sf := m.OS.Pop() //Can either be closure or builtin
			sf_closure, ok := sf.(closure)
			if ok {
				m.RTS.Push(stackFrame{tag: "CALL_FRAME", E: m.E, PC: m.PC + 1})
				m.E.Extend()
				for index, val := range sf_closure.prms {
					m.E.Set_declare(val, args[index])
				}
				if m.spawnThread {
					m.Rrs.NewThread(
						scheduler.Thread{
							Env: m.E,
							Rts: util.Stack{},
							Pc:  sf_closure.addr,
							Os:  util.Stack{},
						},
					)
					m.spawnThread = false
					sf := m.RTS.Pop().(stackFrame)
					m.PC = sf.PC
					m.E = sf.E
					return
				}
				m.PC = sf_closure.addr
				return
			}
			sf_builtin, ok := sf.(builtinType)
			if ok {
				m.PC++
				result, ok := builtin_mapping[sf_builtin.sym](args)
				if ok != nil {
					m.OS.Push(result)
				}
				if m.spawnThread {
					m.spawnThread = false
				}
				return
			}
		},
		"TAIL_CALL": func(instr compiler.Instruction) {
			tailcallInstr, ok := instr.(compiler.TAILCALLInstruction)
			if !ok {
				panic("instr is not of type TAILCALLInstruction")
			}
			arity := tailcallInstr.GetArity()
			args := make([]interface{}, arity)
			for i := arity - 1; i >= 0; i-- {
				args[i] = m.OS.Pop()
			}
			sf := m.OS.Pop() //Can either be closure or builtin
			sf_closure, ok := sf.(closure)
			if ok {
				m.E.Extend()
				for index, val := range sf_closure.prms {
					m.E.Set_declare(val, args[index])
				}
				m.PC = sf_closure.addr
				return
			}
			sf_builtin, ok := sf.(builtinType)
			if ok {
				m.PC++
				result, ok := builtin_mapping[sf_builtin.sym](args)
				if ok != nil {
					m.OS.Push(result)
				}
				return
			}

		},
		"RESET": func(instr compiler.Instruction) {
			top_frame, ok := m.RTS.Pop().(stackFrame)
			if !ok {
				panic("Frame on RTS is not of type stackFrame")
			}
			if top_frame.tag == "CALL_FRAME" {
				m.PC = top_frame.PC
				m.E = top_frame.E
			}
		},
		"GO": func(instr compiler.Instruction) {
			m.spawnThread = true
			m.PC += 1
		},
	}
	return m
}

// TODO: Check allowable types for return
// func run(instrs) interface{} {
func (m *Machine) Run(instrs []compiler.Instruction) interface{} {
	for len(m.Rrs.GetCurrentThreads()) != 0 {
		count := 0
		m.contextSwitch()
		for instrs[m.PC].GetTag() != "DONE" {
			// fmt.Print(m.Rrs.GetCurrentThreads())
			// fmt.Printf(" %d ", m.Rrs.GetCurrThreadID())
			// fmt.Print(instrs[m.PC])
			// fmt.Println("")
			count += 1
			m.microcode[instrs[m.PC].GetTag()](instrs[m.PC])
			// When the thread is done.
			if instrs[m.PC].GetTag() == "RESET" && m.RTS.Size() == 0 {
				break
			}
			if instrs[m.PC].GetTag() == "GO" {
				count = count - 2
			}
			// context switch
			if count > 1 {
				// fmt.Print(m.E.Get("print"))
				m.saveContext()
				break
			}
		}

	}
	return m.OS.Peek()
}

func (m *Machine) contextSwitch() {
	threadId, _ := m.Rrs.GetNextThread()
	m.OS = m.Rrs.ThreadTable[threadId].Os
	m.RTS = m.Rrs.ThreadTable[threadId].Rts
	m.E = m.Rrs.ThreadTable[threadId].Env
	m.PC = m.Rrs.ThreadTable[threadId].Pc
}

func (m *Machine) saveContext() {
	m.Rrs.AddThread(
		scheduler.Thread{
			Os:  m.OS,
			Env: m.E,
			Pc:  m.PC,
			Rts: m.RTS,
		},
	)
}

// Helper for dynamically checking truthy in interface types
func isTruthy(x interface{}) bool {
	switch v := x.(type) {
	case bool:
		return v
	case string:
		return v != ""
	case int, int8, int16, int32, int64:
		return v != 0
	case uint, uint8, uint16, uint32, uint64:
		return v != 0
	case float32, float64:
		return v != 0.0
	case nil:
		return false
	default:
		return false
	}
}
