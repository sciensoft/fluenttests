// Copyright 2023 The FluentTest Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Package strings provides members to register and test string Mock objects.
// It is intended to be used in concert with "*testing.T" functions,
// and the "go test" command, which automates execution of any function
// of the form:
//
// "func TestXxx(*testing.T)" functions.
package strings

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
)

// Fluent returns an "*fluent.Fluent" representing a Mock object
// used on further test cases.
func Fluent(t *testing.T) fluent.IFluent[string, IComparable] {
	return &FluentT{
		testingT: t,
	}
}

// It receives a subject object of type "string", for further testing,
// and returns an "Subject" object.
func (s *FluentT) It(subject string) fluent.ISubject[string, IComparable] {
	sub := Subject{
		Testable: &Testable{},
	}
	sub.testingT = s.testingT
	sub.value = subject

	return &sub
}

// Should starts the testing operations agains the Mock subject, returning
// an "*Comparable" object which provides a set of easy-to-use methods
// to test the subject in an additive maner.
func (s *Subject) Should() IComparable {
	return &Comparable{
		Testable: s.Testable,
	}
}

// And extends the testing operation againts the Mock subject, returning
// an "*Comparable" object which provides a set of easy-to-use methods
// to test the subject in an additive maner.
//
// The "And()" method works as an additive connector between test methods
// from "IComparable", allowing execution of chained testing methods.
func (s *Additional) And() IComparable {
	return s.Comparable
}
