package floats

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
	"golang.org/x/exp/constraints"
)

func Fluent[T constraints.Float](t *testing.T) fluent.IFluent[T, IComparable[T]] {
	return &FluentT[T]{
		testingT: t,
	}
}

func (s *FluentT[T]) It(subject T) fluent.ISubject[T, IComparable[T]] {
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
