package integers

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
	"golang.org/x/exp/constraints"
)

type IComparable[T constraints.Integer] interface {
	fluent.IComparable[T]
	Be(value T) // If comparable to be a Value, no other comparison is required
	BeZero()    // If comparable to Zero, no other comparison is required
	BeNegative() fluent.IAdditional[T, IComparable[T]]
	BePositive() fluent.IAdditional[T, IComparable[T]]
	// BeOfType(t types.Type)
	NotBe(value T) fluent.IAdditional[T, IComparable[T]]
	NotBeZero() fluent.IAdditional[T, IComparable[T]]
	NotBeNegative() fluent.IAdditional[T, IComparable[T]]
	NotBePositive() fluent.IAdditional[T, IComparable[T]]
	BeLowerThan(value T) fluent.IAdditional[T, IComparable[T]]
	BeLowerThanOrEqualTo(value T) fluent.IAdditional[T, IComparable[T]]
	BeGreaterThan(value T) fluent.IAdditional[T, IComparable[T]]
	BeGreaterThanOrEqualTo(value T) fluent.IAdditional[T, IComparable[T]]
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
