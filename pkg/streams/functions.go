package streams

import (
	"github.com/jucardi/go-streams/v2/streams"
)

func Map[From, To comparable](from Stream[From], f ConvertFunc[From, To]) Stream[To] {
	to := streams.Map[From, To](from.ToArray(), func(from From) To {
		return f(from)
	}).ToArray()
	return Build(to)
}
