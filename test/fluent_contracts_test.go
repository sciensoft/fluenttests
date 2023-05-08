package tdd

import (
	"reflect"
	"testing"

	"github.com/sciensoft/fluenttests/fluent/contracts"
)

type MyObject struct {
	Id int `json:"id"`
}

func TestFluentContractsShouldBeOfType(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj := struct {
		A int
	}{
		A: 1,
	}
	rtype := reflect.TypeOf(obj)

	// Assert
	fluent.It(obj).
		Should().BeOfType(rtype)
}

func TestFluentContractsShouldNotBeOfType(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj1 := struct{ A int }{A: 1}
	obj2 := struct{ B int }{B: 1}
	rtype := reflect.TypeOf(obj2)

	// Assert
	fluent.It(obj1).
		Should().NotBeOfType(rtype)
}

func TestFluentContractsShouldHaveMember(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj := struct{ Name string }{Name: "Robert Griesemer"}

	// Assert
	fluent.It(obj).
		Should().HaveMember("Name")
}

func TestFluentContractsShouldHaveField(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj := struct{ Value float32 }{Value: 105.4}

	// Assert
	fluent.It(obj).
		Should().HaveField("Value")
}

func TestFluentContractsShouldHaveFieldWithTag(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj := struct {
		Value float32 `json:"value"`
	}{
		Value: 105.4,
	}

	// Assert
	fluent.It(obj).
		Should().HaveFieldWithTag("Value", "json")
}

func TestFluentContractsShouldHaveAllFieldsWithTag(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj := struct {
		Title string  `json:"title"`
		Value float32 `json:"value"`
	}{
		Title: "Book Xyz",
		Value: 105.4,
	}

	// Assert
	fluent.It(obj).
		Should().HaveAllFieldsWithTag("json")
}

func TestFluentContractsShouldHaveMethod(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	// Assert
	fluent.It(fluent).
		Should().HaveMethod("It")
}

func TestFluentContractsShouldHaveAnyOfMembers(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj := struct {
		Title     string
		Author    string
		Publisher string
	}{
		Title:     "A Concurrent Window System",
		Author:    "Rob Pike",
		Publisher: "AT&T Bell Laboratories",
	}

	// Assert
	fluent.It(obj).
		Should().
		HaveAnyOfMembers([]string{
			"Title",
			"Description",
			"PublicationDate",
		})
}

func TestFluentContractsShouldHaveAllOfMembers(t *testing.T) {
	// Arrange
	fluent := contracts.Fluent[any](t)

	// Act
	obj := struct {
		Title     string
		Author    string
		Publisher string
	}{
		Title:     "A Concurrent Window System",
		Author:    "Rob Pike",
		Publisher: "AT&T Bell Laboratories",
	}

	// Assert
	fluent.It(obj).
		Should().
		HaveAllOfMembers([]string{
			"Title",
			"Author",
			"Publisher",
		})
}
