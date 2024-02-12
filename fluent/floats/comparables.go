package floats

import (
	"github.com/sciensoft/fluenttests/fluent"
)

// Be asserts a Mock object as having the constrained-typed value as per "value T" argument.
func (s *Comparable[T]) Be(value T) {
	be(s.testingT, s.value, value)
}

// BeZero asserts a Mock object as having value of "0.0".
func (s *Comparable[T]) BeZero() {
	be(s.testingT, s.value, 0.0, "Expected value [%d] to be zero (%d).")
}

// BeNegative asserts a Mock object as having a negative value "<0.0".
func (s *Comparable[T]) BeNegative() fluent.IAdditional[T, IComparable[T]] {
	lowerThan(s.testingT, s.value, 0.0, "Expected value [%d] to be negative.")
	return s.createAdditional()
}

// BePositive asserts a Mock object as having a positive value ">=0.0".
func (s *Comparable[T]) BePositive() fluent.IAdditional[T, IComparable[T]] {
	greaterThanOrEqualTo(s.testingT, s.value, 0.0, "Expected value [%d] to be positive.")
	return s.createAdditional()
}

// NotBe asserts a Mock object as NOT having the same value as per "value T" argument.
func (s *Comparable[T]) NotBe(value T) fluent.IAdditional[T, IComparable[T]] {
	notBe(s.testingT, s.value, value)
	return s.createAdditional()
}

// NotBeZero asserts a Mock object as NOT having value of "0.0".
func (s *Comparable[T]) NotBeZero() fluent.IAdditional[T, IComparable[T]] {
	notBe(s.testingT, s.value, 0.0, "Expected value [%d] to be greater than zero (%d).")
	return s.createAdditional()
}

// NotBeNegative asserts a Mock object as NOT having a negative value.
func (s *Comparable[T]) NotBeNegative() fluent.IAdditional[T, IComparable[T]] {
	greaterThanOrEqualTo(s.testingT, s.value, 0.0, "Expected value [%d] not to be negative.")
	return s.createAdditional()
}

// NotBePositive asserts a Mock object as NOT having a positive value.
func (s *Comparable[T]) NotBePositive() fluent.IAdditional[T, IComparable[T]] {
	lowerThan(s.testingT, s.value, 0.0, "Expected value [%d] not to be positive.")
	return s.createAdditional()
}

// BeLowerThan asserts a Mock object as having a value lower than the "value T" argument.
func (s *Comparable[T]) BeLowerThan(value T) fluent.IAdditional[T, IComparable[T]] {
	lowerThan(s.testingT, s.value, value)
	return s.createAdditional()
}

// BeLowerThanOrEqualTo asserts a Mock object as having a value lower than, or equal to the "value T" argument.
func (s *Comparable[T]) BeLowerThanOrEqualTo(value T) fluent.IAdditional[T, IComparable[T]] {
	lowerThanOrEqualTo(s.testingT, s.value, value)
	return s.createAdditional()
}

// BeGreaterThan asserts a Mock object as having a value greater than the "value T" argument.
func (s *Comparable[T]) BeGreaterThan(value T) fluent.IAdditional[T, IComparable[T]] {
	greaterThan(s.testingT, s.value, value)
	return s.createAdditional()
}

// BeGreaterThanOrEqualTo asserts a Mock object as having a value greater than, or equal to the "value T" argument.
func (s *Comparable[T]) BeGreaterThanOrEqualTo(value T) fluent.IAdditional[T, IComparable[T]] {
	greaterThanOrEqualTo(s.testingT, s.value, value)
	return s.createAdditional()
}

// createAdditional supports the creation of chained easy-to-use testing methods.
func (s *Comparable[T]) createAdditional() fluent.IAdditional[T, IComparable[T]] {
	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}
