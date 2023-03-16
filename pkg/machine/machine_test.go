package machine

import(
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	// instrs := []Instruction{
	// 	{"LDC",1},
	// 	{"DONE",0},
	// }

	//Actual Instructions
	instrs := []map[string]interface{}{
		// {"tag": "ENTER_SCOPE", "syms": make([]interface{}, 0)},
		{"tag": "LDC", "val": 1},
		// {"tag": "EXIT_SCOPE"},
		{"tag": "DONE"},
	}

	fmt.Println(instrs)
    // run(instrs)
	val := run(instrs)
    if val != 1 {
        t.Errorf("value is %d, expected %d", val, 1)
    }
}