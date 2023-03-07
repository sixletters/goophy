package main

import (
	"cs4215/goophy/pkg/environment"
	"fmt"
)

func main() {
	fmt.Println("HELLO WORLD")
	sf := environment.NewStackFrame()

	sf.SetVar("x", 10)
	sf.SetVar("y", "hello")
	square := func(x int) int {
		return x * x
	}

	sf.SetVar("a", square)

	x, ok := sf.GetVar("x")
	if ok {
		fmt.Printf("x = %v\n", x)
	}

	y, ok := sf.GetVar("y")
	if ok {
		fmt.Printf("y = %v\n", y)
	}

	fn, ok := sf.GetVar("a")
	if ok {
		result := fn.(func(int) int)(5)
		fmt.Printf("square(5) = %d\n", result)
	}

	frame_address := heap_allocate_Frame(primitive_values.length)

	empty_environment = nil

	global_environment = extendEnvironment(global_frame, empty_environment)

	global_environment := environment.heap_Environment_extend(frame_address,
		heap_empty_Environment)

}
