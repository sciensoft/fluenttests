package integers

import (
	"testing"

	"golang.org/x/exp/constraints"
	"sciensoft.dev/testing/fluent"
)

type IComparable[T constraints.Integer] interface {
	fluent.IComparable[T]
	Be(comparable T) // If comparable to be a Value, no other comparison is required
	BeZero()         // If comparable to Zero, no other comparison is required
	BeNegative() fluent.IAdditional[T, IComparable[T]]
	BePositive() fluent.IAdditional[T, IComparable[T]]
	// BeOfType(t types.Type)
	NotBe(comparable T) fluent.IAdditional[T, IComparable[T]]
	NotBeZero() fluent.IAdditional[T, IComparable[T]]
	NotBeNegative() fluent.IAdditional[T, IComparable[T]]
	NotBePositive() fluent.IAdditional[T, IComparable[T]]
	BeLowerThan(comparable T) fluent.IAdditional[T, IComparable[T]]
	BeLowerThanOrEqualTo(comparable T) fluent.IAdditional[T, IComparable[T]]
	BeGreaterThan(comparable T) fluent.IAdditional[T, IComparable[T]]
	BeGreaterThanOrEqualTo(comparable T) fluent.IAdditional[T, IComparable[T]]
}

type FluentT[T constraints.Integer] struct {
	testingT *testing.T
}

type Subject[T constraints.Integer] struct {
	*Testable[T]
}

type Additional[T constraints.Integer] struct {
	*Comparable[T]
}

type Comparable[T constraints.Integer] struct {
	*Testable[T]
}

type Testable[T constraints.Integer] struct {
	testingT *testing.T
	value    T
}
