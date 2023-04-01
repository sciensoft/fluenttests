package contracts

import (
	"reflect"
	"testing"

	"sciensoft.dev/testing/fluent"
)

type IComparable[T any] interface {
	fluent.IComparable[T]
	BeOfType(expected reflect.Type) fluent.IAdditional[T, IComparable[T]]
	// Add more methods here
}

type FluentT[T any] struct {
	testingT *testing.T
}

type Subject[T any] struct {
	*Testable[T]
}

type Additional[T any] struct {
	*Comparable[T]
}

type Comparable[T any] struct {
	*Testable[T]
}

type Testable[T any] struct {
	testingT *testing.T
	value    T
}
