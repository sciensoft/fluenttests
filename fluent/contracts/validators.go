package contracts

import (
	"reflect"
	"testing"

	"sciensoft.dev/testing/fluent"
)

func ofType[T any](t *testing.T, invert bool, value T, comparable reflect.Type, messagesf ...string) {
	vtype := reflect.TypeOf(value)
	isOneOf := vtype != comparable

	if invert {
		isOneOf = !isOneOf
	}

	if isOneOf {
		message := fluent.GetMessage("Expected type [%t] is not of type [%v].", messagesf...)
		t.Errorf(message, vtype, comparable.Name())
	}
}

func hasMember[T any](t *testing.T, invert bool, mType MemberType, value T, comparable string, messagesf ...string) {
	vtype := reflect.TypeOf(value)
	hasMember := false

	hasField := tryGetMemberByName(MemberTypeField, vtype, comparable)
	hasMethod := tryGetMemberByName(MemberTypeMethod, vtype, comparable)

	switch mType {
	case MemberTypeField:
		hasMember = hasField
	case MemberTypeMethod:
		hasMember = hasMethod
	default:
		hasMember = hasField || hasMethod
	}

	if invert {
		hasMember = !hasMember
	}

	if !hasMember {
		message := fluent.GetMessage("Expected object of [%t] does not have a member called [%v].", messagesf...)
		t.Errorf(message, vtype, comparable)
	}
}

func tryGetMemberByName(mType MemberType, vtype reflect.Type, memberName string) bool {
	defer func() bool {
		recover()
		return false
	}()

	if mType == MemberTypeField {
		_, hasField := vtype.FieldByName(memberName)
		return hasField
	} else {
		_, hasMethod := vtype.MethodByName(memberName)
		return hasMethod
	}
}
