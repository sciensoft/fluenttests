package floats

import (
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
	"golang.org/x/exp/constraints"
)

const comparableNotOfTypeMessageFormat string = "Comparison value is not of %t"

func be[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	if subject != comparable {
		message := fluent.GetMessage("Expected value [%d] is not equal to [%d].", messagesf...)
		t.Errorf(message, subject, comparable)
	}
}

func notBe[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	if subject == comparable {
		message := fluent.GetMessage("Expected value [%d] is equal to [%d].", messagesf...)
		t.Errorf(message, subject, comparable)
	}
}

func lowerThan[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	isLower := subject < comparable
	if !isLower {
		t.Errorf("Expected value [%d] is not lower than [%d].", subject, comparable)
	}
}

func lowerThanOrEqualTo[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	isLowerOrEqual := subject <= comparable
	if !isLowerOrEqual {
		t.Errorf("Expected value [%d] is not lower than or equal to [%d].", subject, comparable)
	}
}

func greaterThan[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	isGreater := subject > comparable
	if !isGreater {
		t.Errorf("Expected value [%d] is not greater than [%d].", subject, comparable)
	}
}

func greaterThanOrEqualTo[T constraints.Float](t *testing.T, subject T, comparable T, messagesf ...string) {
	isGreaterOrEqual := subject >= comparable
	if !isGreaterOrEqual {
		t.Errorf("Expected value [%d] is not greater than or equal to [%d].", subject, comparable)
	}
}
