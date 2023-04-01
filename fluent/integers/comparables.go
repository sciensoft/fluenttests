package integers

import "sciensoft.dev/testing/fluent"

func (s *Comparable[T]) Be(comparable T) {
	be(s.testingT, s.value, comparable)
}

func (s *Comparable[T]) BeZero() {
	be(s.testingT, s.value, 0, "Expected value [%d] to be zero (%d).")
}

func (s *Comparable[T]) BeNegative() fluent.IAdditional[T, IComparable[T]] {
	lowerThan(s.testingT, s.value, 0, "Expected value [%d] to be negative.")
	return s.createAdditional()
}

func (s *Comparable[T]) BePositive() fluent.IAdditional[T, IComparable[T]] {
	greaterThanOrEqualTo(s.testingT, s.value, 0, "Expected value [%d] to be positive.")
	return s.createAdditional()
}

func (s *Comparable[T]) NotBe(comparable T) fluent.IAdditional[T, IComparable[T]] {
	notBe(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) NotBeZero() fluent.IAdditional[T, IComparable[T]] {
	notBe(s.testingT, s.value, 0, "Expected value [%d] to be greater than zero (%d).")
	return s.createAdditional()
}

func (s *Comparable[T]) NotBeNegative() fluent.IAdditional[T, IComparable[T]] {
	greaterThanOrEqualTo(s.testingT, s.value, 0, "Expected value [%d] not to be negative.")
	return s.createAdditional()
}

func (s *Comparable[T]) NotBePositive() fluent.IAdditional[T, IComparable[T]] {
	lowerThan(s.testingT, s.value, 0, "Expected value [%d] not to be positive.")
	return s.createAdditional()
}

func (s *Comparable[T]) BeLowerThan(comparable T) fluent.IAdditional[T, IComparable[T]] {
	lowerThan(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) BeLowerThanOrEqualTo(comparable T) fluent.IAdditional[T, IComparable[T]] {
	lowerThanOrEqualTo(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) BeGreaterThan(comparable T) fluent.IAdditional[T, IComparable[T]] {
	greaterThan(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) BeGreaterThanOrEqualTo(comparable T) fluent.IAdditional[T, IComparable[T]] {
	greaterThanOrEqualTo(s.testingT, s.value, comparable)
	return s.createAdditional()
}

func (s *Comparable[T]) createAdditional() fluent.IAdditional[T, IComparable[T]] {
	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}
