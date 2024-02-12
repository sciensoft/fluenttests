package floats

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
	"golang.org/x/exp/constraints"
)

const comparableNotOfTypeMessageFormat string = "Comparison value is not of %t"

func be[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	if subject != comparable {
		message := fluent.GetMessage("Expected value [%f] is not equal to [%f].", messagesf...)
		t.Errorf(message, subject, comparable)
	}
}

func notBe[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	if subject == comparable {
		message := fluent.GetMessage("Expected value [%f] is equal to [%f].", messagesf...)
		t.Errorf(message, subject, comparable)
	}
}

func lowerThan[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	isLower := subject < comparable
	if !isLower {
		t.Errorf("Expected value [%f] is not lower than [%f].", subject, comparable)
	}
}

func lowerThanOrEqualTo[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	isLowerOrEqual := subject <= comparable
	if !isLowerOrEqual {
		t.Errorf("Expected value [%f] is not lower than or equal to [%f].", subject, comparable)
	}
}

func greaterThan[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	isGreater := subject > comparable
	if !isGreater {
		t.Errorf("Expected value [%f] is not greater than [%f].", subject, comparable)
	}
}

func greaterThanOrEqualTo[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	isGreaterOrEqual := subject >= comparable
	if !isGreaterOrEqual {
		t.Errorf("Expected value [%f] is not greater than or equal to [%f].", subject, comparable)
	}
}
