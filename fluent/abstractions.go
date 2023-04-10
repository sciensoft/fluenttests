package fluent

import "testing"

type IFluent[T any, Y IComparable[T]] interface {
	It(subject T) ISubject[T, Y]
}

type ISubject[T any, Y IComparable[T]] interface {
	Should() Y
}

type IAdditional[T any, Y IComparable[T]] interface {
	And() Y
}

type IComparable[T any] interface{}

type IValidator[TValue any, TExpected any] interface {
	Validate(t *testing.T, value TValue, expected TExpected, messagesf ...string)
}

/*
 * Enums
 */
type MemberType int64

const (
	MemberTypeAll    MemberType = 0
	MemberTypeField  MemberType = 1
	MemberTypeMethod MemberType = 2
)

type AdditiveInverse bool

const (
	NegativeDefault AdditiveInverse = false
	NegativeInvert  AdditiveInverse = true
)

type Match int

const (
	MatchAny Match = 1
	MatchAll Match = 2
)
