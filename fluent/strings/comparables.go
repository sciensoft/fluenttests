package strings

func (s *Comparable) Be(comparable string) IAdditional {
	match(s.testingT, true, s.value, comparable, "Expected value '%v' is not equal to '%v'.")
	return s.createAdditional()
}

func (s *Comparable) BeEmpty() {
	match(s.testingT, true, s.value, `^\s{0}$`, "Value '%v' is not empty. It's expected to match expression `%v`.")
}

func (s *Comparable) BeWhiteSpace() {
	match(s.testingT, true, s.value, `^\s{1,}$`, "Value '%v' does not match pattern `s%v`.")
}

func (s *Comparable) BeOneOf(comparables []string) IAdditional {
	beOneOf(s.testingT, true, s.value, comparables, "Expected value [%v] is not one of [%v].")
	return s.createAdditional()
}

func (s *Comparable) NotBe(comparable string) IAdditional {
	match(s.testingT, false, s.value, comparable, "Expected value '%v' is equal to '%v'.")
	return s.createAdditional()
}

func (s *Comparable) NotBeEmpty() IAdditional {
	match(s.testingT, false, s.value, `^\s{0}$`, "Value '%v' is empty. It's expected to match pattern `%v`.")
	return s.createAdditional()
}

func (s *Comparable) NotBeOneOf(comparables []string) IAdditional {
	beOneOf(s.testingT, false, s.value, comparables, "Expected value [%v] is one of [%v].")
	return s.createAdditional()
}

func (s *Comparable) Match(pattern string) IAdditional {
	match(s.testingT, true, s.value, pattern, "Value '%v' does not match pattern `%v`.")
	return s.createAdditional()
}

func (s *Comparable) NotMatch(pattern string) IAdditional {
	match(s.testingT, false, s.value, pattern, "Value '%v' does match pattern `%v`.")
	return s.createAdditional()
}

func (s *Comparable) StartWith(prefix string) IAdditional {
	startWith(s.testingT, true, s.value, prefix, "Value [%v] does not start with [%v]")
	return s.createAdditional()
}

func (s *Comparable) NotStartWith(prefix string) IAdditional {
	startWith(s.testingT, false, s.value, prefix, "Value [%v] does start with [%v]")
	return s.createAdditional()
}

func (s *Comparable) EndWith(prefix string) IAdditional {
	endWith(s.testingT, true, s.value, prefix, "Value [%v] does not end with [%v]")
	return s.createAdditional()
}

func (s *Comparable) NotEndWith(prefix string) IAdditional {
	endWith(s.testingT, false, s.value, prefix, "Value [%v] does end with [%v]")
	return s.createAdditional()
}

func (s *Comparable) HaveLengthOf(length int) IAdditional {
	haveLengthOf(s.testingT, s.value, length)
	return s.createAdditional()
}

func (s *Comparable) createAdditional() IAdditional {
	additional := &Additional{}
	additional.Comparable = s

	return additional
}
