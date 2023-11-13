package streams

type ConditionalFunc[T comparable] func(T) bool

type ConvertFunc[From, To any] func(From) To

type IterFunc[T any] func(T)

type SortFunc[T comparable] func(T, T) int

type Stream[T comparable] interface {
	SetThreads(threads int) Stream[T]
	Filter(f ConditionalFunc[T]) Stream[T]
	Except(f ConditionalFunc[T]) Stream[T]
	Sort(f SortFunc[T], desc ...bool) Stream[T]
	Distinct() Stream[T]
	First(defaultValue ...T) T
	Last(defaultValue ...T) T
	At(index int, defaultValue ...T) T
	AtReverse(pos int, defaultValue ...T) T
	Count() int
	IsEmpty() bool
	Contains(value T) bool
	AnyMatch(f ConditionalFunc[T]) bool
	AllMatch(f ConditionalFunc[T]) bool
	NotAllMatch(f ConditionalFunc[T]) bool
	NoneMatch(f ConditionalFunc[T]) bool
	ForEach(f IterFunc[T])
	ParallelForEach(f IterFunc[T], threads int, skipWait ...bool)
	ToArray() []T
}
