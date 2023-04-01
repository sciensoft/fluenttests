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
