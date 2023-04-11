package strings

import (
	"regexp"
	"strings"
	"testing"

	"github.com/sciensoft/fluenttests/fluent"
	"golang.org/x/exp/constraints"
)

const comparableNotOfTypeMessageFormat string = "Comparison value is not of [%t]."

func beOneOf(t *testing.T, invert bool, subject string, comparables []string, messagesf ...string) {
	isOneOf := false

	for _, c := range comparables {
		if subject == c {
			isOneOf = true
		}
	}

	if invert {
		isOneOf = !isOneOf
	}

	if !isOneOf {
		message := fluent.GetMessage("Value [%d] one of mis-match [%d].", messagesf...)
		t.Errorf(message, subject, comparables)
	}
}

func match(t *testing.T, invert bool, subject string, pattern string, messagesf ...string) {
	isMatch, err := regexp.MatchString(pattern, subject)

	if invert {
		isMatch = !isMatch
	}

	if !isMatch || err != nil {
		message := fluent.GetMessage("Value [%d] mis-match [%d].", messagesf...)
		t.Errorf(message, subject, pattern)
	}
}

func startWith(t *testing.T, invert bool, subject string, comparable string, messagesf ...string) {
	hasPrefix := strings.HasPrefix(subject, comparable)

	if invert {
		hasPrefix = !hasPrefix
	}

	if !hasPrefix {
		message := fluent.GetMessage("Value [%d] prefix mis-match [%d].", messagesf...)
		t.Errorf(message, subject, comparable)
	}
}

func endWith(t *testing.T, invert bool, subject string, comparable string, messagesf ...string) {
	hasSuffix := strings.HasSuffix(subject, comparable)

	if invert {
		hasSuffix = !hasSuffix
	}

	if !hasSuffix {
		message := fluent.GetMessage("Value [%d] prefix mis-match [%d].", messagesf...)
		t.Errorf(message, subject, comparable)
	}
}

func haveLengthOf(t *testing.T, subject string, comparable int) {
	if len(subject) != comparable {
		t.Errorf("Expected value [%v] to have length of [%d] but it has [%d].", subject, comparable, len(subject))
	}
}

func haveFnv[T constraints.Integer](t *testing.T, fvnSumFx func() T, subject string, comparable T) {
	var sum T = fvnSumFx()

	if sum != comparable {
		t.Errorf("Expected value [%v] to have Fvn32a sum of [%d] but it's sum is [%d].", subject, comparable, sum)
	}
}
