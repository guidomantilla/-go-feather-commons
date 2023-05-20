package util

func ValueToPtr[T any](value T) *T {
	return &value
}
