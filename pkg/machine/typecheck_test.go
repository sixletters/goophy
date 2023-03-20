package machine

import (
    "testing"
)

func TestIsNumber(t *testing.T) {
    cases := []struct {
        input interface{}
        want  bool
    }{
        {1, true},
        {1.0, true},
        {float32(1), true},
        {float64(1), true},
        {int32(1), true},
        {int64(1), true},
        {"1", false},
        {true, false},
        {nil, false},
    }

    for _, c := range cases {
        got := isNumber(c.input)
        if got != c.want {
            t.Errorf("isNumber(%v) == %t, want %t", c.input, got, c.want)
        }
    }
}

func TestIsString(t *testing.T) {
    cases := []struct {
        input interface{}
        want  bool
    }{
        {"hello", true},
        {123, false},
        {true, false},
        {nil, false},
    }

    for _, c := range cases {
        got := isString(c.input)
        if got != c.want {
            t.Errorf("isString(%v) == %t, want %t", c.input, got, c.want)
        }
    }
}

func TestIsBool(t *testing.T) {
    cases := []struct {
        input interface{}
        want  bool
    }{
        {true, true},
        {false, true},
        {123, false},
        {"true", false},
        {nil, false},
    }

    for _, c := range cases {
        got := isBool(c.input)
        if got != c.want {
            t.Errorf("isBool(%v) == %t, want %t", c.input, got, c.want)
        }
    }
}
