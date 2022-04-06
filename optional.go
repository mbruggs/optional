// Package optional provides a generic optional type
package optional

// Optional either holds a value or not.
type Optional[T any] struct {
	value *T
}

// New creates a new Optional type with the given value.
func New[T any](value T) Optional[T] {
	return Optional[T]{
		value: &value,
	}
}

// Empty creates an empty Optional type.
func Empty[T any]() Optional[T] {
	return Optional[T]{
		value: nil,
	}
}

// Value returns the value in the Optional and whether the value was set.
//
// If the value is unset, the zero value is used.
func (o Optional[T]) Value() (T, bool) {
	var zero T
	if o.value == nil {
		return zero, false
	}

	return *o.value, true
}

// Valid returns if the value has been set.
func (o Optional[T]) Valid() bool {
	return o.value != nil
}
