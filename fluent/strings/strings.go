package strings

import (
	"regexp"
	"strings"
	"testing"

	"sciensoft.dev/testing/fluent"
)

const comparableNotOfTypeMessageFormat string = "Comparison value is not of [%t]."

func be(t *testing.T, subject string, comparable string) {
	if subject != comparable {
		t.Errorf("Expected value [%v] is not equal to [%v].", subject, comparable)
	}
}

func beEmpty(t *testing.T, toBe bool, subject string, messagesf ...string) {
	isEmpty := subject == "" && len(subject) == 0

	if !toBe {
		isEmpty = !isEmpty
	}

	if !isEmpty {
		message := fluent.GetMessage("Value [%d] mis-match [%d].", messagesf...)
		t.Errorf(message, subject)
	}
}

func beOneOf(t *testing.T, toBe bool, subject string, comparables []string, messagesf ...string) {
	isOneOf := false

	for _, c := range comparables {
		if subject == c {
			isOneOf = true
		}
	}

	if !toBe {
		isOneOf = !isOneOf
	}

	if !isOneOf {
		message := fluent.GetMessage("Value [%d] one of mis-match [%d].", messagesf...)
		t.Errorf(message, subject, comparables)
	}
}

func notBe(t *testing.T, subject string, comparable string) {
	if subject == comparable {
		t.Errorf("Expected value [%v] is not equal to [%v].", subject, comparable)
	}
}

func match(t *testing.T, subject string, pattern string, comparableTest func(bool) bool, messagesf ...string) {
	b, err := regexp.MatchString(pattern, subject)

	if comparableTest(b) || err != nil {
		message := fluent.GetMessage("Value [%d] mis-match [%d].", messagesf...)
		t.Errorf(message, subject, pattern)
	}
}

func startWith(t *testing.T, toBe bool, subject string, comparable string, messagesf ...string) {
	hasPrefix := strings.HasPrefix(subject, comparable)

	if !toBe {
		hasPrefix = !hasPrefix
	}

	if !hasPrefix {
		message := fluent.GetMessage("Value [%d] prefix mis-match [%d].", messagesf...)
		t.Errorf(message, subject, comparable)
	}
}

func endWith(t *testing.T, toBe bool, subject string, comparable string, messagesf ...string) {
	hasSuffix := strings.HasSuffix(subject, comparable)

	if !toBe {
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
