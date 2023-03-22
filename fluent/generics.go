package fluent

import (
	"testing"

	"golang.org/x/exp/constraints"
)

const comparableNotOfTypeMessageFormat string = "Comparison value is not of %t"

func be[T constraints.Integer | constraints.Float | constraints.Complex](t *testing.T, subject T, comparable interface{}) {
	comparableValue, ok := comparable.(T)

	if !ok {
		t.Errorf(comparableNotOfTypeMessageFormat, comparable)
	}

	if subject != comparableValue {
		t.Errorf("Expected value [%v] is not equal to [%v].", subject, comparable)
	}
}

func notBe[T constraints.Integer | constraints.Float | constraints.Complex](t *testing.T, subject T, comparable interface{}) {
	comparableValue, ok := comparable.(T)

	if !ok {
		t.Errorf(comparableNotOfTypeMessageFormat, comparable)
	}

	if subject == comparableValue {
		t.Errorf("Expected value [%v] is not equal to [%v].", subject, comparable)
	}
}

func lowerThan[T constraints.Integer | constraints.Float](t *testing.T, subject T, comparable interface{}) {
	comparableValue, ok := comparable.(T)

	if !ok {
		t.Errorf(comparableNotOfTypeMessageFormat, comparable)
	}

	isLower := subject < comparableValue
	if !isLower {
		t.Errorf("Expected value [%v] IS NOT LOWER than [%v].", subject, comparable)
	}
}

func lowerThanOrEqualTo[T constraints.Integer | constraints.Float](t *testing.T, subject T, comparable interface{}) {
	comparableValue, ok := comparable.(T)

	if !ok {
		t.Errorf(comparableNotOfTypeMessageFormat, comparable)
	}

	isLowerOrEqual := subject <= comparableValue
	if !isLowerOrEqual {
		t.Errorf("Expected value [%v] IS NOT LOWER than [%v].", subject, comparable)
	}
}

func greaterThan[T constraints.Integer | constraints.Float](t *testing.T, subject T, comparable interface{}) {
	comparableValue, ok := comparable.(T)

	if !ok {
		t.Errorf(comparableNotOfTypeMessageFormat, comparable)
	}

	isGreater := subject > comparableValue
	if !isGreater {
		t.Errorf("Expected value [%v] IS NOT GREATER than [%v].", subject, comparable)
	}
}

func greaterThanOrEqualTo[T constraints.Integer | constraints.Float](t *testing.T, subject T, comparable interface{}) {
	comparableValue, ok := comparable.(T)

	if !ok {
		t.Errorf(comparableNotOfTypeMessageFormat, comparable)
	}

	isGreaterOrEqual := subject >= comparableValue
	if !isGreaterOrEqual {
		t.Errorf("Expected value [%v] IS NOT GREATER than [%v].", subject, comparable)
	}
}
