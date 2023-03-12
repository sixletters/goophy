package main

import(
	"fmt"
	"errors"
	"encoding/json" // To convert array in string form back into array
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
	fmt.Println(arr)
}