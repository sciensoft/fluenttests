package contracts

import (
	"reflect"

	"github.com/sciensoft/fluenttests/fluent"
	f "github.com/sciensoft/fluenttests/fluent"
)

func (c *Comparable[T]) BeOfType(typeName reflect.Type) f.IAdditional[T, IComparable[T]] {
	ofType(c.testingT, f.NegativeDefault, c.value, typeName)
	return c.createAdditional()
}

func (c *Comparable[T]) NotBeOfType(typeName reflect.Type) f.IAdditional[T, IComparable[T]] {
	ofType(c.testingT, f.NegativeInvert, c.value, typeName, "Type [%T] is to not be of type [%v].")
	return c.createAdditional()
}

func (c *Comparable[T]) HaveMember(memberName string) f.IAdditional[T, IComparable[T]] {
	haveMembers(c.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeAll, c.value, []string{memberName})
	return c.createAdditional()
}

func (c *Comparable[T]) HaveField(fieldName string) IAdditionalWith[T] {
	haveMembers(c.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeField, c.value, []string{fieldName}, "Object of [%T] does not have a field called [%v].")

	return c.createAdditionalWith(Params{
		ptype: fluent.ParamType.Field,
		value: fieldName,
	})
}

func (c *Comparable[T]) HaveFieldWithTag(fieldName string, tagName string) f.IAdditional[T, IComparable[T]] {
	haveFieldWithTag(c.testingT, f.NegativeDefault, c.value, fieldName, tagName, "....")
	return c.createAdditional()
}

func (c *Comparable[T]) HaveAllFieldsWithTag(tagName string) f.IAdditional[T, IComparable[T]] {
	haveAllFieldsWithTag(c.testingT, f.NegativeDefault, c.value, tagName, "....")
	return c.createAdditional()
}

func (c *Comparable[T]) HaveMethod(methodName string) f.IAdditional[T, IComparable[T]] {
	haveMembers(c.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeMethod, c.value, []string{methodName}, "Object of [%T] does not have a method called [%v].")
	return c.createAdditional()
}

func (c *Comparable[T]) HaveAnyOfMembers(membersNames []string) f.IAdditional[T, IComparable[T]] {
	haveMembers(c.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeAll, c.value, membersNames, "Object of [%T] does not have any of members [%v].")
	return c.createAdditional()
}

func (c *Comparable[T]) HaveAllOfMembers(membersNames []string) f.IAdditional[T, IComparable[T]] {
	haveMembers(c.testingT, f.NegativeDefault, f.MatchAll, f.MemberTypeAll, c.value, membersNames, "Object of [%T] does not have all of members [%v].")
	return c.createAdditional()
}

func (c *Comparable[T]) createAdditional() f.IAdditional[T, IComparable[T]] {
	additional := &Additional[T]{
		Comparable: c,
	}

	return additional
}

func (c *Comparable[T]) createAdditionalWith(params ...Params) IAdditionalWith[T] {
	additional := &AdditionalWith[T]{
		Additional: &Additional[T]{
			Comparable: c,
		},
		params: params,
	}

	return additional
}
