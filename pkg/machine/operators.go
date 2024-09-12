package machine

import (
	"reflect"
	// "fmt"
)

var unop_microcode = map[string]func(interface{}) interface{}{
	"-unary": func(x interface{}) interface{} {
		switch x.(type) {
		case int:
			return -x.(int)
		case float64:
			return -x.(float64)
		default:
			return nil // return an error message or value here
		}
	},
	"!": func(x interface{}) interface{} {
		if ok := isBool(x); ok {
			return !x.(bool)
		}
		return nil // return an error message or value here
	},
}

var binop_microcode = map[string]func(x, y interface{}) interface{}{
	"=": func(x, y interface{}) interface{} {
		return x == y
	},

	"+": func(x, y interface{}) interface{} {
		if x == nil || y == nil {
			return nil
		}
		if isNumber(x) && isNumber(y) {
			switch x.(type) {
			case int:
				return x.(int) + y.(int)
			case float64:
				return x.(float64) + y.(float64)
			}
		} else if isString(x) && isString(y) {
			return x.(string) + y.(string)
		}
		return "error: + expects two numbers or two strings, got:" +
			reflect.TypeOf(x).String() + " and " +
			reflect.TypeOf(y).String()
	},

	"*": func(x, y interface{}) interface{} {
		// return x.(float64) * y.(float64)
		switch x.(type) {
		case int:
			return x.(int) * y.(int)
		case float64:
			return x.(float64) * y.(float64)
		default:
			panic("invalid type")
		}
	},
	"-": func(x, y interface{}) interface{} {
		switch x.(type) {
		case int:
			return x.(int) - y.(int)
		case float64:
			return x.(float64) - y.(float64)
		default:
			panic("invalid type")
		}
	},
	"/": func(x, y interface{}) interface{} {
		switch x.(type) {
		case int:
			return x.(int) / y.(int)
		case float64:
			return x.(float64) / y.(float64)
		default:
			panic("invalid type")
		}
	},
	"%": func(x, y interface{}) interface{} {
		switch x.(type) {
		case int:
			return x.(int) % y.(int)
		default:
			panic("invalid type")
		}
	},
	"<": func(x, y interface{}) interface{} {
		switch x.(type) {
		case int:
			return x.(int) < y.(int)
		case float64:
			return x.(float64) < y.(float64)
		default:
			panic("invalid type")
		}
	},
	"<=": func(x, y interface{}) interface{} {
		switch x.(type) {
		case int:
			return x.(int) <= y.(int)
		case float64:
			return x.(float64) <= y.(float64)
		default:
			panic("invalid type")
		}
	},
	">=": func(x, y interface{}) interface{} {
		switch x.(type) {
		case int:
			return x.(int) >= y.(int)
		case float64:
			return x.(float64) >= y.(float64)
		default:
			panic("invalid type")
		}
	},
	">": func(x, y interface{}) interface{} {
		switch x.(type) {
		case int:
			return x.(int) > y.(int)
		case float64:
			return x.(float64) > y.(float64)
		default:
			panic("invalid type")
		}
	},
	"==": func(x, y interface{}) interface{} {
		return x == y
	},
	"!=": func(x, y interface{}) interface{} {
		return x != y
	},
}

func apply_binop(op string, v2 interface{}, v1 interface{}) interface{} {
	return binop_microcode[op](v1, v2)
}

func apply_unop(op string, v interface{}) interface{} {
	return unop_microcode[op](v)
}
