package floats

import (
	"testing"

	"golang.org/x/exp/constraints"
)

type IFluent[T constraints.Float] interface {
	It(subject T) ISubject[T]
}

type ISubject[T constraints.Float] interface {
	Should() IComparable[T]
}

type IComparable[T constraints.Float] interface {
	Be(comparable T) // If comparable to be a Value, no other comparison is required
	BeZero()         // If comparable to Zero, no other comparison is required
	BeNegative() IAdditional[T]
	BePositive() IAdditional[T]
	// BeOfType(t types.Type)
	NotBe(comparable T) IAdditional[T]
	NotBeZero() IAdditional[T]
	NotBeNegative() IAdditional[T]
	NotBePositive() IAdditional[T]
	BeLowerThan(comparable T) IAdditional[T]
	BeLowerThanOrEqualTo(comparable T) IAdditional[T]
	BeGreaterThan(comparable T) IAdditional[T]
	BeGreaterThanOrEqualTo(comparable T) IAdditional[T]
}

type IAdditional[T constraints.Float] interface {
	And() IComparable[T]
}

type FluentT[T constraints.Float] struct {
	testingT *testing.T
}

type Testable[T constraints.Float] struct {
	testingT *testing.T
	value    T
}

type Subject[T constraints.Float] struct {
	*Testable[T]
}

type Comparable[T constraints.Float] struct {
	*Testable[T]
}

type Additional[T constraints.Float] struct {
	*Comparable[T]
}
