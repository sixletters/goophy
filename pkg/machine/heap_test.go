package machine

import "testing"

func TestHeapMake(t *testing.T) {
    heap_make(10)

    if len(HEAP) != 10*mega {
        t.Errorf("HEAP size is %d bytes, expected %d", len(HEAP), 10*mega)
    }
}

func TestHeapGetAndSetWord(t *testing.T) {
    heap_make(10)

    // Set a word and get it back
    x := 42.0
    heap_set_word_at_index(0, x)
    y := heap_get_word_at_index(0)

    if x != y {
        t.Errorf("heap_set_word_at_index(0, %f); heap_get_word_at_index(0) = %f, expected %f", x, y, x)
    }

    // Set another word and get it back
    a := 3.14159
    heap_set_word_at_index(2, a)
    b := heap_get_word_at_index(2)

    if a != b {
        t.Errorf("heap_set_word_at_index(2, %f); heap_get_word_at_index(2) = %f, expected %f", a, b, a)
    }
}


func TestHeapDisplay(t *testing.T) {
    // TODO: write a test for heap_display()
    // You can use a buffer to capture the output and check that it is correct.
}
