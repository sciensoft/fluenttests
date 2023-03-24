package strings

func (s *Comparable) Be(comparable string) IAdditional {
	be(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable) BeEmpty() {
	beEmpty(s.testingT, true, s.value, "Value [%v] is not empty.")
}

func (s *Comparable) BeOneOf(comparables []string) IAdditional {
	beOneOf(s.testingT, true, s.value, comparables, "Expected value [%v] is not one of [%v].")
	return s.createAdditional()
}

func (s *Comparable) NotBe(comparable string) IAdditional {
	notBe(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable) NotBeEmpty() IAdditional {
	beEmpty(s.testingT, false, s.value, "Value [%v] is empty.")
	return s.createAdditional()
}

func (s *Comparable) NotBeOneOf(comparables []string) IAdditional {
	beOneOf(s.testingT, false, s.value, comparables, "Expected value [%v] is one of [%v].")
	return s.createAdditional()
}

func (s *Comparable) Match(pattern string) IAdditional {
	comparable := func(b bool) bool { return !b }
	match(s.testingT, s.value, pattern, comparable, "Value [%v] does not match pattern [%v].")
	return s.createAdditional()
}

func (s *Comparable) NotMatch(pattern string) IAdditional {
	comparable := func(b bool) bool { return b }
	match(s.testingT, s.value, pattern, comparable, "Value [%v] does match pattern [%v].")
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
