package integers

import (
	"testing"

	"golang.org/x/exp/constraints"
)

const comparableNotOfTypeMessageFormat string = "Comparison value is not of %t"

func be[T constraints.Integer](t *testing.T, subject T, comparable T) {
	if subject != comparable {
		t.Errorf("Expected value [%v] is not equal to [%v].", subject, comparable)
	}
}

func notBe[T constraints.Integer](t *testing.T, subject T, comparable T) {
	if subject == comparable {
		t.Errorf("Expected value [%v] is not equal to [%v].", subject, comparable)
	}
}

func lowerThan[T constraints.Integer](t *testing.T, subject T, comparable T) {
	isLower := subject < comparable
	if !isLower {
		t.Errorf("Expected value [%v] IS NOT LOWER than [%v].", subject, comparable)
	}
}

func lowerThanOrEqualTo[T constraints.Integer](t *testing.T, subject T, comparable T) {
	isLowerOrEqual := subject <= comparable
	if !isLowerOrEqual {
		t.Errorf("Expected value [%v] IS NOT LOWER than [%v].", subject, comparable)
	}
}

func greaterThan[T constraints.Integer](t *testing.T, subject T, comparable T) {
	isGreater := subject > comparable
	if !isGreater {
		t.Errorf("Expected value [%v] IS NOT GREATER than [%v].", subject, comparable)
	}
}

func greaterThanOrEqualTo[T constraints.Integer](t *testing.T, subject T, comparable T) {
	isGreaterOrEqual := subject >= comparable
	if !isGreaterOrEqual {
		t.Errorf("Expected value [%v] IS NOT GREATER than [%v].", subject, comparable)
	}
}
