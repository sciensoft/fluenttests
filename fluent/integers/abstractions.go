package integers

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
	"golang.org/x/exp/constraints"
)

// IComparable[T constraints.Integer] interface for the floats package. Following
// the Fluent-Interface design pattern to provide chained testing methods.
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

// FluentT[T constraints.Integer] is a type returned by calling "integers.Fluent[constraints.Integer](t)"
// for enabling easy-to-use testing methods.
//
// It carries the "*testing.T" object to be further used in the methods chain.
type FluentT[T constraints.Integer] struct {
	testingT *testing.T
}

// Subject[T constraints.Integer] is a type returned by calling "fluentObj.It(-100)"
// used to encapsulate the "*testing.T" object and testable subject.
//
// NOTE: The "It()" method in the example accepts any integer number as per constraint "constraints.Integer".
type Subject[T constraints.Integer] struct {
	*Testable[T]
}

// Additional[T constraints.Integer] is a type returned by some easy-to-use testing methods,
// providing basic additive test chaining with "And()".
type Additional[T constraints.Integer] struct {
	*Comparable[T]
}

// Comparable[T constraints.Integer] is a type returned by calling "fluentObj.It(-100).Should()",
// encapsulating a "*Testable[T]" object for further testing, and providing
// all the integers package's easy-to-use set of methods with signatures declared at
// the "integers.IComparable[T constraints.Integer]" interface.
//
// NOTE: The "It()" method in the example accepts any integer number as per constraint "constraints.Integer".
type Comparable[T constraints.Integer] struct {
	*Testable[T]
}

// Testable[T constraints.Integer] is a type that holds the "*testing.T" object and the underline
// Mock object, supporting further testing using the chain of easy-to-use methods.
// It's used by the "Subject[T constraints.Integer]" and "Comparable[T constraints.Integer]" types.
type Testable[T constraints.Integer] struct {
	testingT *testing.T
	value    T
}
