package functionalgo

type Stream[T any] chan T

func MapStream[T, U any](in Stream[T], f func(in T) U) Stream[U] {
	out := make(Stream[U])
	go func() {
		defer close(out)
		for e := range in {
			out <- f(e)
		}
	}()
	return out
}

func FilterStream[T any](in Stream[T], f func(in T) bool) Stream[T] {
	out := make(Stream[T])
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

func ReduceStream[T any](in Stream[T], f func(l, r T) T, init T) Stream[T] {
	out := make(Stream[T])
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

func FinalValue[T any](in Stream[T]) T {
	var acc T
	for e := range in {
		acc = e
	}
	return acc
}

func StreamToCollection[T any](stream Stream[T]) chan Collection[T] {
	slice := make([]T, 0)
	out := make(chan Collection[T])
	go func() {
		defer close(out)
		for e := range stream {
			slice = append(slice, e)
		}
		out <- slice
	}()
	return out
}
