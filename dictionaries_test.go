package functionalgo

import (
	"reflect"
	"slices"
	"testing"
)

func TestMapDictionary(t *testing.T) {
	var start = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	var double = func(in int) int {
		return 2 * in
	}
	var expected = Dictionary[string, int]{"one": 2, "two": 4, "three": 6, "four": 8}
	var end = MapDictionary(start, double)

	if !reflect.DeepEqual(end, expected) {
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
	var expected = Dictionary[string, int]{"two": 2, "four": 4, "six": 6, "eight": 8}

	var end = FilterDictionary(start, evens)

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("Evens didn't work properly, expected: %v, got: %v", expected, end)
	}
}

func TestReduceDictionary(t *testing.T) {
	var start = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}
	var sum = func(l, r int) int {
		return l + r
	}

	var expected = 15

	var end = ReduceDictionary(start, sum, 0)

	if end != expected {
		t.Errorf("Sum didn't work properly, expected: %d, got: %d", expected, end)
	}
}

func TestDictionaryToStream(t *testing.T) {
	start := map[string]int{"one": 1, "two": 2, "three": 3}
	endChan := DictionaryToStream(start)
	var end []int
	expected := []int{1, 2, 3}

	for e := range endChan {
		end = append(end, e)
	}

	slices.Sort(end)

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("Sum didn't work properly, expected: %d, got: %d", expected, end)
	}
}

func TestDictionaryToCollection(t *testing.T) {
	start := map[string]int{"one": 1, "two": 2, "three": 3}
	end := DictionaryToCollection(start)
	slices.Sort(end)
	expected := Collection[int]([]int{1, 2, 3})

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("Sum didn't work properly, expected: %d, got: %d", expected, end)
	}
}
