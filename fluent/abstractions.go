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
