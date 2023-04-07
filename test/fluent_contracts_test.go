package tdd

import (
	"reflect"
	"testing"

	"sciensoft.dev/testing/fluent/contracts"
)

func TestFluentContractsShouldBeOfType(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj := struct {
		A int
	}{
		A: 1,
	}
	rtype := reflect.TypeOf(obj)

	// Assert
	fluent.It(obj).
		Should().BeOfType(rtype)
}

func TestFluentContractsShouldNotBeOfType(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj1 := struct{ A int }{A: 1}
	obj2 := struct{ B int }{B: 1}
	rtype := reflect.TypeOf(obj2)

	// Assert
	fluent.It(obj1).
		Should().NotBeOfType(rtype)
}

func TestFluentContractsShouldHaveMemberNamed(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj := struct{ Name string }{Name: "Robert Griesemer"}

	// Assert
	fluent.It(obj).
		Should().HaveMemberNamed("Name")
}

func TestFluentContractsShouldHaveFieldNamed(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj := struct{ Value float32 }{Value: 105.4}

	// Assert
	fluent.It(obj).
		Should().HaveFieldNamed("Value")
}

func TestFluentContractsShouldHaveMethodNamed(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	// Assert
	fluent.It(fluent).
		Should().HaveMethodNamed("It")
}
