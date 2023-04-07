package contracts

import (
	"reflect"
	"testing"

	"sciensoft.dev/testing/fluent"
)

type IComparable[T any] interface {
	fluent.IComparable[T]
	BeOfType(comparable reflect.Type) fluent.IAdditional[T, IComparable[T]]
	NotBeOfType(comparable reflect.Type) fluent.IAdditional[T, IComparable[T]]
	HaveMemberNamed(comparable string) fluent.IAdditional[T, IComparable[T]]
	HaveFieldNamed(comparable string) fluent.IAdditional[T, IComparable[T]]
	HaveMethodNamed(comparable string) fluent.IAdditional[T, IComparable[T]]
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

type MemberType int64

const (
	MemberTypeAll    MemberType = 0
	MemberTypeField  MemberType = 1
	MemberTypeMethod MemberType = 2
)
