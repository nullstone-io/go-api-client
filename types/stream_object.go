package types

type StreamObject[T any] struct {
	Object T
	Err    error
}
