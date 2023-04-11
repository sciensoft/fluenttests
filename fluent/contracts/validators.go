package contracts

import (
	"reflect"
	"testing"

	f "github.com/sciensoft/fluenttests/fluent"
)

func ofType[T any](t *testing.T, invert f.AdditiveInverse, value T, comparable reflect.Type, messagesf ...string) {
	vtype := reflect.TypeOf(value)
	isOneOf := vtype != comparable

	if invert {
		isOneOf = !isOneOf
	}

	if isOneOf {
		message := f.GetMessage("Expected type [%t] is not of type [%v].", messagesf...)
		t.Errorf(message, vtype, comparable.Name())
	}
}

func haveMembers[T any](t *testing.T, invert f.AdditiveInverse, matchType f.Match, mType f.MemberType, value T, comparable []string, messagesf ...string) {
	vtype := reflect.TypeOf(value)
	hasMember := false
	matchAll := matchType == f.MatchAll

	for _, member := range comparable {
		hasField := tryGetMemberByName(f.MemberTypeField, vtype, member)
		hasMethod := tryGetMemberByName(f.MemberTypeMethod, vtype, member)

		if matchAll || !hasMember {
			switch mType {
			case f.MemberTypeField:
				hasMember = hasField
			case f.MemberTypeMethod:
				hasMember = hasMethod
			default:
				hasMember = hasField || hasMethod
			}
		}

		if matchAll && !hasMember {
			break
		}
	}

	if invert {
		hasMember = !hasMember
	}

	if !hasMember {
		message := f.GetMessage("Object of [%t] does not have a member called [%v].", messagesf...)
		t.Errorf(message, vtype, comparable)
	}
}

func haveFieldWithTag[T any](t *testing.T, invert f.AdditiveInverse, value T, comparableFieldName string, comparableTagName string, messagesf ...string) {
	vtype := reflect.TypeOf(value)
	field, hasField := vtype.FieldByName(comparableFieldName)

	if hasField {
		_, hasTag := field.Tag.Lookup(comparableTagName)

		if !hasTag {
			message := f.GetMessage("Object of [%t] does not have a member [%v] with tag [%v].", messagesf...)
			t.Errorf(message, vtype, comparableFieldName, comparableTagName)
		}
	}
}

func haveAllFieldsWithTag[T any](t *testing.T, invert f.AdditiveInverse, value T, comparableTagName string, messagesf ...string) {
	vtype := reflect.TypeOf(value)
	numFields := vtype.NumField()

	for i := 0; i < numFields; i++ {
		field := vtype.Field(i)
		_, hasTag := field.Tag.Lookup(comparableTagName)

		if !hasTag {
			message := f.GetMessage("Object of [%t] does not have all fields with tag [%v].", messagesf...)
			t.Errorf(message, vtype, comparableTagName)
		}
	}
}

func tryGetMemberByName(mType f.MemberType, vtype reflect.Type, memberName string) bool {
	defer func() bool {
		recover()
		return false
	}()

	if mType == f.MemberTypeField {
		_, hasField := vtype.FieldByName(memberName)
		return hasField
	} else {
		_, hasMethod := vtype.MethodByName(memberName)
		return hasMethod
	}
}
