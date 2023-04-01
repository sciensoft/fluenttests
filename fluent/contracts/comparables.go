package contracts

import (
	"reflect"

	"sciensoft.dev/testing/fluent"
)

func (s *Comparable[T]) BeOfType(expected reflect.Type) fluent.IAdditional[T, IComparable[T]] {
	beOfType(s.testingT, s.value, expected)
	return s.createAdditional()
}

func (s *Comparable[T]) createAdditional() fluent.IAdditional[T, IComparable[T]] {
	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}
