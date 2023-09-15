package functionalgo

func MakeSlice[T any](items ...T) []T {
	out := make([]T, len(items))
	copy(out, items)
	return out
}

func MapSlice[T, U any](in []T, f func(in T) U) []U {
	out := make([]U, len(in))

	for i, e := range in {
		out[i] = f(e)
	}

	return out
}

func FilterSlice[T any](in []T, f func(in T) bool) []T {
	out := make([]T, 0, len(in))

	for _, e := range in {
		if f(e) {
			out = append(out, e)
		}
	}

	return out
}

func ReduceSlice[T any](in []T, f func(l, r T) T, init T) T {
	acc := init

	for _, e := range in {
		acc = f(acc, e)
	}

	return acc
}

func SliceToChan[T any](collection []T) chan T {
	channel := make(chan T, len(collection))
	for _, e := range collection {
		channel <- e
	}
	close(channel)
	return channel
}
