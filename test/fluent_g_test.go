package tdd

import (
	"testing"

	test "sciensoft.dev/testing/fluent"
)

func TestNumber30ShouldBeGreaterThan10And25(t *testing.T) {
	fluent := test.Fluent[int](t)

	fluent.It(30).
		Should().
		BeGreaterThan(10).
		And().
		BeGreaterThan(25)
}

func TestNumber100ShouldBeLowerThan150And233(t *testing.T) {
	fluent := test.Fluent[int](t)

	fluent.It(100).
		Should().
		BeLowerThan(150).
		And().
		BeLowerThan(233)
}

func TestNumber10ShouldBeLowerThanOrEqualTo15(t *testing.T) {
	fluent := test.Fluent[int](t)

	fluent.It(10).Should().
		Be(10).And().
		BeLowerThanOrEqualTo(15)
}
