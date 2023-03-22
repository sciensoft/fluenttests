package strings

func (s *Comparable) Be(comparable string) IAdditional {
	be(s.testingT, s.value, comparable)

	additional := &Additional{}
	additional.Comparable = s

	return additional
}

func (s *Comparable) NotBe(comparable string) IAdditional {
	notBe(s.testingT, s.value, comparable)

	additional := &Additional{}
	additional.Comparable = s

	return additional
}

func (s *Comparable) BeLowerThan(comparable string) IAdditional {
	lowerThan(s.testingT, s.value, comparable)

	additional := &Additional{}
	additional.Comparable = s

	return additional
}

func (s *Comparable) BeLowerThanOrEqualTo(comparable string) IAdditional {
	lowerThanOrEqualTo(s.testingT, s.value, comparable)

	additional := &Additional{}
	additional.Comparable = s

	return additional
}

func (s *Comparable) BeGreaterThan(comparable string) IAdditional {
	greaterThan(s.testingT, s.value, comparable)

	additional := &Additional{}
	additional.Comparable = s

	return additional
}

func (s *Comparable) BeGreaterThanOrEqualTo(comparable string) IAdditional {
	greaterThanOrEqualTo(s.testingT, s.value, comparable)

	additional := &Additional{}
	additional.Comparable = s

	return additional
}
