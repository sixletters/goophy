package compiler

import (
	"testing"
)

func TestInstructions(t *testing.T) {
	instrs := Instructions{
		instrs: []Instruction{
			&LDCBInstruction{tag: "LDCB", val: true},
			&LDCNInstruction{tag: "LDCN", val: 42},
			&LDSInstruction{tag: "LDS", sym: "someSymbol"},
			&DONEInstruction{tag: "DONE"},
			&UNOPSInstruction{tag: "UNOP", sym: negative},
			&BINOPSInstruction{tag: "BINOP", sym: add},
			&POPInstruction{tag: "POP"},
		},
	}

	for _, instr := range instrs.instrs {
		tag := instr.getTag()
		if tag == "" {
			t.Error("Expected non-empty tag, got empty string")
		}

		switch instruction := instr.(type) {
		case *LDCBInstruction:
			if !instruction.getValue() {
				t.Error("Expected true, got false")
			}
		case *LDCNInstruction:
			if instruction.getValue() != 42 {
				t.Errorf("Expected 42, got %d", instruction.getValue())
			}
		case *LDSInstruction:
			if instruction.getSym() != "someSymbol" {
				t.Errorf("Expected someSymbol, got %s", instruction.getSym())
			}
		case *UNOPSInstruction:
			if instruction.getSym() != negative {
				t.Errorf("Expected negative, got %s", instruction.getSym())
			}
		case *BINOPSInstruction:
			if instruction.getSym() != add {
				t.Errorf("Expected add, got %s", instruction.getSym())
			}
		}
	}
}
