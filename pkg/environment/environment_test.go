package environment

import (
	"testing"
)

func TestEnvironment(t *testing.T) {
	// create a new environment and set a variable
	env := NewEnvironment()
	env.Set_assign("x", 42)
	env.Set_assign("z", 80)

	// check that the variable was set correctly
	if val, ok := env.Get("x"); !ok || val != 42 {
		t.Errorf("Failed to set variable x to 42")
	}
	if val, ok := env.Get("z"); !ok || val != 80 {
		t.Errorf("Failed to set variable z to 80")
	}

	// extend the environment and set a new variable
	childEnv := env.Extend()
	childEnv.Set_assign("y", "hello")

	// check that both variables are accessible in the child environment
	if val, ok := childEnv.Get("x"); !ok || val != 42 {
		t.Errorf("Failed to get variable x from child environment")
	}
	if val, ok := childEnv.Get("y"); !ok || val != "hello" {
		t.Errorf("Failed to get variable y from child environment")
	}

	// check that the variables are not accessible in the parent environment
	if _, ok := env.Get("y"); ok {
		t.Errorf("Variable y should not be accessible in parent environment")
	}
	if _, ok := childEnv.Get("z"); ok {
		t.Errorf("Variable z should not be accessible in child environment")
	}
}
