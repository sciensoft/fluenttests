// Copyright 2023 The FluentTest Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Package contracts provides members to register and test Mock objects.
// It is intended to be used in concert with "*testing.T" functions,
// and the "go test" command, which automates execution of any function
// of the form:
//
// "func TestXxx(*testing.T)" functions.
package contracts

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
)

// Fluent[T any] returns an "*fluent.Fluent[T]" representing a Mock object
// used on further test cases. The generic type parameter represents any
// testable type subject.
func Fluent[T any](t *testing.T) fluent.IFluent[T, IComparable[T]] {
	return &FluentT[T]{
		testingT: t,
	}
}

// It receives a subject object of generic type "T" for further testing,
// and returns an "Subject[T]" object.
func (s *FluentT[T]) It(subject T) fluent.ISubject[T, IComparable[T]] {
	sub := Subject[T]{
		Testable: &Testable[T]{},
	}
	sub.testingT = s.testingT
	sub.value = subject

	return &sub
}

// Should starts the testing operations agains the Mock subject, returning
// an "*Comparable[T]" object which provides a set of easy-to-use methods
// to test the subject in an additive maner.
func (s *Subject[T]) Should() IComparable[T] {
	return &Comparable[T]{
		Testable: s.Testable,
	}
}

// And extends the testing operation againts the Mock subject, returning
// an "*Comparable[T]" object which provides a set of easy-to-use methods
// to test the subject in an additive maner.
//
// The "And()" method works as an additive connector between test methods
// from "IComparable[T any]", allowing execution of chained testing methods.
//
//	package tdd
//
//	import "testing"
//	import "github.com/sciensoft/fluenttests/fluent/contracts"
//
//	func TestFluentContractsShouldHaveAllOfMembers(t *testing.T) {
//	  // Arrange
//	  fluent := contracts.Fluent[any](t)
//
//	  // ... code hiden for brevity
//
//	  // Assert
//	  fluent.It(obj).
//		  Should().BeOfType(objType).
//		  And().HaveField("A").OfType(fieldType).WithValue(1)
//	}
func (s *Additional[T]) And() IComparable[T] {
	return s.Comparable
}
