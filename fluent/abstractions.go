package fluent

import "testing"

type IFluent[T any] interface {
	It(subject T) ISubject[T]
}

type ISubject[T any] interface {
	Should() IComparable[T]
}

type IComparable[T any] interface {
	Be(comparable T) IAdditional[T]
	NotBe(comparable T) IAdditional[T]
	BeLowerThan(comparable T) IAdditional[T]
	BeLowerThanOrEqualTo(comparable T) IAdditional[T]
	BeGreaterThan(comparable T) IAdditional[T]
	BeGreaterThanOrEqualTo(comparable T) IAdditional[T]
}

type IAdditional[T any] interface {
	And() IComparable[T]
}

type FluentT[T any] struct {
	testingT *testing.T
}

type Testable[T any] struct {
	TestingT *testing.T
	Value    T
}

type Subject[T any] struct {
	*Testable[T]
}

type Comparable[T any] struct {
	*Testable[T]
}

type Additional[T any] struct {
	*Comparable[T]
}
