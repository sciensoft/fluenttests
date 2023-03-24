package tdd

import (
	"testing"

	"sciensoft.dev/testing/fluent/strings"
)

func TestFluentStringShouldBe(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	text1 := "Hello world!"
	text2 := "Hello world!"

	// Assert
	fluent.It(text1).
		Should().Be(text2)
}

func TestFluentStringShouldBeEmpty(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("").
		Should().BeEmpty()
}

func TestFluentStringShouldBeOneOf(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	texts := []string{
		"This is a message",
		"which contains",
		"Hello World!",
	}

	// Assert
	fluent.It("Hello World!").
		Should().BeOneOf(texts)
}

func TestFluentStringShouldNotBe(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello World!").
		Should().NotBe("This is not hello world!")
}

func TestFluentStringShouldNotBeEmpty(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello there!").
		Should().NotBeEmpty()
}

func TestFluentStringShouldNotBeOneOf(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	texts := []string{
		"This is a message",
		"which contains",
		"Hello World!",
	}

	// Assert
	fluent.It("Hello Word!").
		Should().NotBeOneOf(texts)
}

func TestFluentStringShouldMatch(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello World!").
		Should().
		Match(`Hello (\w\W)*`).
		And().
		Match(`[a-zA-Z]+\!{1}`).
		And().
		Match(`(Hello){1}`).
		And().
		Match(`(World){1}`)
}

func TestFluentStringShouldNotMatch(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("This is Go!").
		Should().
		NotMatch(`Hello (\w\W)*`).
		And().
		NotMatch(`(Hello){1}`).
		And().
		NotMatch(`(World){1}`)
}

func TestFluentStringShouldStartWith(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello world!").
		Should().StartWith("He")
}

func TestFluentStringShouldNotStartWith(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello world!").
		Should().NotStartWith("Hall")
}

func TestFluentStringShouldEndWith(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello world!").
		Should().EndWith("ld!")
}

func TestFluentStringShouldNotEndWith(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello world!").
		Should().NotEndWith("word!")
}

func TestFluentStringShouldHaveLengthOf(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	length := len("Hello World!")

	// Assert
	fluent.It("Hello World!").
		Should().HaveLengthOf(length)
}
