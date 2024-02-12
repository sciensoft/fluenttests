package strings

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
)

type String interface {
	~string
}

// IComparable interface for the strings package. Following
// the Fluent-Interface design pattern to provide chained testing methods.
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

// FluentT is a type returned by calling "strings.Fluent(t)"
// for enabling easy-to-use testing methods.
//
// It carries the "*testing.T" object to be further used in the methods chain.
type FluentT struct {
	testingT *testing.T
}

// Subject is a type returned by calling "fluentObj.It("My text string.")"
// used to encapsulate the "*testing.T" object and testable subject.
//
// NOTE: The "It()" method in the example accepts string objects.
type Subject struct {
	*Testable
}

// Additional is a type returned by some easy-to-use testing methods,
// providing basic additive test chaining with "And()".
type Additional struct {
	*Comparable
}

// Comparable is a type returned by calling "fluentObj.It("My text string.").Should()",
// encapsulating a "*Testable[T]" object for further testing, and providing
// all the strings package's easy-to-use set of methods with signatures declared at
// the "strings.IComparable" interface.
//
// NOTE: The "It()" method in the example accepts string objects.
type Comparable struct {
	*Testable
}

// Testable is a type that holds the "*testing.T" object and the underline
// Mock object, supporting further testing using the chain of easy-to-use methods.
// It's used by the "Subject" and "Comparable" types.
type Testable struct {
	testingT *testing.T
	value    string
}
