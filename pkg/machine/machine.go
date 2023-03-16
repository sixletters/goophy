// Idealized Virtual Machine

package machine

import(
	// "fmt"
	// "errors"
	// "encoding/json" // To convert array in string form back into array
	// "cs4215/goophy/pkg/environment"
)
// Parser
// Microcode
// Scheduler

// const global_environment = heap
// RTS, OS, PC, E
var RTS Stack
var OS Stack
var E EnvironmentFrame
var PC int

// Microcode
// var microcode = map[string]func(Instruction){
//     "LDC": func(instr Instruction) {
//         PC++
// 		OS.Push(instr.val)
//     },
var microcode = map[string]func(map[string]interface{}){
    "LDC": func(instr map[string]interface{}) {
        PC++
        OS.Push(instr["val"])
    },
}

    // "UNOP": func(instr Instruction) {
    //     PC++
	// 	OS.Push(apply_unop(instr.sym, OS.Pop()))
    // },
    // "BINOP": func(instr *Instruction) {
    //     PC++
	// 	OS.Push(apply_binop(instr.sym, OS.Pop(), OS.Pop()))
    // },
    // "POP": func(instr *Instruction) {
    //     PC++
    //     OS.Pop()
    // },
    // "JOF": func(instr *Instruction) {
    //     PC = boolToInt(OS.Pop())*instr.addr + boolToInt(!OS.Peek())*(PC + 1 - instr.addr)
    // },
    // "GOTO": func(instr *Instruction) {
    //     PC = instr.addr
    // },
    // "ENTER_SCOPE": func(instr *Instruction) {
    //     PC++
    //     RTS.Push(&frame{tag: "BLOCK_FRAME", env: E})
    //     locals := instr.syms
    //     unassigneds := make([]interface{}, len(locals))
    //     for i := range unassigneds {
    //         unassigneds[i] = unassigned
    //     }
    //     E = extend(locals, unassigneds, E)
    // },
    // "EXIT_SCOPE": func(instr *Instruction) {
    //     PC++
    //     E = RTS.Pop().env
    // },
    // "LD": func(instr *Instruction) {
    //     PC++
	// 	OS.Push(lookup(instr.sym, E))
    // },
    // "ASSIGN": func(instr *Instruction) {
    //     PC++
    //     assign_value(instr.sym, OS.Peek(), E)
    // },
    // "LDF": func(instr *Instruction) {
    //     PC++
	// 	OS.Push({tag: "CLOSURE", prms: instr.prms, addr: instr.addr, env: E})
    // },
    // "CALL": func(instr *Instruction) {
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
    // "TAIL_CALL": func(instr *Instruction) {
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
    // "RESET": func(instr *Instruction) {
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
// type Instruction struct{
// 	tag string
// 	val interface{}
// 	// sym string
// }

// TODO: Check allowable types for return
// func run(instrs) interface{} {
func run(instrs[]map[string]interface{}) interface{} {
	// global_environment := NewEnvironmentStack()
	// global_environment.Extend()
    OS = Stack{}
    PC = 0
    // E = global_environment
    // RTS = Stack{}

	//Placeholder instruction
	// instrs := []Instruction{
	// 	{"LDC",1},
	// 	{"DONE",0},
	// }

    for instrs[PC]["tag"] != "DONE" {
        instr := instrs[PC]
        // fmt.Println(instr)
        microcode[instr["tag"].(string)](instr)
    }
	// fmt.Println(instrs[0])
    return OS.Peek()
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