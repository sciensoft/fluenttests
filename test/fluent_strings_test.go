package tdd

import (
	"hash/fnv"
	"testing"

	"sciensoft.dev/testing/fluent/strings"
)

func TestFluentStringsShouldBe(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	text1 := "Ola, tudo bem?"
	text2 := "Ola, tudo bem?"

	// Assert
	fluent.It(text1).
		Should().Be(text2)
}

func TestFluentStringsShouldBeEmpty(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("").
		Should().BeEmpty()
}

func TestFluentStringsShouldBeOneOf(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	texts := []string{
		"This is a message",
		"which contains",
		"Bonjour le monde!",
	}

	// Assert
	fluent.It("Bonjour le monde!").
		Should().BeOneOf(texts)
}

func TestFluentStringsShouldNotBe(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("¡Hola Mundo!").
		Should().NotBe("¡Esto no es Hola Mundo!")
}

func TestFluentStringsShouldNotBeEmpty(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello there!").
		Should().NotBeEmpty()
}

func TestFluentStringsShouldNotBeOneOf(t *testing.T) {
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

func TestFluentStringsShouldMatch(t *testing.T) {
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

func TestFluentStringsShouldNotMatch(t *testing.T) {
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

func TestFluentStringsShouldStartWith(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello world!").
		Should().StartWith("He")
}

func TestFluentStringsShouldNotStartWith(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello world!").
		Should().NotStartWith("Hall")
}

func TestFluentStringsShouldEndWith(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello world!").
		Should().EndWith("ld!")
}

func TestFluentStringsShouldNotEndWith(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	// Assert
	fluent.It("Hello world!").
		Should().NotEndWith("word!")
}

func TestFluentStringsShouldHaveLengthOf(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)

	// Act
	length := len("Hello World!")

	// Assert
	fluent.It("Hello World!").
		Should().HaveLengthOf(length)
}

func TestFluentStringsShouldHaveFnv32aSumOf(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)
	fnvHasher := fnv.New32a()

	// Act
	fnvHasher.Write([]byte("Hello World!"))
	fnvSum32 := fnvHasher.Sum32()

	// Assert
	fluent.It("Hello World!").
		Should().HaveFnv32aSumOf(fnvSum32)
}

func TestFluentStringsShouldHaveFnv64aSumOf(t *testing.T) {
	// Arrange
	fluent := strings.Fluent(t)
	fnvHasher := fnv.New64a()

	// Act
	fnvHasher.Write([]byte("Hello World!"))
	fnvSum64 := fnvHasher.Sum64()

	// Assert
	fluent.It("Hello World!").
		Should().HaveFnv64aSumOf(fnvSum64)
}
