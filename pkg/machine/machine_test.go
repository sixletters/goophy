package machine

import "testing"

func TestRun(t *testing.T) {
	instrs := []Instruction{
		{"LDC",1},
		{"DONE",0},
	}
    // run(instrs)
	val := run(instrs)
    if val != 1 {
        t.Errorf("value is %d, expected %d", val, 1)
    }
}