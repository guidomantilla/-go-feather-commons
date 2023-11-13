package streams

import (
	"github.com/jucardi/go-streams/v2/streams"
)

type DefaultStream[T comparable] struct {
	stream streams.IStream[T]
}

func (stream *DefaultStream[T]) SetThreads(threads int) Stream[T] {
	stream.stream.SetThreads(threads)
	return stream
}

func (stream *DefaultStream[T]) Filter(f ConditionalFunc[T]) Stream[T] {
	stream.stream.Filter(func(item T) bool {
		return f(item)
	})
	return stream
}

func (stream *DefaultStream[T]) Except(f ConditionalFunc[T]) Stream[T] {
	stream.stream.Except(func(item T) bool {
		return f(item)
	})
	return stream
}

func (stream *DefaultStream[T]) Sort(f SortFunc[T], desc ...bool) Stream[T] {
	stream.stream.Sort(func(a, b T) int {
		return f(a, b)
	}, desc...)
	return stream
}

func (stream *DefaultStream[T]) Distinct() Stream[T] {
	stream.stream.Distinct()
	return stream
}

func (stream *DefaultStream[T]) First(defaultValue ...T) T {
	return stream.stream.First(defaultValue...)
}

func (stream *DefaultStream[T]) Last(defaultValue ...T) T {
	return stream.stream.Last(defaultValue...)
}

func (stream *DefaultStream[T]) At(index int, defaultValue ...T) T {
	return stream.stream.At(index, defaultValue...)
}

func (stream *DefaultStream[T]) AtReverse(pos int, defaultValue ...T) T {
	return stream.stream.AtReverse(pos, defaultValue...)
}

func (stream *DefaultStream[T]) Count() int {
	return stream.stream.Count()
}

func (stream *DefaultStream[T]) IsEmpty() bool {
	return stream.stream.IsEmpty()
}

func (stream *DefaultStream[T]) Contains(value T) bool {
	return stream.stream.Contains(value)
}

func (stream *DefaultStream[T]) AnyMatch(f ConditionalFunc[T]) bool {
	return stream.stream.AnyMatch(func(item T) bool {
		return f(item)
	})
}

func (stream *DefaultStream[T]) AllMatch(f ConditionalFunc[T]) bool {
	return stream.stream.AllMatch(func(item T) bool {
		return f(item)
	})
}

func (stream *DefaultStream[T]) NotAllMatch(f ConditionalFunc[T]) bool {
	return stream.stream.NotAllMatch(func(item T) bool {
		return f(item)
	})
}

func (stream *DefaultStream[T]) NoneMatch(f ConditionalFunc[T]) bool {
	return stream.stream.NoneMatch(func(item T) bool {
		return f(item)
	})
}

func (stream *DefaultStream[T]) ForEach(f IterFunc[T]) {
	stream.stream.ForEach(func(item T) {
		f(item)
	})
}

func (stream *DefaultStream[T]) ParallelForEach(f IterFunc[T], threads int, skipWait ...bool) {
	stream.stream.ParallelForEach(func(item T) {
		f(item)
	}, threads, skipWait...)
}

func (stream *DefaultStream[T]) ToArray() []T {
	return stream.stream.ToArray()
}

func Build[T comparable](set []T, threads ...int) *DefaultStream[T] {
	return &DefaultStream[T]{
		stream: streams.FromArray[T](set, threads...),
	}
}
