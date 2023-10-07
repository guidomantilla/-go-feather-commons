package collections

import "math/rand"

// func ToMap[ID IDAllowedTypes, M ValueAllowedTypes](items ...Item[ID, M]) map[ID]*M {
func ToMap[ID IDAllowedTypes, M ValueAllowedTypes](items ...Item[ID, M]) map[ID]*M {

	if len(items) == 0 {
		return make(map[ID]*M)
	}

	modelsMap := make(map[ID]*M)
	for _, m := range items {
		if m.ID() == nil {
			continue
		}
		modelsMap[*m.ID()] = m.Value()
	}

	return modelsMap
}

func Filter[V ValueAllowedTypes](collection []V, predicate func(item V, index int) bool) []V {
	result := make([]V, 0, len(collection))

	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}

func Map[T ValueAllowedTypes, R ValueAllowedTypes](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

func Reduce[T ValueAllowedTypes, R ValueAllowedTypes](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for i, item := range collection {
		initial = accumulator(initial, item, i)
	}

	return initial
}

func Shuffle[T ValueAllowedTypes](collection []T) []T {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})

	return collection
}

func Deduplicate[T ComparableAllowedTypes](collection []T) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

func Count[T ValueAllowedTypes](collection []T, predicate func(item T, index int) bool) int {
	var count int
	for index, item := range collection {
		if predicate(item, index) {
			count++
		}
	}

	return count
}

func Occurs[T ComparableAllowedTypes](collection []T) map[T]int {
	result := make(map[T]int)

	for _, item := range collection {
		result[item]++
	}

	return result
}

func Replace[T ComparableAllowedTypes](collection []T, old T, new T) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	for i := range result {
		if result[i] == old {
			result[i] = new
		}
	}

	return result
}

func Keys[K ComparableAllowedTypes, V ValueAllowedTypes](in map[K]V) []K {
	result := make([]K, 0, len(in))

	for k := range in {
		result = append(result, k)
	}

	return result
}

func Values[K ComparableAllowedTypes, V ValueAllowedTypes](in map[K]V) []V {
	result := make([]V, 0, len(in))

	for _, v := range in {
		result = append(result, v)
	}

	return result
}

func SelectByKeys[K ComparableAllowedTypes, V ValueAllowedTypes](in map[K]V, keys []K) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if Contains(keys, k) {
			r[k] = v
		}
	}
	return r
}

func Contains[T ComparableAllowedTypes](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}
