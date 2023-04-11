package strings

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
)

func Fluent(t *testing.T) fluent.IFluent[string, IComparable] {
	return &FluentT{
		testingT: t,
	}
}

func (s *FluentT) It(subject string) fluent.ISubject[string, IComparable] {
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
