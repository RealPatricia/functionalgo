package functionalgo

func MakeChan[T any](items ...T) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, e := range items {
			out <- e
		}
	}()
	return out
}

func MapChan[T, U any](in chan T, f func(in T) U) chan U {
	out := make(chan U)
	go func() {
		defer close(out)
		for e := range in {
			out <- f(e)
		}
	}()
	return out
}

func FilterChan[T any](in chan T, f func(in T) bool) chan T {
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

func ReduceChan[T any](in chan T, f func(l, r T) T, init T) chan T {
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

func FinalValue[T any](in chan T) T {
	var acc T
	for e := range in {
		acc = e
	}
	return acc
}

func ChanToSlice[T any](stream chan T) chan []T {
	slice := make([]T, 0)
	out := make(chan []T)
	go func() {
		defer close(out)
		for e := range stream {
			slice = append(slice, e)
		}
		out <- slice
	}()
	return out
}
