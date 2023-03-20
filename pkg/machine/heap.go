package machine

import (
    // "fmt"
    "encoding/binary"
    "math"
)

/* *************************

HEAP
*************************/
// HEAP is a byte slice
var HEAP []byte

const word_size = 8
const mega = 1 << 20 // equivalent to 2 ** 20

// heap_make allocates a heap of given size (in megabytes)
// and initializes HEAP with a byte slice of that size.
func heap_make(megaBytes int) {
	HEAP = make([]byte, mega*megaBytes)
}

// we randomly pick a heap size of 10 megabytes
func init() {
	heap_make(10)
}

// free is the next free address;
// we keep allocating as if there was no tomorrow
var free = 0

// for debugging: display all bits of the heap
// func heap_display() {
// 	fmt.Println("heap:")
// 	for i := 0; i < free; i += word_size {
// 		fmt.Printf("%d %f\n", i, heap_get_word_at_index(i))
// 	}
// }
func heap_get_word_at_index(index int) float64 {
    wordBytes := HEAP[index*word_size : (index+1)*word_size]
    bits := binary.LittleEndian.Uint64(wordBytes)
    return math.Float64frombits(bits)
}

func heap_set_word_at_index(index int, x float64) {
    wordBytes := make([]byte, word_size)
    bits := math.Float64bits(x)
    binary.LittleEndian.PutUint64(wordBytes, bits)
    copy(HEAP[index*word_size:(index+1)*word_size], wordBytes)
}

