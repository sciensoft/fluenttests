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

type ISubject interface {
	Should() IComparable
}

type IComparable interface {
	Be(comparable string) IAdditional
	BeEmpty()
	BeOneOf(comparables []string) IAdditional
	NotBe(comparable string) IAdditional
	NotBeEmpty() IAdditional
	NotBeOneOf(comparables []string) IAdditional
	Match(pattern string) IAdditional
	NotMatch(pattern string) IAdditional
	StartWith(prefix string) IAdditional
	NotStartWith(prefix string) IAdditional
	EndWith(comparable string) IAdditional
	NotEndWith(comparable string) IAdditional
	HaveLengthOf(length int) IAdditional

	// TODO : Review
	// BeLowerThan(comparable string) IAdditional
	// BeLowerThanOrEqualTo(comparable string) IAdditional
	// BeGreaterThan(comparable string) IAdditional
	// BeGreaterThanOrEqualTo(comparable string) IAdditional
}

type IAdditional interface {
	And() IComparable
}

type FluentT struct {
	testingT *testing.T
}

type Testable struct {
	testingT *testing.T
	value    string
}

type Subject struct {
	*Testable
}

type Comparable struct {
	*Testable
}

type Additional struct {
	*Comparable
}
