package contracts

import (
	"reflect"

	f "github.com/sciensoft/fluenttests/fluent"
)

func (s *Comparable[T]) BeOfType(typeName reflect.Type) f.IAdditional[T, IComparable[T]] {
	ofType(s.testingT, f.NegativeDefault, s.value, typeName)
	return s.createAdditional()
}

func (s *Comparable[T]) NotBeOfType(typeName reflect.Type) f.IAdditional[T, IComparable[T]] {
	ofType(s.testingT, f.NegativeInvert, s.value, typeName, "Type [%t] is to not be of type [%v].")
	return s.createAdditional()
}

func (s *Comparable[T]) HaveMember(memberName string) f.IAdditional[T, IComparable[T]] {
	haveMembers(s.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeAll, s.value, []string{memberName})
	return s.createAdditional()
}

func (s *Comparable[T]) HaveField(fieldName string) f.IAdditional[T, IComparable[T]] {
	haveMembers(s.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeField, s.value, []string{fieldName}, "Object of [%t] does not have a field called [%v].")
	return s.createAdditional()
}

func (s *Comparable[T]) HaveFieldWithTag(fieldName string, tagName string) f.IAdditional[T, IComparable[T]] {
	haveFieldWithTag(s.testingT, f.NegativeDefault, s.value, fieldName, tagName, "....")
	return s.createAdditional()
}

func (s *Comparable[T]) HaveAllFieldsWithTag(tagName string) f.IAdditional[T, IComparable[T]] {
	haveAllFieldsWithTag(s.testingT, f.NegativeDefault, s.value, tagName, "....")
	return s.createAdditional()
}

func (s *Comparable[T]) HaveMethod(methodName string) f.IAdditional[T, IComparable[T]] {
	haveMembers(s.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeMethod, s.value, []string{methodName}, "Object of [%t] does not have a method called [%v].")
	return s.createAdditional()
}

func (s *Comparable[T]) HaveAnyOfMembers(membersNames []string) f.IAdditional[T, IComparable[T]] {
	haveMembers(s.testingT, f.NegativeDefault, f.MatchAny, f.MemberTypeAll, s.value, membersNames, "Object of [%t] does not have any of members [%v].")
	return s.createAdditional()
}

func (s *Comparable[T]) HaveAllOfMembers(membersNames []string) f.IAdditional[T, IComparable[T]] {
	haveMembers(s.testingT, f.NegativeDefault, f.MatchAll, f.MemberTypeAll, s.value, membersNames, "Object of [%t] does not have all of members [%v].")
	return s.createAdditional()
}

func (s *Comparable[T]) createAdditional() f.IAdditional[T, IComparable[T]] {
	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}
