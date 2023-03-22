package fluent

import "testing"

type IFluent[T any] interface {
	It(subject T) ISubject[T]
}

type FluentT[T any] struct {
	testingT *testing.T
}

type Testable[T any] struct {
	testingT *testing.T
	value    T
}

type ISubject[T any] interface {
	Should() IComparable[T]
}

type Subject[T any] struct {
	*Testable[T]
}

type IComparable[T any] interface {
	Be(comparable T) IAdditional[T]
	NotBe(comparable T) IAdditional[T]
	BeLowerThan(comparable T) IAdditional[T]
	BeLowerThanOrEqualTo(comparable T) IAdditional[T]
	BeGreaterThan(comparable T) IAdditional[T]
	BeGreaterThanOrEqualTo(comparable T) IAdditional[T]
}

type Comparable[T any] struct {
	*Testable[T]
}

type IAdditional[T any] interface {
	And() IComparable[T]
}

type Additional[T any] struct {
	*Comparable[T]
}

/// OLD

// type ITestable[T any] interface {
// 	Should() *Comparable[T]
// 	// And() IComparable[T]
// }

// type ITestableAdditional[T any] interface {
// 	And() IComparable[T]
// }

// type ITestable[T any] interface {
// 	NewTestable[T]
// 	Should() IComparable[T]
// 	// And() IComparable[T]
// }

// type IComparable[T any] interface {
// 	GetValue() T
// 	BeGreaterThan(v T) ITestableAdditional[T]
// }

// type NewTestable[T any] struct {
// }

// type Testable[T any] struct {
// 	testingT *testing.T
// 	Value    T
// }

// type Comparable[T any] struct {
// 	Testable *INewTestable[T]
// }
