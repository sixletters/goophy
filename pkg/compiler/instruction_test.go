package compiler

// import (
// 	"fmt"
// 	"testing"
// )

// func TestInstructions(t *testing.T) {
// 	instrs := Instructions{
// 		instrs: []Instruction{
// 			// Test LDCBInstruction
// 			&LDCBInstruction{tag: "LDCB", val: true},

// 			// Test LDCNInstruction
// 			&LDCNInstruction{tag: "LDCN", val: 42},

// 			// Test LDSInstruction
// 			&LDSInstruction{tag: "LDS", sym: "x"},

// 			// Test DONEInstruction
// 			&DONEInstruction{tag: "DONE"},

// 			// Test UNOPSInstruction
// 			&UNOPInstruction{tag: "UNOP", sym: negative},

// 			// Test BINOPSInstruction
// 			&BINOPInstruction{tag: "BINOP", sym: add},

// 			// Test POPInstruction
// 			&POPInstruction{tag: "POP"},

// 			// Test JOFInstruction
// 			&JOFInstruction{tag: "JOF", addr: 10},

// 			// Test GOTOInstruction
// 			&GOTOInstruction{tag: "GOTO", addr: 20},

// 			// Test ENTERSCOPEInstruction
// 			&ENTERSCOPEInstruction{tag: "ENTER_SCOPE", syms: []string{"x", "y"}},

// 			// Test EXITSCOPEInstruction
// 			&EXITSCOPEInstruction{tag: "EXIT_SCOPE"},

// 			// Test LDFInstruction
// 			&LDFInstruction{tag: "LDF", prms: []string{"a", "b"}, addr: 30},

// 			// Test CALLInstruction
// 			&CALLInstruction{tag: "CALL", arity: 2},

// 			// Test TAILCALLInstruction
// 			&TAILCALLInstruction{tag: "TAIL_CALL", arity: 2},

// 			// Test RESETInstruction
// 			&RESETInstruction{tag: "RESET"},
// 		},
// 	}
// 	fmt.Println(instrs)
// 	// Check if the tags and parameters are correct
// 	for _, instr := range instrs.instrs {
// 		tag := instr.getTag()
// 		if tag == "" {
// 			t.Errorf("Empty tag found in instruction: %+v", instr)
// 		}

// 		switch i := instr.(type) {
// 		case *LDCBInstruction:
// 			if i.val != true {
// 				t.Errorf("Invalid value found in LDCBInstruction: %+v", i)
// 			}
// 		case *LDCNInstruction:
// 			if i.val != 42 {
// 				t.Errorf("Invalid value found in LDCNInstruction: %+v", i)
// 			}
// 		case *LDSInstruction:
// 			if i.sym != "x" {
// 				t.Errorf("Invalid sym found in LDSInstruction: %+v", i)
// 			}
// 		case *UNOPInstruction:
// 			if i.sym != negative {
// 				t.Errorf("Invalid sym found in UNOPSInstruction: %+v", i)
// 			}
// 		case *BINOPInstruction:
// 			if i.sym != add {
// 				t.Errorf("Invalid sym found in BINOPSInstruction: %+v", i)
// 			}
// 		case *JOFInstruction:
// 			if i.addr != 10 {
// 				t.Errorf("Invalid addr found in JOFInstruction: %+v", i)
// 			}
// 		case *GOTOInstruction:
// 			if i.addr != 20 {
// 				t.Errorf("Invalid addr found in GOTOInstruction: %+v", i)
// 			}
// 		case *ENTERSCOPEInstruction:
// 			if len(i.syms) != 2 || i.syms[0] != "x" || i.syms[1] != "y" {
// 				t.Errorf("Invalid syms found in ENTERSCOPEInstruction: %+v", i)
// 			}
// 		case *EXITSCOPEInstruction:
// 			// No parameters to validate
// 		case *LDFInstruction:
// 			if len(i.prms) != 2 || i.prms[0] != "a" || i.prms[1] != "b" || i.addr != 30 {
// 				t.Errorf("Invalid parameters found in LDFInstruction: %+v", i)
// 			}
// 		case *CALLInstruction:
// 			if i.arity != 2 {
// 				t.Errorf("Invalid arity found in CALLInstruction: %+v", i)
// 			}
// 		case *TAILCALLInstruction:
// 			if i.arity != 2 {
// 				t.Errorf("Invalid arity found in TAILCALLInstruction: %+v", i)
// 			}
// 		case *RESETInstruction:
// 			// No parameters to validate
// 		case *POPInstruction:
// 			// No parameters to validate
// 		case *DONEInstruction:
// 			// No parameters to validate
// 		default:
// 			t.Errorf("Unknown instruction type: %+v", i)
// 		}
// 	}
// }
