package functionalgo

type Dictionary[T, U comparable] map[T]U

func MapDictionary[T, U, V comparable](in Dictionary[T, U], f func(in U) V) Dictionary[T, V] {
	out := make(Dictionary[T, V])
	for k, v := range in {
		out[k] = f(v)
	}
	return out
}

func FilterDictionary[T, U comparable](in Dictionary[T, U], f func(in U) bool) Dictionary[T, U] {
	out := in
	for k, v := range out {
		if !f(v) {
			delete(out, k)
		}
	}
	return out
}

func ReduceDictionary[T, U comparable](in Dictionary[T, U], f func(l, r U) U, init U) U {
	acc := init
	for _, v := range in {
		acc = f(acc, v)
	}
	return acc
}

func DictionaryToStream[T, U comparable](dictionary Dictionary[T, U]) Stream[U] {
	channel := make(chan U, len(dictionary))
	for _, v := range dictionary {
		channel <- v
	}
	close(channel)
	return Stream[U](channel)
}

func DictionaryToCollection[T, U comparable](dictionary Dictionary[T, U]) Collection[U] {
	collection := Collection[U](make([]U, 0, len(dictionary)))
	for _, v := range dictionary {
		collection = append(collection, v)
	}
	return collection
}
