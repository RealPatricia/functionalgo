package functionalgo

import (
	"reflect"
	"testing"
)

func TestMapCollection(t *testing.T) {
	var start = Collection[int]{1, 2, 3, 4}
	var double = func(in int) int {
		return 2 * in
	}
	var expected = Collection[int]{2, 4, 6, 8}
	var end = MapCollection(start, double)

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("Double did not work properly, expected: %v, got: %v", expected, end)
	}
}

func TestFilterCollection(t *testing.T) {
	var start = Collection[int]{1, 2, 3, 4, 5, 6, 7, 8}
	var evens = func(in int) bool {
		return in%2 == 0
	}
	var expected = Collection[int]{2, 4, 6, 8}

	var end = FilterCollection(start, evens)

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("Evens didn't work properly, expected: %v, got: %v", expected, end)
	}
}

func TestReduceCollection(t *testing.T) {
	var start = Collection[int]{1, 2, 3, 4, 5}
	var sum = func(l, r int) int {
		return l + r
	}

	var expected = 15

	var end = ReduceCollection(start, sum, 0)

	if end != expected {
		t.Errorf("Sum didn't work properly, expected: %d, got: %d", expected, end)
	}
}

func TestCollectionToStream(t *testing.T) {
	start := Collection[int]{1, 2, 3}
	endChan := CollectionToStream(start)
	expected := []int{1, 2, 3}

	end := []int(<-StreamToCollection(endChan))

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("CollectionToStream didn't work properly, expected: %d, got: %d", expected, end)
	}
}
