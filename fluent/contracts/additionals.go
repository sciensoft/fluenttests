package contracts

import (
	"reflect"

	"github.com/sciensoft/fluenttests/fluent"
)

// OfType is an extra method used to support assertion by type and made
// available through some easy-to-use testing methods.
func (a *AdditionalWith[T]) OfType(ctype reflect.Type) IAdditionalWith[T] {
	param := a.params[0]

	if param.ptype == fluent.ParamType.Field {
		vtype := reflect.ValueOf(a.Comparable.value)
		field := vtype.FieldByName(param.value.(string))

		ofType(a.testingT, fluent.NegativeDefault, field.Interface(), ctype)
	}

	return a
}

// WithValue is an extra method used to support assertion by a specific value
// and made available through some easy-to-use testing methods.
func (a *AdditionalWith[T]) WithValue(value any) IAdditionalWith[T] {
	param := a.params[0]

	if param.ptype == fluent.ParamType.Field {
		vtype := reflect.ValueOf(a.Comparable.value)
		field := vtype.FieldByName(param.value.(string))
		fieldValue := field.Interface()

		if fieldValue != value {
			a.testingT.Errorf("Member value [%v] differs from [%v].", fieldValue, value)
		}
	}

	return a
}
