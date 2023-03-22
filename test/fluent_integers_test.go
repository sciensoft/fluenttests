package tdd

import (
	"math"
	"testing"

	"sciensoft.dev/testing/fluent/integers"
)

func TestFluentShouldHaveIntToBeGreaterThan(t *testing.T) {
	fluent := integers.Fluent[int](t)

	fluent.It(30).
		Should().
		BeGreaterThan(10).
		And().
		BeGreaterThan(25)
}

func TestFluentShouldHaveFloatToBeGreaterThan(t *testing.T) {
	fluent := integers.Fluent[int64](t)

	fluent.It(math.MaxInt64).
		Should().
		BeGreaterThan(math.MaxInt64 - 1000).
		And().
		BeGreaterThan(math.MaxInt64 - 1)
}

func TestFluentShouldHaveIntToBeLowerThan(t *testing.T) {
	fluent := integers.Fluent[int](t)

	fluent.It(100).
		Should().
		BeLowerThan(150).
		And().
		BeLowerThan(233)
}

func TestFluentShouldHaveIntToBeLowerThanOrEqualTo(t *testing.T) {
	fluent := integers.Fluent[int](t)

	fluent.It(10).Should().
		Be(10).And().
		BeLowerThanOrEqualTo(15)
}
