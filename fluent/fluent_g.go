package fluent

import (
	"testing"
)

func Fluent[T any](t *testing.T) IFluent[T] {
	return &FluentT[T]{
		testingT: t,
	}
}

func (s *FluentT[T]) It(subject T) ISubject[T] {
	sub := Subject[T]{
		Testable: &Testable[T]{},
	}
	sub.testingT = s.testingT
	sub.value = subject

	return &sub
}

func (s *Subject[T]) Should() IComparable[T] {
	return &Comparable[T]{
		Testable: s.Testable,
	}
}

func (s *Additional[T]) And() IComparable[T] {
	return s.Comparable
}
