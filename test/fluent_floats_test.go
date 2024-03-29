package tdd

import (
	"math"
	"math/rand"
	"testing"

	"github.com/sciensoft/fluenttests/fluent/floats"
)

func TestFluentFloat32ShouldBe(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)

	// Act
	// ... Noop

	// Assert
	fluent.It(10.11).
		Should().
		Be(10.11)
}

func TestFluentFloat64ShouldBe(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)

	// Act
	// ... Noop

	// Assert
	fluent.It(0.6645600532184904).
		Should().
		Be(0.6645600532184904)
}

func TestFluentFloat32ShouldBeZero(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)

	// Act
	// ... Noop

	// Assert
	fluent.It(0).
		Should().
		BeZero()
}

func TestFluentFloat64ShouldBeZero(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)

	// Act
	// ... Noop

	// Assert
	fluent.It(0.0).
		Should().
		BeZero()
}

func TestFluentFloat32ShouldBeNegative(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)

	// Act
	// ... Noop

	// Assert
	fluent.It(-0.0114131).
		Should().
		BeNegative()
}

func TestFluentFloat64ShouldBeNegative(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)

	// Act
	// ... Noop

	// Assert
	fluent.It(-0.6645600532184904).
		Should().
		BeNegative()
}

func TestFluentFloat32ShouldBePositive(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)
	number := rand.Float32()

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		BePositive()
}

func TestFluentFloat64ShouldBePositive(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)
	number := rand.Float64()

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		BePositive()
}

func TestFluentFloat32ShouldToNotBe(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)
	number1 := rand.Float32()
	number2 := rand.Float32()

	// Act
	// ... Noop

	// Assert
	fluent.It(number1).
		Should().
		NotBe(number2)
}

func TestFluentFloat64ShouldToNotBe(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)
	number1 := rand.Float64()
	number2 := rand.Float64()

	// Act
	// ... Noop

	// Assert
	fluent.It(number1).
		Should().
		NotBe(number2)
}

func TestFluentFloat32ShouldToNotBeZero(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)
	number := (rand.Float32() + 10) * 10

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		NotBeZero()
}

func TestFluentFloat64ShouldToNotBeZero(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)
	var base, expo float64 = 10, 5

	// Act
	number := math.Pow(rand.Float64()+base, expo)

	// Assert
	fluent.It(number).
		Should().
		NotBeZero()
}

func TestFluentShouldFloatToNotBeNegative(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)

	// Act
	// ... Noop

	// Assert
	fluent.It(rand.Float32()).
		Should().
		NotBeNegative()
}

func TestFluentFloat32ShouldToNotBePositive(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)
	number := -(rand.Float32())

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		NotBePositive()
}

func TestFluentFloat64ShouldToNotBePositive(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)
	number := -(rand.Float64())

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		NotBePositive()
}

func TestFluentFloat32ShouldBeLowerThan(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)
	number := rand.Float32()

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		BeLowerThan(number * 100.55)
}

func TestFluentFloat64ShouldBeLowerThan(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)
	number := rand.Float64()

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		BeLowerThan(number * 1000.523232235)
}

func TestFluentFloat32ShouldBeLowerThanOrEqualTo(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)
	number := rand.Float32()

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		BeLowerThanOrEqualTo(number * 1000).
		And().
		BeGreaterThanOrEqualTo(number / 100)
}

func TestFluentFloat64ShouldBeLowerThanOrEqualTo(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)
	number := rand.Float64()

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		BeLowerThanOrEqualTo(number * 1000).
		And().
		BeGreaterThanOrEqualTo(number / 100)
}

func TestFluentFloat32ShouldBeGreaterThan(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)
	number := rand.Float32()

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		BeGreaterThan(number / 10)
}

func TestFluentFloat64ShouldBeGreaterThan(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)
	number := rand.Float64()

	// Act
	// ... Noop

	// Assert
	fluent.It(number).
		Should().
		BeGreaterThan(number / 100).
		And().
		BeLowerThan(number * 1000)
}

func TestFluentFloat32ShouldBeGreaterThanOrEqualTo(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float32](t)

	// Act
	// ... Noop

	// Assert
	fluent.It(math.MaxFloat32).
		Should().
		BeGreaterThanOrEqualTo(math.MaxFloat32 - 1000)
}

func TestFluentFloat64ShouldBeGreaterThanOrEqualTo(t *testing.T) {
	// Arrange
	fluent := floats.Fluent[float64](t)

	// Act
	// ... Noop

	// Assert
	fluent.It(math.MaxFloat64).
		Should().
		BeGreaterThanOrEqualTo(math.MaxFloat64 - 1000)
}
