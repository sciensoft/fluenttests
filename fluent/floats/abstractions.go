package floats

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
	"golang.org/x/exp/constraints"
)

// IComparable[T constraints.Float] interface for the floats package. Following
// the Fluent-Interface design pattern to provide chained testing methods.
type IComparable[T constraints.Float] interface {
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

// FluentT[T constraints.Float] is a type returned by calling "contracts.Fluent[constraints.Float](t)"
// for enabling easy-to-use testing methods.
//
// It carries the "*testing.T" object to be further used in the methods chain.
type FluentT[T constraints.Float] struct {
	testingT *testing.T
}

// Subject[T constraints.Float] is a type returned by calling "fluentObj.It(0.9999)"
// used to encapsulate the "*testing.T" object and testable subject.
//
// NOTE: The "It()" method in the example accepts any float number as per constraint "constraints.Float".
type Subject[T constraints.Float] struct {
	*Testable[T]
}

// Additional[T constraints.Float] is a type returned by some easy-to-use testing methods,
// providing basic additive test chaining with "And()".
type Additional[T constraints.Float] struct {
	*Comparable[T]
}

// Comparable[T constraints.Float] is a type returned by calling "fluentObj.It(0.9999).Should()",
// encapsulating a "*Testable[T]" object for further testing, and providing
// all the floats package's easy-to-use set of methods with signatures declared at
// the "floats.IComparable[T constraints.Float]" interface.
//
// NOTE: The "It()" method in the example accepts any float number as per constraint "constraints.Float".
type Comparable[T constraints.Float] struct {
	*Testable[T]
}

// Testable[T constraints.Float] is a type that holds the "*testing.T" object and the underline
// Mock object, supporting further testing using the chain of easy-to-use methods.
// It's used by the "Subject[T constraints.Float]" and "Comparable[T constraints.Float]" types.
type Testable[T constraints.Float] struct {
	testingT *testing.T
	value    T
}
