package machine

import (
	// "fmt"
	"testing"
	// "reflect"
)

func TestUnOpMicrocode(t *testing.T) {
	testCases := []struct {
		operator string
		arg      interface{}
		want     interface{}
	}{
		{"-unary", 3, -3},
		{"-unary", 3.5, -3.5},
		{"-unary", "test", nil},
		{"!", true, false},
		{"!", false, true},
		{"!", 3, nil},
	}

	for _, tc := range testCases {
		got := unop_microcode[tc.operator](tc.arg)
		if got != tc.want {
			t.Errorf("%s(%v) = %v; want %v", tc.operator, tc.arg, got, tc.want)
		}
	}
}

func TestBinOpMicrocode(t *testing.T) {
	// Test addition with two numbers
	result := binop_microcode["+"](2.5, 3.5)
	expected := 6.0
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test addition with two strings
	var tests = []struct {
		op       string
		x, y     interface{}
		expected interface{}
	}{
		{"+", 2, 3, 5},
		// {"+", 3.5, -3.5, 0},
		{"+", "hello", "world", "helloworld"},
		{"+", 2, "test", "error: + expects two numbers or two strings, got:int and string"},
		// {"+", nil, 1, "error: + expects two numbers or two strings, got:<nil> and int"},
	}

	for _, test := range tests {
		result := binop_microcode[test.op](test.x, test.y)
		if result != test.expected {
			t.Errorf("Error: %v + %v = %v, want %v", test.x, test.y, result, test.expected)
		}
	}

	// Test subtraction with two numbers
	result = binop_microcode["-"](4.5, 3.5)
	expected = 1.0
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test multiplication with two numbers
	result = binop_microcode["*"](2.5, 3.5)
	expected = 8.75
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test division with two numbers
	result = binop_microcode["/"](4.5, 1.5)
	expected = 3.0
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test modulo with two numbers
	result = binop_microcode["%"](5.0, 2.0)
	expectedInt := 1
	if result != expectedInt {
		t.Errorf("Expected %v but got %v", expectedInt, result)
	}

	// Test less than with two numbers
	result = binop_microcode["<"](2.5, 3.5)
	expectedbool := true
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}

	// Test less than or equal with two numbers
	result = binop_microcode["<="](3.5, 3.5)
	expectedbool = true
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}

	// Test greater than or equal with two numbers
	result = binop_microcode[">="](4.5, 3.5)
	expectedbool = true
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}

	// Test greater than with two numbers
	result = binop_microcode[">"](3.5, 2.5)
	expectedbool = true
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}

	// Test equality with two numbers
	result = binop_microcode["=="](3.5, 3.5)
	expectedbool = true
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}

	// Test equality with two numbers
	result = binop_microcode["=="](2.0, 2.0)
	expectedbool = true
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}

	// Test inequality with two strings
	result = binop_microcode["=="]("hello", "hello")
	expectedbool = true
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}

	// Test equality with a string and a nil value
	var s interface{} = "hello"
	result = binop_microcode["=="](s, nil)
	expectedbool = false
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}

	// Test inequality with two integers
	result = binop_microcode["!="](5, 2)
	expectedbool = true
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}

	// Test inequality with two strings
	result = binop_microcode["!="]("hello", "world")
	expectedbool = true
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}

	// Test inequality with two floats
	result = binop_microcode["!="](3.14, 3.14)
	expectedbool = false
	if result != expectedbool {
		t.Errorf("Expected %v but got %v", expectedbool, result)
	}
}
