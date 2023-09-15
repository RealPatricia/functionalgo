package functionalgo

import (
	"github.com/RealPatricia/pth/assert"
	"slices"
	"testing"
)

func TestMapDictionary(t *testing.T) {
	var start = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	var double = func(in int) int {
		return 2 * in
	}
	var expected = map[string]int{"one": 2, "two": 4, "three": 6, "four": 8}
	var end = MapMap(start, double)

	if assert.Equal(end, expected, false) {
		t.Errorf("Double did not work properly, expected: %v, got: %v", expected, end)
	}
}

func TestFilterDictionary(t *testing.T) {
	var start = map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8,
	}
	var evens = func(in int) bool {
		return in%2 == 0
	}
	var expected = map[string]int{"two": 2, "four": 4, "six": 6, "eight": 8}

	var end = FilterMap(start, evens)

	if assert.Equal(end, expected, false) {
		t.Errorf("Evens didn't work properly, expected: %v, got: %v", expected, end)
	}
}

func TestReduceDictionary(t *testing.T) {
	var start = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}
	var sum = func(l, r int) int {
		return l + r
	}

	var expected = 15

	var end = ReduceMap(start, sum, 0)

	if end != expected {
		t.Errorf("Sum didn't work properly, expected: %d, got: %d", expected, end)
	}
}

func TestDictionaryToStream(t *testing.T) {
	start := map[string]int{"one": 1, "two": 2, "three": 3}
	endChan := MapToChan(start)
	expected := MakeSlice(1, 2, 3)

	end := FinalValue(ChanToSlice(endChan))
	slices.Sort(end)

	if assert.Equal(end, expected, false) {
		t.Errorf("Sum didn't work properly, expected: %d, got: %d", expected, end)
	}
}

func TestDictionaryToCollection(t *testing.T) {
	start := map[string]int{"one": 1, "two": 2, "three": 3}
	end := MapToSlice(start)
	slices.Sort(end)
	expected := MakeSlice(1, 2, 3)

	if assert.Equal(end, expected, false) {
		t.Errorf("Sum didn't work properly, expected: %d, got: %d", expected, end)
	}
}
