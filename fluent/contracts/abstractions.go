package contracts

import (
	"reflect"
	"testing"

	"sciensoft.dev/testing/fluent"
)

type IComparable[T any] interface {
	fluent.IComparable[T]
	BeOfType(typeName reflect.Type) fluent.IAdditional[T, IComparable[T]]
	NotBeOfType(typeName reflect.Type) fluent.IAdditional[T, IComparable[T]]
	HaveMember(memberName string) fluent.IAdditional[T, IComparable[T]]
	HaveField(fieldName string) fluent.IAdditional[T, IComparable[T]]
	HaveFieldWithTag(fieldName string, tagName string) fluent.IAdditional[T, IComparable[T]]
	HaveAllFieldsWithTag(tagName string) fluent.IAdditional[T, IComparable[T]]
	HaveMethod(methodName string) fluent.IAdditional[T, IComparable[T]]
	HaveAnyOfMembers(membersNames []string) fluent.IAdditional[T, IComparable[T]]
	HaveAllOfMembers(membersNames []string) fluent.IAdditional[T, IComparable[T]]
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
