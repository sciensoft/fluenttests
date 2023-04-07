package strings

import (
	"hash/fnv"

	"sciensoft.dev/testing/fluent"
)

func (s *Comparable) Be(comparable string) fluent.IAdditional[string, IComparable] {
	match(s.testingT, false, s.value, comparable, "Expected value '%v' is not equal to '%v'.")
	return s.createAdditional()
}

func (s *Comparable) BeEmpty() {
	match(s.testingT, false, s.value, `^\s{0}$`, "Value '%v' is not empty. It's expected to match expression `%v`.")
}

func (s *Comparable) BeWhiteSpace() {
	match(s.testingT, false, s.value, `^\s{1,}$`, "Value '%v' does not match pattern `s%v`.")
}

func (s *Comparable) BeOneOf(comparables []string) fluent.IAdditional[string, IComparable] {
	beOneOf(s.testingT, false, s.value, comparables, "Expected value [%v] is not one of [%v].")
	return s.createAdditional()
}

func (s *Comparable) NotBe(comparable string) fluent.IAdditional[string, IComparable] {
	match(s.testingT, true, s.value, comparable, "Expected value '%v' is equal to '%v'.")
	return s.createAdditional()
}

func (s *Comparable) NotBeEmpty() fluent.IAdditional[string, IComparable] {
	match(s.testingT, true, s.value, `^\s{0}$`, "Value '%v' is empty. It's expected to match pattern `%v`.")
	return s.createAdditional()
}

func (s *Comparable) NotBeOneOf(comparables []string) fluent.IAdditional[string, IComparable] {
	beOneOf(s.testingT, true, s.value, comparables, "Expected value [%v] is one of [%v].")
	return s.createAdditional()
}

func (s *Comparable) Match(pattern string) fluent.IAdditional[string, IComparable] {
	match(s.testingT, false, s.value, pattern, "Value '%v' does not match pattern `%v`.")
	return s.createAdditional()
}

func (s *Comparable) NotMatch(pattern string) fluent.IAdditional[string, IComparable] {
	match(s.testingT, true, s.value, pattern, "Value '%v' does match pattern `%v`.")
	return s.createAdditional()
}

func (s *Comparable) StartWith(prefix string) fluent.IAdditional[string, IComparable] {
	startWith(s.testingT, false, s.value, prefix, "Value [%v] does not start with [%v]")
	return s.createAdditional()
}

func (s *Comparable) NotStartWith(prefix string) fluent.IAdditional[string, IComparable] {
	startWith(s.testingT, true, s.value, prefix, "Value [%v] does start with [%v]")
	return s.createAdditional()
}

func (s *Comparable) EndWith(prefix string) fluent.IAdditional[string, IComparable] {
	endWith(s.testingT, false, s.value, prefix, "Value [%v] does not end with [%v]")
	return s.createAdditional()
}

func (s *Comparable) NotEndWith(prefix string) fluent.IAdditional[string, IComparable] {
	endWith(s.testingT, true, s.value, prefix, "Value [%v] does end with [%v]")
	return s.createAdditional()
}

func (s *Comparable) HaveLengthOf(length int) fluent.IAdditional[string, IComparable] {
	haveLengthOf(s.testingT, s.value, length)
	return s.createAdditional()
}

func (s *Comparable) HaveFnv32aSumOf(comparable uint32) fluent.IAdditional[string, IComparable] {
	fnvSumFx := func() uint32 {
		h := fnv.New32a()
		h.Write([]byte(s.value))
		return h.Sum32()
	}

	haveFnv(s.testingT, fnvSumFx, s.value, comparable)

	return s.createAdditional()
}

func (s *Comparable) HaveFnv64aSumOf(comparable uint64) fluent.IAdditional[string, IComparable] {
	fnvSumFx := func() uint64 {
		h := fnv.New64a()
		h.Write([]byte(s.value))
		return h.Sum64()
	}

	haveFnv(s.testingT, fnvSumFx, s.value, comparable)

	return s.createAdditional()
}

func (s *Comparable) createAdditional() fluent.IAdditional[string, IComparable] {
	additional := &Additional{}
	additional.Comparable = s

	return additional
}
