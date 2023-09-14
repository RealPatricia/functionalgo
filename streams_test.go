package functionalgo

import (
	"reflect"
	"testing"
)

func TestMapStream(t *testing.T) {
	var start = make(chan int, 4)
	var double = func(in int) int {
		return 2 * in
	}

	var expected = []int{2, 4, 6, 8}
	var endChan = MapStream(start, double)

	start <- 1
	start <- 2
	start <- 3
	start <- 4
	close(start)

	end := []int(<-StreamToCollection(endChan))

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("Double did not work properly, expected: %v, got: %v", expected, end)
	}
}

func TestFilterStream(t *testing.T) {
	start := make(chan int, 8)
	var odd = func(in int) bool {
		return in%2 == 1
	}

	var expected = []int{1, 3, 5, 7}
	var endStream = FilterStream(start, odd)

	start <- 1
	start <- 2
	start <- 3
	start <- 4
	start <- 5
	start <- 6
	start <- 7
	start <- 8
	close(start)

	end := []int(<-StreamToCollection(endStream))

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("Odds didn't work properly, expected: %v, got: %v", expected, end)
	}
}

func TestReduceStream(t *testing.T) {
	start := make(chan int, 10)
	var sum = func(l, r int) int {
		return l + r
	}

	expected := 55
	var endStream = ReduceStream(start, sum, 0)

	start <- 1
	start <- 2
	start <- 3
	start <- 4
	start <- 5
	start <- 6
	start <- 7
	start <- 8
	start <- 9
	start <- 10
	close(start)

	end := FinalValue(endStream)

	if end != expected {
		t.Errorf("Sum didn't work properly, expected: %d, got: %d", expected, end)
	}
}

func TestStreamToCollection(t *testing.T) {
	start := make(chan int, 3)
	endStream := StreamToCollection(start)
	expected := Collection[int]{1, 2, 3}

	start <- 1
	start <- 2
	start <- 3
	close(start)

	end := FinalValue(endStream)

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("StreamToCollection didn't work properly, expected: %d, got: %d", expected, end)
	}
}
