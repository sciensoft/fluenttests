package tdd

import (
	"math"
	"testing"

	"github.com/sciensoft/fluenttests/fluent/integers"
)

func TestFluentIntShouldBe(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(10).
		Should().
		Be(10)
}

func TestFluentIntShouldBeZero(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(0).
		Should().
		BeZero()
}

func TestFluentIntShouldBeNegative(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(-1).
		Should().
		BeNegative()
}

func TestFluentIntShouldBePositive(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(10).
		Should().
		BePositive()
}

func TestFluentIntShouldNotBe(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(142).
		Should().
		NotBe(10)
}

func TestFluentIntShouldNotBeZero(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(1).
		Should().
		NotBeZero()
}

func TestFluentIntShouldNotBeNegative(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(10).
		Should().
		NotBeNegative()
}

func TestFluentIntShouldNotBePositive(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(-10).
		Should().
		NotBePositive()
}

func TestFluentIntShouldBeLowerThan(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(100).
		Should().
		BeLowerThan(150).
		And().
		BeLowerThan(233)
}

func TestFluentIntShouldBeLowerThanOrEqualTo(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(10).
		Should().
		BeLowerThanOrEqualTo(15)
}

func TestFluentIntShouldBeGreaterThan(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int](t)

	// Act
	// Assert
	fluent.It(30).
		Should().
		BeGreaterThan(25).
		And().
		BeLowerThan(50)
}

func TestFluentIntShouldBeGreaterThanOrEqualTo(t *testing.T) {
	// Arrange
	fluent := integers.Fluent[int64](t)

	// Act
	// Assert
	fluent.It(math.MaxInt64).
		Should().
		BeGreaterThanOrEqualTo(math.MaxInt64 - 1000)
}
