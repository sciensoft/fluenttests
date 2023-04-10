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
	Be(value string) fluent.IAdditional[string, IComparable]
	BeEmpty()
	BeOneOf(values []string) fluent.IAdditional[string, IComparable]
	NotBe(value string) fluent.IAdditional[string, IComparable]
	NotBeEmpty() fluent.IAdditional[string, IComparable]
	NotBeOneOf(values []string) fluent.IAdditional[string, IComparable]
	Match(pattern string) fluent.IAdditional[string, IComparable]
	NotMatch(pattern string) fluent.IAdditional[string, IComparable]
	StartWith(prefix string) fluent.IAdditional[string, IComparable]
	NotStartWith(prefix string) fluent.IAdditional[string, IComparable]
	EndWith(suffix string) fluent.IAdditional[string, IComparable]
	NotEndWith(suffix string) fluent.IAdditional[string, IComparable]
	HaveLengthOf(length int) fluent.IAdditional[string, IComparable]
	HaveFnv32aSumOf(sumValue uint32) fluent.IAdditional[string, IComparable]
	HaveFnv64aSumOf(sumValue uint64) fluent.IAdditional[string, IComparable]
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
