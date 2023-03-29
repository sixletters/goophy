// Idealized Virtual Machine

package machine

// "errors"
// "encoding/json" // To convert array in string form back into array
// "cs4215/goophy/pkg/environment"
import (
	"cs4215/goophy/pkg/compiler"
	"fmt"
)

// Parser
// Microcode
// Scheduler

// const global_environment = heap
// RTS, OS, PC, E
var RTS Stack
var OS Stack
var E *Environment
var PC int

// Microcode
//
//	var microcode = map[string]func(compiler.Instruction){
//	    "LDC": func(instr compiler.Instruction) {
//	        PC++
//			OS.Push(instr.val)
//	    },

// Closures to wait for compiler
type closure struct {
	tag  string
	prms []string //an array of symbols
	addr int
	env  *Environment
}

// Helper for dynamically checking truthy in interface types
func IsTruthy(x interface{}) bool {
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

var microcode = map[string]func(instr compiler.Instruction){
	"LDCN": func(instr compiler.Instruction) {
		ldcnInstr, ok := instr.(compiler.LDCNInstruction)
		if !ok {
			panic("instr is not of type LDCNInstruction")
		}
		PC++
		OS.Push(ldcnInstr.Val)
	},
	"LDCB": func(instr compiler.Instruction) {
		ldcbInstr, ok := instr.(compiler.LDCBInstruction)
		if !ok {
			panic("instr is not of type LDCBInstruction")
		}
		PC++
		OS.Push(ldcbInstr.Val)
	},
	"UNOP": func(instr compiler.Instruction) {
		unopInstr, ok := instr.(compiler.UNOPInstruction)
		if !ok {
			panic("instr is not of type UNOPInstruction")
		}
		PC++
		OS.Push(apply_unop(string(unopInstr.Sym), OS.Pop()))
	},
	"BINOP": func(instr compiler.Instruction) {
		binopInstr, ok := instr.(compiler.BINOPInstruction)
		if !ok {
			panic("instr is not of type BINOPInstruction")
		}
		PC++
		OS.Push(apply_binop(string(binopInstr.Sym), OS.Pop(), OS.Pop()))
	},
	"POP": func(instr compiler.Instruction) {
		PC++
		OS.Pop()
	},
	"LDS": func(instr compiler.Instruction) {
		ldsInstr, ok := instr.(compiler.LDSInstruction)
		if !ok {
			panic("instr is not of type LDSInstruction")
		}
		PC++
		val, ok := E.Get(ldsInstr.GetSym())
		if !ok || val == nil {
			panic(fmt.Sprintf("symbol %s not found", ldsInstr.GetSym()))
		}
		OS.Push(val) //Note this pushes interface{} type into OS
	},
	// throw error if cannot find value
	"ASSIGN": func(instr compiler.Instruction) {
		assignInstr, ok := instr.(compiler.ASSIGNInstruction)
		if !ok {
			panic("instr is not of type ASSIGNInstruction")
		}
		PC++
		E.Set(assignInstr.GetSym(), OS.Peek())
		// assign_value(assignInstr.GetSym(), OS.Peek(), &E)
	},
	"ENTER_SCOPE": func(instr compiler.Instruction) {
		PC++
		enterscopeInstr, ok := instr.(compiler.ENTERSCOPEInstruction)
		if !ok {
			panic("instr is not of type BINOPInstruction")
		}
		RTS.Push(stackFrame{tag: "BLOCK_FRAME", E: E}) //TODO: Include RTS back for OS, PC
		locals := enterscopeInstr.GetSyms()
		E = E.Extend()
		for _, i := range locals {
			E.Set(i, unassigned)
		}
		// unassigneds := make([]interface{}, len(locals)) //TODO: Change the type to String?
		// for i := range unassigneds {
		// 	unassigneds[i] = unassigned
		// }
		// E.Extend()
		// E.Set()
	},
	"EXIT_SCOPE": func(instr compiler.Instruction) {
		PC++
		sf, ok := RTS.Pop().(stackFrame)
		if !ok {
			panic("Frame on RTS is not of type stackFrame")
		}
		E = sf.E
	},
	"LDF": func(instr compiler.Instruction) {
		PC++
		ldfInstr, ok := instr.(compiler.LDFInstruction)
		if !ok {
			panic("instr is not of type LDFInstruction")
		}
		closure_var := closure{
			tag:  ldfInstr.GetTag(),
			prms: ldfInstr.GetPrms(),
			addr: ldfInstr.GetAddr(),
			env:  E,
		}
		OS.Push(closure_var)
	},
	"JOF": func(instr compiler.Instruction) {
		cond := OS.Pop()
		jofInstr, ok := instr.(compiler.JOFInstruction)
		if !ok {
			panic("instr is not of type JOFInstruction")
		}
		if IsTruthy(cond) {
			// fmt.Println("x is truthy")
			PC = PC + 1
		} else {
			// fmt.Println("x is falsy")
			PC = jofInstr.GetAddr()
		}
	},
	"GOTO": func(instr compiler.Instruction) {
		gotoInstr, ok := instr.(compiler.GOTOInstruction)
		if !ok {
			panic("instr is not of type GOTOInstruction")
		}
		// fmt.Println(instr)
		PC = gotoInstr.GetAddr()
	},
	"CALL": func(instr compiler.Instruction) {
		callInstr, ok := instr.(compiler.CALLInstruction)
		if !ok {
			panic("instr is not of type CALLInstruction")
		}
		arity := callInstr.GetArity()
		args := make([]interface{}, arity)
		for i := arity - 1; i >= 0; i-- {
			args[i] = OS.Pop()
		}
		sf := OS.Pop().(closure)
		// Assume there are no built-ins for our vm first
		// if sf.tag == "BUILTIN" {
		// 	PC++
		// 	push(OS, apply_builtin(sf.sym, args))
		// 	return
		// }
		RTS.Push(stackFrame{tag: "CALL_FRAME", E: E, PC: PC + 1})
		E.Extend()
		for index, val := range sf.prms {
			E.Set(val, args[index])
		}
		PC = sf.addr
	},
	"TAIL_CALL": func(instr compiler.Instruction) {
		tailcallInstr, ok := instr.(compiler.TAILCALLInstruction)
		if !ok {
			panic("instr is not of type TAILCALLInstruction")
		}
		arity := tailcallInstr.GetArity()
		args := make([]interface{}, arity)
		for i := arity - 1; i >= 0; i-- {
			args[i] = OS.Pop()
		}
		sf := OS.Pop().(*closure)
		//Assume no built in
		// if sf.tag == "BUILTIN" {
		// 	PC++
		// 	// push(OS, apply_builtin(sf.sym, args))
		// 	return
		// }
		E.Extend()
		for index, val := range sf.prms {
			E.Set(val, args[index])
		}
		PC = sf.addr
	},
	"RESET": func(instr compiler.Instruction) {
		top_frame, ok := RTS.Pop().(stackFrame)
		if !ok {
			panic("Frame on RTS is not of type stackFrame")
		}
		if top_frame.tag == "CALL_FRAME" {
			PC = top_frame.PC
			E = top_frame.E
		}
	},
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
func Run(instrs []compiler.Instruction) interface{} {
	global_environment := NewEnvironment() //TODO: Populate global environment with all built-ins
	global_environment.Extend()            //Redundant?
	OS = Stack{}
	PC = 0
	E = global_environment
	RTS = Stack{}
	instrs_arr := instrs
	for instrs_arr[PC].GetTag() != "DONE" {
		// fmt.Println(instrs_arr[PC])
		// fmt.Println(PC)
		microcode[instrs_arr[PC].GetTag()](instrs_arr[PC])
	}
	return OS.Peek()
	// for _, i := range instrs_a {
	// 	// instr := instrs_a[PC]
	// 	fmt.Println(PC)
	// 	// fmt.Println(instrs_a[PC].GetTag())
	// 	// microcode[instr.GetTag()](instr)
	// 	// fmt.Println(i.GetTag())
	// 	// instr, ok := microcode[i.GetTag()]
	// 	if i.GetTag() == "DONE" {
	// 		break
	// 	}
	// 	microcode[i.GetTag()](i)
	// 	// if !ok {
	// 	// 	continue
	// 	// }
	// 	// instr(i)
	// }
	// fmt.Println(instrs[0])
	// return OS.Peek()
}

// Testing for Virtual Machine
// func main() {

// 	// s := Stack{}
// 	// s.Push(1)
// 	// s.Push("two")
// 	// s.Push(3)
// 	// fmt.Println(s.Pop()) // 3, nil
// 	// fmt.Println(s.Pop()) // "two", nil
// 	// fmt.Println(s.Pop()) // 1, nil
// 	// fmt.Println(s.Pop()) // nil, error: stack is empty
// 	str := `["+", [["literal", [2, null]], [["literal", [3, null]], null]]]`
// 	var arr interface{}
// 	// RTS := Stack{}
// 	// E := Stack{}
// 	// PC := Stack{}
// 	// OS := Stack{}
// 	err := json.Unmarshal([]byte(str), &arr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// instrs := json.Unmarshal([]byte(str), &arr)//input instruction array here
// 	fmt.Println(arr)
// }
