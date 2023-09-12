package functionalgo

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func GenerateRange[T Number](start, end, step T) (chan T, error) {
	var checkStart, checkEnd, checkStep = float64(start), float64(end), float64(step)
	var steppingDown = math.Signbit(checkStep)
	if (start > end && !steppingDown) || (start < end && steppingDown) {
		var zero chan T
		close(zero)
		return zero, fmt.Errorf("cannot make range with the given step size and start and end points")
	}

	var rangeBreadth = math.Abs(checkStart - checkEnd)
	var rangeSteps = math.Abs(rangeBreadth / checkStep)
	var maxChanSize = int(math.Ceil(rangeSteps))
	out := make(chan T, maxChanSize)

	go func() {
		defer close(out)
		for i := start; i < end; i += step {
			out <- i
		}
	}()

	return out, nil
}

func GenerateFibonacciUpTo(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}

func GenerateNFibonacci(n int) chan int {
	out := make(chan int, n)
	go func() {
		defer close(out)
		i, j := 0, 1
		for k := 0; k < n; k++ {
			out <- i
			i, j = i+j, i
		}
	}()
	return out
}

func CollectSlice[T, U any](in []T, f func(in T) U) []U {
	out := make([]U, len(in))

	for i := range in {
		out[i] = f(in[i])
	}

	return out
}

func SelectSlice[T any](in []T, f func(in T) bool) []T {
	out := make([]T, 0, len(in))

	for _, e := range in {
		if f(e) {
			out = append(out, e)
		}
	}

	return out
}

func InjectSlice[T any](in []T, init T, f func(l, r T) T) T {
	acc := init

	for _, e := range in {
		acc = f(acc, e)
	}

	return acc
}

func CollectChan[T, U any](in chan T, f func(in T) U) chan U {
	out := make(chan U)
	go func() {
		defer close(out)
		for e := range in {
			out <- f(e)
		}
	}()
	return out
}

func SelectChan[T any](in chan T, f func(in T) bool) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for e := range in {
			if f(e) {
				out <- e
			}
		}
	}()

	return out
}

func InjectChan[T any](in chan T, init T, f func(l, r T) T) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		acc := init
		for e := range in {
			acc = f(acc, e)
		}
		out <- acc
	}()
	return out
}
