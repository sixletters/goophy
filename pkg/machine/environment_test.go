package machine

import (
	"testing"
)

func TestEnvironmentFrameSetAndGetVar(t *testing.T) {
	ef := NewEnvironmentFrame()
	ef.setVar("foo", 42)
	val, ok := ef.getVar("foo")
	if !ok {
		t.Errorf("expected ok to be true but got false")
	}
	if val != 42 {
		t.Errorf("expected val to be 42 but got %v", val)
	}
}

func TestEnvironmentStackExtend(t *testing.T) {
	es := NewEnvironmentStack()
	es.Extend()
	if len(es.envFrames) != 1 {
		t.Errorf("expected len(es.envFrames) to be 1 but got %d", len(es.envFrames))
	}
	es.Extend()
	if len(es.envFrames) != 2 {
		t.Errorf("expected len(es.envFrames) to be 2 but got %d", len(es.envFrames))
	}
}

func TestEnvironmentStackPop(t *testing.T) {
	es := NewEnvironmentStack()
	es.Extend()
	es.Extend()
	es.Pop()
	if len(es.envFrames) != 1 {
		t.Errorf("expected len(es.envFrames) to be 1 but got %d", len(es.envFrames))
	}
}

func TestEnvironmentStackSetAndGet(t *testing.T) {
	es := NewEnvironmentStack()
	es.Extend()
	es.Set("foo", 42)
	val, ok := es.Get("foo")
	if !ok {
		t.Errorf("expected ok to be true but got false")
	}
	if val != 42 {
		t.Errorf("expected val to be 42 but got %v", val)
	}
}

func TestEnvironmentStackGetNotFound(t *testing.T) {
	es := NewEnvironmentStack()
	es.Extend()
	val, ok := es.Get("foo")
	if ok {
		t.Errorf("expected ok to be false but got true")
	}
	if val != nil {
		t.Errorf("expected val to be nil but got %v", val)
	}
}

func TestPopWhenEmpty(t *testing.T) {
	env := NewEnvironmentStack()
	env.Pop()
	if len(env.envFrames) != 0 {
		t.Errorf("Expected length of environment stack to be 0, but got %d", len(env.envFrames))
	}
}

