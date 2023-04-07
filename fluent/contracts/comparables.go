package contracts

import (
	"reflect"

	"sciensoft.dev/testing/fluent"
)

func (s *Comparable[T]) BeOfType(comparable reflect.Type) fluent.IAdditional[T, IComparable[T]] {
	ofType(s.testingT, false, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) NotBeOfType(comparable reflect.Type) fluent.IAdditional[T, IComparable[T]] {
	ofType(s.testingT, true, s.value, comparable, "Expected type [%t] is to not be of type [%v].")
	return s.createAdditional()
}

func (s *Comparable[T]) HaveMemberNamed(comparable string) fluent.IAdditional[T, IComparable[T]] {
	hasMember(s.testingT, false, MemberTypeAll, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) HaveFieldNamed(comparable string) fluent.IAdditional[T, IComparable[T]] {
	hasMember(s.testingT, false, MemberTypeField, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) HaveMethodNamed(comparable string) fluent.IAdditional[T, IComparable[T]] {
	hasMember(s.testingT, false, MemberTypeMethod, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) createAdditional() fluent.IAdditional[T, IComparable[T]] {
	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}
