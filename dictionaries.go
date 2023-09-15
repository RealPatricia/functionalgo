package functionalgo

func MapMap[T, U, V comparable](in map[T]U, f func(in U) V) map[T]V {
	out := make(map[T]V)
	for k, v := range in {
		out[k] = f(v)
	}
	return out
}

func FilterMap[T, U comparable](in map[T]U, f func(in U) bool) map[T]U {
	out := in
	for k, v := range out {
		if !f(v) {
			delete(out, k)
		}
	}
	return out
}

func ReduceMap[T, U comparable](in map[T]U, f func(l, r U) U, init U) U {
	acc := init
	for _, v := range in {
		acc = f(acc, v)
	}
	return acc
}

func MapToChan[T, U comparable](dictionary map[T]U) chan U {
	channel := make(chan U, len(dictionary))
	for _, v := range dictionary {
		channel <- v
	}
	close(channel)
	return channel
}

func MapToSlice[T, U comparable](dictionary map[T]U) []U {
	collection := make([]U, 0, len(dictionary))
	for _, v := range dictionary {
		collection = append(collection, v)
	}
	return collection
}
