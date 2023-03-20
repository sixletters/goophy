package machine

// import (
// 	"cs4215/goophy/pkg/compiler"
// 	"testing"
// )

// func TestRun(t *testing.T) {
// 	// instrs := []Instruction{
// 	// 	{"LDC",1},
// 	// 	{"DONE",0},
// 	// }

// 	// //Actual Instructions
// 	// instrs := []map[string]interface{}{
// 	// 	// {"tag": "ENTER_SCOPE", "syms": make([]interface{}, 0)},
// 	// 	{"tag": "LDC", "val": 1},
// 	// 	// {"tag": "EXIT_SCOPE"},
// 	// 	{"tag": "DONE"},
// 	// }
// 	instrs := compiler.Instructions{
// 		instrs: []compiler.Instruction{
// 			// Test LDCBInstruction
// 			&LDCBInstruction{tag: "LDCB", val: true},

// 			// Test LDCNInstruction
// 			&LDCNInstruction{tag: "LDCN", val: 42},

// 			&LDCNInstruction{tag: "LDCN", val: 44},

// 			// // Test LDSInstruction
// 			// &LDSInstruction{tag: "LDS", sym: "x"},

// 			// Test UNOPSInstruction
// 			&UNOPInstruction{tag: "UNOP", sym: negative},

// 			// Test BINOPSInstruction
// 			&BINOPInstruction{tag: "BINOP", sym: add},

// 			// Test DONEInstruction
// 			&DONEInstruction{tag: "DONE"},

// 			// // Test POPInstruction
// 			// &POPInstruction{tag: "POP"},

// 			// // Test JOFInstruction
// 			// &JOFInstruction{tag: "JOF", addr: 10},

// 			// // Test GOTOInstruction
// 			// &GOTOInstruction{tag: "GOTO", addr: 20},

// 			// // Test ENTERSCOPEInstruction
// 			// &ENTERSCOPEInstruction{tag: "ENTER_SCOPE", syms: []string{"x", "y"}},

// 			// // Test EXITSCOPEInstruction
// 			// &EXITSCOPEInstruction{tag: "EXIT_SCOPE"},

// 			// // Test LDFInstruction
// 			// &LDFInstruction{tag: "LDF", prms: []string{"a", "b"}, addr: 30},

// 			// // Test CALLInstruction
// 			// &CALLInstruction{tag: "CALL", arity: 2},

// 			// // Test TAILCALLInstruction
// 			// &TAILCALLInstruction{tag: "TAIL_CALL", arity: 2},

// 			// // Test RESETInstruction
// 			// &RESETInstruction{tag: "RESET"},
// 		},
// 	}

// 	// fmt.Println(instrs)
// 	Run(instrs)
// 	val := Run(instrs)
// 	// return type currently is interface, change this to PV
// 	if val != 1 {
// 		t.Errorf("value is %d, expected %d", val, 42)
// 	}
// }
