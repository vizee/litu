package maps

import "litu/pair"

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Pairs[K comparable, V any](m map[K]V) []pair.Pair[K, V] {
	pairs := make([]pair.Pair[K, V], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, pair.Pair[K, V]{A: k, B: v})
	}
	return pairs
}
