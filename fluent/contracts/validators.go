package contracts

import (
	"reflect"
	"testing"

	"sciensoft.dev/testing/fluent"
)

func beOfType[T any](t *testing.T, value T, expected reflect.Type, messagesf ...string) {
	vtype := reflect.TypeOf(value)

	if vtype != expected {
		message := fluent.GetMessage("Expected type [%t] is not of type [%v].", messagesf...)
		t.Errorf(message, vtype, expected.Name())
	}
}
