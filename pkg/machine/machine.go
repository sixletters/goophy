// Idealized Virtual Machine

package machine

// "errors"
// "encoding/json" // To convert array in string form back into array
// "cs4215/goophy/pkg/environment"
import (
	"cs4215/goophy/pkg/compiler"
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
	OS        Stack
	RTS       Stack
	E         *Environment
	PC        int
	microcode map[string]func(instr compiler.Instruction)
}

// Closures to wait for compiler
type closure struct {
	tag  string
	prms []string //an array of symbols
	addr int
	env  *Environment
}
type builtinType struct {
	tag   string
	sym   string
	arity int
}

func NewMachine() *Machine {
	global_environment := NewEnvironment() //TODO: Populate global environment with all built-ins
	for key, fn := range builtin_mapping {
		fnType := reflect.TypeOf(fn)
		value := builtinType{
			tag:   "BUILTIN",
			sym:   key,
			arity: fnType.NumIn(),
		}
		global_environment.Set(key, value)
	}
	// fmt.Println(global_environment.Get("print"))
	// apply_builtin("print", "1+2")
	// fmt.Print(global_environment)
	global_environment.Extend() //Redundant?
	return &Machine{
		OS:  Stack{},
		PC:  0,
		E:   global_environment,
		RTS: Stack{},
	}
}

func (m *Machine) Init() *Machine {
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
		"UNOP": func(instr compiler.Instruction) {
			unopInstr, ok := instr.(compiler.UNOPInstruction)
			if !ok {
				panic("instr is not of type UNOPInstruction")
			}
			m.PC++
			m.OS.Push(apply_unop(string(unopInstr.Sym), m.OS.Pop()))
		},
		"BINOP": func(instr compiler.Instruction) {
			binopInstr, ok := instr.(compiler.BINOPInstruction)
			if !ok {
				panic("instr is not of type BINOPInstruction")
			}
			m.PC++
			m.OS.Push(apply_binop(string(binopInstr.Sym), m.OS.Pop(), m.OS.Pop()))
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
			m.E.Set(assignInstr.GetSym(), m.OS.Peek())
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
				m.E.Set(i, unassigned)
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
					m.E.Set(val, args[index])
				}
				m.PC = sf_closure.addr
				return
			}
			sf_builtin, ok := sf.(builtinType)
			if ok {
				m.PC++
				result, ok := apply_builtin(sf_builtin.sym, args)
				if ok != nil {
					m.OS.Push(result)
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
					m.E.Set(val, args[index])
				}
				m.PC = sf_closure.addr
				return
			}
			sf_builtin, ok := sf.(builtinType)
			if ok {
				m.PC++
				result, ok := apply_builtin(sf_builtin.sym, args)
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
	}
	return m
}

// "LD": func(instr *compiler.Instruction) {
//     PC++
// 	OS.Push(lookup(instr.sym, E))
// },
// "CALL": func(instr *compiler.Instruction) {
//     arity := instr.arity
//     args := make([]interface{}, arity)
//     for i := arity - 1; i >= 0; i-- {
//         args[i] = OS.Pop()
//     }
//     sf := OS.pop().(*closure)
//     if sf.tag == "BUILTIN" {
//         PC++
//         push(OS, apply_builtin(sf.sym, args))
//         return
//     }
//     RTS.push(&frame{tag: "CALL_FRAME", addr: PC + 1, env: E})
//     E = extend(sf.prms, args, sf.env)
//     PC = sf.addr
// },
// "TAIL_CALL": func(instr *compiler.Instruction) {
//     arity := instr.arity
//     args := make([]interface{}, arity)
//     for i := arity - 1; i >= 0; i-- {
//         args[i] = OS.pop()
//     }
//     sf := OS.pop().(*closure)
//     if sf.tag == "BUILTIN" {
//         PC++
//         push(OS, apply_builtin(sf.sym, args))
//         return
//     }
//     E = extend(sf.prms, args, sf.env)
//     PC = sf.addr
// },
// "RESET": func(instr *compiler.Instruction) {
//     for {
//         top_frame := RTS.pop()
//         if top_frame.tag == "CALL_FRAME" {
//             PC = top_frame.addr
//             E = top_frame.env
//             break
//         }
//     }
// },
// }

//TODO: Finalize struct for instruction
// type compiler.Instruction struct{
// 	tag string
// 	val interface{}
// 	// sym string
// }

// TODO: Check allowable types for return
// func run(instrs) interface{} {
func (m *Machine) Run(instrs []compiler.Instruction) interface{} {
	for instrs[m.PC].GetTag() != "DONE" {
		// fmt.Println(instrs[m.PC])
		m.microcode[instrs[m.PC].GetTag()](instrs[m.PC])
	}
	return m.OS.Peek()
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
