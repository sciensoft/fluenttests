package strings

import (
	"testing"

	"sciensoft.dev/testing/fluent"
)

type String interface {
	~string
}

type IComparable interface {
	fluent.IComparable[string]
	Be(comparable string) fluent.IAdditional[string, IComparable]
	BeEmpty()
	BeOneOf(comparables []string) fluent.IAdditional[string, IComparable]
	NotBe(comparable string) fluent.IAdditional[string, IComparable]
	NotBeEmpty() fluent.IAdditional[string, IComparable]
	NotBeOneOf(comparables []string) fluent.IAdditional[string, IComparable]
	Match(pattern string) fluent.IAdditional[string, IComparable]
	NotMatch(pattern string) fluent.IAdditional[string, IComparable]
	StartWith(prefix string) fluent.IAdditional[string, IComparable]
	NotStartWith(prefix string) fluent.IAdditional[string, IComparable]
	EndWith(comparable string) fluent.IAdditional[string, IComparable]
	NotEndWith(comparable string) fluent.IAdditional[string, IComparable]
	HaveLengthOf(length int) fluent.IAdditional[string, IComparable]
	HaveFnv32aSumOf(comparable uint32) fluent.IAdditional[string, IComparable]
	HaveFnv64aSumOf(comparable uint64) fluent.IAdditional[string, IComparable]
}

type FluentT struct {
	testingT *testing.T
}

type Subject struct {
	*Testable
}

type Additional struct {
	*Comparable
}

type Comparable struct {
	*Testable
}

type Testable struct {
	testingT *testing.T
	value    string
}
