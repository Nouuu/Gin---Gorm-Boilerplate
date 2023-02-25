package optional

type Optional[T any] struct {
	value  T
	exists bool
}

func (o Optional[T]) IsPresent() bool {
	return o.exists
}

func (o Optional[T]) IsEmpty() bool {
	return !o.exists
}

func (o Optional[T]) Get() T {
	return o.value
}

func (o Optional[T]) OrElse(defaultValue T) T {
	if o.exists {
		return o.value
	}
	return defaultValue
}

func (o Optional[T]) OrElseGet(defaultValue func() T) T {
	if o.exists {
		return o.value
	}
	return defaultValue()
}

func (o Optional[T]) OrElsePanic(panicMessage string) T {
	if o.exists {
		return o.value
	}
	panic(panicMessage)
}

func Empty[T any]() Optional[T] {
	return Optional[T]{exists: false}
}

func Of[T any](value T) Optional[T] {
	return Optional[T]{value: value, exists: true}
}

func OfNillable[T any](value *T) Optional[T] {
	if value == nil {
		return Optional[T]{exists: false}
	}
	return Optional[T]{value: *value, exists: true}
}
