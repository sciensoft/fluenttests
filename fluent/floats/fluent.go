// Copyright 2023 The FluentTest Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Package floats provides members to register and test float Mock objects.
// It is intended to be used in concert with "*testing.T" functions,
// and the "go test" command, which automates execution of any function
// of the form:
//
// "func TestXxx(*testing.T)" functions.
package floats

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
	"golang.org/x/exp/constraints"
)

// Fluent[T constraints.Float] returns an "*fluent.Fluent[T]" representing a Mock object
// used on further test cases. The generic type parameter represents any
// testable float32 or float64 subject.
func Fluent[T constraints.Float](t *testing.T) fluent.IFluent[T, IComparable[T]] {
	return &FluentT[T]{
		testingT: t,
	}
}

// It receives a subject object of generic type "T", a constraint of Float
// for further testing, and returns an "Subject[T]" object.
func (s *FluentT[T]) It(subject T) fluent.ISubject[T, IComparable[T]] {
	sub := Subject[T]{
		Testable: &Testable[T]{},
	}
	sub.testingT = s.testingT
	sub.value = subject

	return &sub
}

// Should starts the testing operations agains the Mock subject, returning
// an "*Comparable[T]" object which provides a set of helpful methods
// to test the subject in an additive maner.
func (s *Subject[T]) Should() IComparable[T] {
	return &Comparable[T]{
		Testable: s.Testable,
	}
}

// And extends the testing operation againts the Mock subject, returning
// an "*Comparable[T]" object which provides a set of helpful methods
// to test the subject in an additive maner.
//
// The "And()" method works as an additive connector between test methods
// from "IComparable[T]", allowing execution of chained testing methods.
func (s *Additional[T]) And() IComparable[T] {
	return s.Comparable
}
