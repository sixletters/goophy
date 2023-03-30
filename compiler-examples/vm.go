package main

import (
	"encoding/json" // To convert array in string form back into array
	"errors"
)

// Parser
// Microcode
// Scheduler
// RTS, OS, PC, E
type Stack []interface{}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, nil
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

type Environment struct {
	Variables map[string]interface{}
	Functions map[string]interface{}
}

// func run(prog) int {
// 	// Initialize rts,os,pc,e

// 	// Return head of OS
// 	return 1
// }

// Operand Stack

func main() {
	// s := Stack{}
	// s.Push(1)
	// s.Push("two")
	// s.Push(3)
	// fmt.Println(s.Pop()) // 3, nil
	// fmt.Println(s.Pop()) // "two", nil
	// fmt.Println(s.Pop()) // 1, nil
	// fmt.Println(s.Pop()) // nil, error: stack is empty
	str := `["+", [["literal", [2, null]], [["literal", [3, null]], null]]]`
	var arr interface{}
	// RTS := Stack{}
	// E := Stack{}
	// PC := Stack{}
	// OS := Stack{}
	err := json.Unmarshal([]byte(str), &arr)
	if err != nil {
		panic(err)
	}
	// instrs := json.Unmarshal([]byte(str), &arr)//input instruction array here
	// fmt.Println(arr)
	// u need an operand stack aka stash
	// u need a stack for the commands <- instruction aka agenda
	// ^ need handle using array operations
	// u also need current environment <- for evaluation of names/variables
	// lambda and apply is fking hard, need to consider the parameters wrt to current environment and leave mark so that dont fully execute if theres an ifelse block
	// builtin instructions are special, like math stuff <- we slap these in the global frame under builtin mapping
	// lookup needed and applied when the top of agenda is a name
}
