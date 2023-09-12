package functionalgo

import (
	"reflect"
	// "fmt"
	"testing"
)

func TestSanity(t *testing.T) {
	if 2+2 != 4 {
		t.Errorf("Sorry, math's broken, Moose out front should have told ya")
	}
}

func TestCollectSlice(t *testing.T) {
	var start = []int{1, 2, 3, 4}
	var double = func(in int) int {
		return 2 * in
	}
	var expected = []int{2, 4, 6, 8}
	var end = CollectSlice(start, double)

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("DoubleSlice did not work properly, expected: %v, got: %v", expected, end)
	}
}

func TestSelectSlice(t *testing.T) {
	var start = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var evens = func(in int) bool {
		return in%2 == 0
	}
	var expected = []int{2, 4, 6, 8}

	var end = SelectSlice(start, evens)

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("EvensSlice didn't work properly, expected: %v, got: %v", expected, end)
	}
}

func TestInjectSlice(t *testing.T) {
	var start = []int{1, 2, 3, 4, 5}
	var sum = func(l, r int) int {
		return l + r
	}

	var expected = 15

	var end = InjectSlice(start, 0, sum)

	if end != expected {
		t.Errorf("SumSlice didn't work properly, expected: %d, got: %d", expected, end)
	}
}

func TestCollectChan(t *testing.T) {
	var start = make(chan int, 4)
	var double = func(in int) int {
		return 2 * in
	}

	var expected = []int{2, 4, 6, 8}
	var endChan = CollectChan(start, double)
	var end []int

	start <- 1
	start <- 2
	start <- 3
	start <- 4
	close(start)

	for i := range endChan {
		end = append(end, i)
	}

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("DoubleChan did not work properly, expected: %v, got: %v", expected, end)
	}
}

func TestSelectChan(t *testing.T) {
	start := make(chan int, 8)
	var odd = func(in int) bool {
		return in%2 == 1
	}

	var expected = []int{1, 3, 5, 7}
	var endChan = SelectChan(start, odd)
	var end []int

	start <- 1
	start <- 2
	start <- 3
	start <- 4
	start <- 5
	start <- 6
	start <- 7
	start <- 8
	close(start)

	for i := range endChan {
		end = append(end, i)
	}

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("OddsChan didn't work properly, expected: %v, got: %v", expected, end)
	}
}

func TestInjectChan(t *testing.T) {
	start := make(chan int, 10)
	var sum = func(l, r int) int {
		return l + r
	}

	expected := 55
	var endChan = InjectChan(start, 0, sum)
	var end int

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

	for i := range endChan {
		end = i
	}

	if end != expected {
		t.Errorf("SumChan didn't work properly, expected: %d, got: %d", expected, end)
	}
}

func TestGenerateRangeInt(t *testing.T) {
	rangeChan, err := GenerateRange(1, 10, 2)
	var end []int
	var expected = []int{1, 3, 5, 7, 9}
	if err != nil {
		t.Errorf("GenerateRange returned an error when it wasn't supposed to")
	}

	for i := range rangeChan {
		end = append(end, i)
	}

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("GenerateRange[int] didn't work properly, expected: %v, got: %v", expected, end)
	}
}

func TestGenerateRangeFloat64(t *testing.T) {
	rangeChan, err := GenerateRange(0.0, 2.0, 0.25)
	var end []float64
	var expected = []float64{0.0, 0.25, 0.5, 0.75, 1.0, 1.25, 1.5, 1.75}
	if err != nil {
		t.Errorf("GenerateRange returned an error when it wasn't supposed to")
	}

	for i := range rangeChan {
		end = append(end, i)
	}

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("GenerateRange[float64] didn't work properly, expected: %v, got: %v", expected, end)
	}
}

func TestGenerateFibonacciUpTo(t *testing.T) {
	fibs := GenerateFibonacciUpTo(400)
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}
	var end []int

	for i := range fibs {
		end = append(end, i)
	}

	if !reflect.DeepEqual(end, expected) {
		t.Errorf("GenerateRange[float64] didn't work properly, expected: %v, got: %v", expected, end)
	}
}
