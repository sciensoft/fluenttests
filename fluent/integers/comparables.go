package integers

func (s *Comparable[T]) Be(comparable T) {
	be(s.testingT, s.value, comparable)
}

func (s *Comparable[T]) BeZero() {
	be(s.testingT, s.value, 0, "Expected value [%d] to be zero (%d).")
}

func (s *Comparable[T]) BeNegative() IAdditional[T] {
	lowerThan(s.testingT, s.value, 0, "Expected value [%d] to be negative.")
	return s.createAdditional()
}

func (s *Comparable[T]) BePositive() IAdditional[T] {
	greaterThanOrEqualTo(s.testingT, s.value, 0, "Expected value [%d] to be positive.")
	return s.createAdditional()
}

func (s *Comparable[T]) NotBe(comparable T) IAdditional[T] {
	notBe(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) NotBeZero() IAdditional[T] {
	notBe(s.testingT, s.value, 0, "Expected value [%d] to be greater than zero (%d).")
	return s.createAdditional()
}

func (s *Comparable[T]) NotBeNegative() IAdditional[T] {
	greaterThanOrEqualTo(s.testingT, s.value, 0, "Expected value [%d] not to be negative.")
	return s.createAdditional()
}

func (s *Comparable[T]) NotBePositive() IAdditional[T] {
	lowerThan(s.testingT, s.value, 0, "Expected value [%d] not to be positive.")
	return s.createAdditional()
}

func (s *Comparable[T]) BeLowerThan(comparable T) IAdditional[T] {
	lowerThan(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) BeLowerThanOrEqualTo(comparable T) IAdditional[T] {
	lowerThanOrEqualTo(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) BeGreaterThan(comparable T) IAdditional[T] {
	greaterThan(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) BeGreaterThanOrEqualTo(comparable T) IAdditional[T] {
	greaterThanOrEqualTo(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) createAdditional() IAdditional[T] {
	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}
