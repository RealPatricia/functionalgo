package functionalgo

import (
	"github.com/RealPatricia/pth/assert"
	"testing"
)

func TestMapChan(t *testing.T) {
	var start = MakeChan(1, 2, 3, 4)
	var double = func(in int) int {
		return 2 * in
	}

	var expected = MakeSlice(2, 4, 6, 8)
	var endChan = MapChan(start, double)

	end := FinalValue(ChanToSlice(endChan))

	if assert.Equal(end, expected, false) {
		t.Errorf("Double did not work properly, expected: %v, got: %v", expected, end)
	}
}

func TestFilterChan(t *testing.T) {
	start := MakeChan(1, 2, 3, 4, 5, 6, 7, 8)
	var odd = func(in int) bool {
		return in%2 == 1
	}

	var expected = []int{1, 3, 5, 7}
	var endChan = FilterChan(start, odd)

	end := FinalValue(ChanToSlice(endChan))

	if assert.Equal(end, expected, false) {
		t.Errorf("Odds didn't work properly, expected: %v, got: %v", expected, end)
	}
}

func TestReduceChan(t *testing.T) {
	start := MakeChan(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	var sum = func(l, r int) int {
		return l + r
	}

	expected := 55
	var endChan = ReduceChan(start, sum, 0)

	end := FinalValue(endChan)

	if end != expected {
		t.Errorf("Sum didn't work properly, expected: %d, got: %d", expected, end)
	}
}

func TestChanToCollection(t *testing.T) {
	start := MakeChan(1, 2, 3)
	endChan := ChanToSlice(start)
	expected := MakeSlice(1, 2, 3)

	end := FinalValue(endChan)

	if assert.Equal(end, expected, false) {
		t.Errorf("ChanToCollection didn't work properly, expected: %d, got: %d", expected, end)
	}
}
