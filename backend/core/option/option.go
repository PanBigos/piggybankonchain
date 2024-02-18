package option

import "fmt"

type Option[T any] struct {
	set bool
	v   T
}

func (o Option[T]) String() string {
	if o.Some() {
		return fmt.Sprintf("%v", o.v)
	} else {
		return fmt.Sprintf("%v", nil)
	}
}

func Some[T any](v T) Option[T] {
	return Option[T]{
		set: true,
		v:   v,
	}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (p Option[T]) None() bool        { return !p.set }
func (p Option[T]) Some() bool        { return p.set }
func (p Option[T]) Unwrap() (T, bool) { return p.v, p.set }
func (p Option[T]) MustUnwrap() T {
	if !p.set {
		panic("tried to unwrap none")
	}
	return p.v
}
func (p Option[T]) Value() T {
	return p.v
}
