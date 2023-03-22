package strings

import (
	"testing"
)

type String interface {
	~string
}

type IFluent interface {
	It(subject string) ISubject
}

type FluentT struct {
	testingT *testing.T
}

type Testable struct {
	testingT *testing.T
	value    string
}

type ISubject interface {
	Should() IComparable
}

type Subject struct {
	*Testable
}

type IComparable interface {
	Be(comparable string) IAdditional
	NotBe(comparable string) IAdditional
	BeLowerThan(comparable string) IAdditional
	BeLowerThanOrEqualTo(comparable string) IAdditional
	BeGreaterThan(comparable string) IAdditional
	BeGreaterThanOrEqualTo(comparable string) IAdditional
}

type Comparable struct {
	*Testable
}

type IAdditional interface {
	And() IComparable
}

type Additional struct {
	*Comparable
}
