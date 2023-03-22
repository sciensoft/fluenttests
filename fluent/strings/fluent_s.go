package strings

import (
	"testing"
)

func Fluent(t *testing.T) IFluent {
	return &FluentT{
		testingT: t,
	}
}

func (s *FluentT) It(subject string) ISubject {
	sub := Subject{
		Testable: &Testable{},
	}
	sub.testingT = s.testingT
	sub.value = subject

	return &sub
}

func (s *Subject) Should() IComparable {
	return &Comparable{
		Testable: s.Testable,
	}
}

func (s *Additional) And() IComparable {
	return s.Comparable
}
