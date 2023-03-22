package integers

import (
	"testing"

	"golang.org/x/exp/constraints"
)

type IFluent[T constraints.Integer] interface {
	It(subject T) ISubject[T]
}

type FluentT[T constraints.Integer] struct {
	testingT *testing.T
}

type Testable[T constraints.Integer] struct {
	testingT *testing.T
	value    T
}

type ISubject[T constraints.Integer] interface {
	Should() IComparable[T]
}

type Subject[T constraints.Integer] struct {
	*Testable[T]
}

type IComparable[T constraints.Integer] interface {
	Be(comparable T) IAdditional[T]
	NotBe(comparable T) IAdditional[T]
	BeLowerThan(comparable T) IAdditional[T]
	BeLowerThanOrEqualTo(comparable T) IAdditional[T]
	BeGreaterThan(comparable T) IAdditional[T]
	BeGreaterThanOrEqualTo(comparable T) IAdditional[T]
}

type Comparable[T constraints.Integer] struct {
	*Testable[T]
}

type IAdditional[T constraints.Integer] interface {
	And() IComparable[T]
}

type Additional[T constraints.Integer] struct {
	*Comparable[T]
}
