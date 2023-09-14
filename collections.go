package functionalgo

type Collection[T any] []T

func MapCollection[T, U any](in Collection[T], f func(in T) U) Collection[U] {
	out := make(Collection[U], len(in))

	for i, e := range in {
		out[i] = f(e)
	}

	return out
}

func FilterCollection[T any](in Collection[T], f func(in T) bool) Collection[T] {
	out := make(Collection[T], 0, len(in))

	for _, e := range in {
		if f(e) {
			out = append(out, e)
		}
	}

	return out
}

func ReduceCollection[T any](in Collection[T], f func(l, r T) T, init T) T {
	acc := init

	for _, e := range in {
		acc = f(acc, e)
	}

	return acc
}

func CollectionToStream[T any](collection Collection[T]) Stream[T] {
	channel := make(chan T, len(collection))
	for _, e := range collection {
		channel <- e
	}
	close(channel)
	return Stream[T](channel)
}
