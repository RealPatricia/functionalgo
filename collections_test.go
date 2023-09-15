package functionalgo

import (
	"github.com/RealPatricia/pth/assert"
	"testing"
)

func TestMapSlice(t *testing.T) {
	var start = MakeSlice(1, 2, 3, 4)
	var double = func(in int) int {
		return 2 * in
	}
	var expected = MakeSlice(2, 4, 6, 8)
	var end = MapSlice(start, double)

	if assert.Equal(end, expected, false) {
		t.Errorf("Double did not work properly, expected: %v, got: %v", expected, end)
	}
}

func TestFilterSlice(t *testing.T) {
	var start = MakeSlice(1, 2, 3, 4, 5, 6, 7, 8)
	var evens = func(in int) bool {
		return in%2 == 0
	}
	var expected = MakeSlice(2, 4, 6, 8)
	var end = FilterSlice(start, evens)

	if assert.Equal(end, expected, false) {
		t.Errorf("Evens didn't work properly, expected: %v, got: %v", expected, end)
	}
}

func TestReduceSlice(t *testing.T) {
	var start = MakeSlice(1, 2, 3, 4, 5)
	var sum = func(l, r int) int {
		return l + r
	}

	var expected = 15

	var end = ReduceSlice(start, sum, 0)

	if end != expected {
		t.Errorf("Sum didn't work properly, expected: %d, got: %d", expected, end)
	}
}

func TestSliceToChan(t *testing.T) {
	start := []int{1, 2, 3}
	endChan := SliceToChan(start)
	expected := []int{1, 2, 3}

	end := FinalValue(ChanToSlice(endChan))

	if assert.Equal(end, expected, false) {
		t.Errorf("SliceToChan didn't work properly, expected: %d, got: %d", expected, end)
	}
}
