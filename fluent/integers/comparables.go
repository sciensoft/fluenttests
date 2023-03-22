package integers

func (s *Comparable[T]) Be(comparable T) IAdditional[T] {
	be(s.testingT, s.value, comparable)

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}

func (s *Comparable[T]) NotBe(comparable T) IAdditional[T] {
	notBe(s.testingT, s.value, comparable)

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}

func (s *Comparable[T]) BeLowerThan(comparable T) IAdditional[T] {
	lowerThan(s.testingT, s.value, comparable)

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}

func (s *Comparable[T]) BeLowerThanOrEqualTo(comparable T) IAdditional[T] {
	lowerThanOrEqualTo(s.testingT, s.value, comparable)

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}

func (s *Comparable[T]) BeGreaterThan(comparable T) IAdditional[T] {
	greaterThan(s.testingT, s.value, comparable)

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}

func (s *Comparable[T]) BeGreaterThanOrEqualTo(comparable T) IAdditional[T] {
	greaterThanOrEqualTo(s.testingT, s.value, comparable)

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}
