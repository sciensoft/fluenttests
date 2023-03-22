package tdd

import (
	"testing"

	"sciensoft.dev/testing/fluent/strings"
)

func TestFluentStringsShouldBeGreaterThan(t *testing.T) {
	fluent := strings.Fluent(t)

	fluent.It("D").
		Should().
		BeGreaterThan("B").
		And().
		BeGreaterThan("C")
}

func TestFluentStringsShouldBeLowerThan(t *testing.T) {
	fluent := strings.Fluent(t)

	fluent.It("A").
		Should().
		BeLowerThan("B").
		And().
		BeLowerThan("C")
}

func TestFluentStringsShouldBeLowerThanOrEqualTo(t *testing.T) {
	fluent := strings.Fluent(t)

	fluent.It("D").
		Should().
		Be("D").
		And().
		BeLowerThanOrEqualTo("E")
}
