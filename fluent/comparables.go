package fluent

func (s *Comparable[T]) Be(comparable T) IAdditional[T] {
	testable := (interface{})(s.value)

	switch value := testable.(type) {
	case int:
		be(s.testingT, value, comparable)
	case int8:
		be(s.testingT, value, comparable)
	case int16:
		be(s.testingT, value, comparable)
	case int32:
		be(s.testingT, value, comparable)
	case int64:
		be(s.testingT, value, comparable)
	case uint:
		be(s.testingT, value, comparable)
	case uint8:
		be(s.testingT, value, comparable)
	case uint16:
		be(s.testingT, value, comparable)
	case uint32:
		be(s.testingT, value, comparable)
	case uint64:
		be(s.testingT, value, comparable)
	case uintptr:
		be(s.testingT, value, comparable)
	case float32:
		be(s.testingT, value, comparable)
	case float64:
		be(s.testingT, value, comparable)
	case complex64:
		be(s.testingT, value, comparable)
	case complex128:
		be(s.testingT, value, comparable)
	case string:
		// TODO : Add string implementation
	default:
		// TODO : Skipping unmapped type
		s.testingT.Skip()
	}

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}

func (s *Comparable[T]) NotBe(comparable T) IAdditional[T] {
	testable := (interface{})(s.value)

	switch value := testable.(type) {
	case int:
		notBe(s.testingT, value, comparable)
	case int8:
		notBe(s.testingT, value, comparable)
	case int16:
		notBe(s.testingT, value, comparable)
	case int32:
		notBe(s.testingT, value, comparable)
	case int64:
		notBe(s.testingT, value, comparable)
	case uint:
		notBe(s.testingT, value, comparable)
	case uint8:
		notBe(s.testingT, value, comparable)
	case uint16:
		notBe(s.testingT, value, comparable)
	case uint32:
		notBe(s.testingT, value, comparable)
	case uint64:
		notBe(s.testingT, value, comparable)
	case uintptr:
		notBe(s.testingT, value, comparable)
	case float32:
		notBe(s.testingT, value, comparable)
	case float64:
		notBe(s.testingT, value, comparable)
	case complex64:
		notBe(s.testingT, value, comparable)
	case complex128:
		notBe(s.testingT, value, comparable)
	case string:
		// TODO : Add string implementation
	default:
		// TODO : Skipping unmapped type
		s.testingT.Skip()
	}

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}

func (s *Comparable[T]) BeLowerThan(comparable T) IAdditional[T] {
	testable := (interface{})(s.value)

	switch value := testable.(type) {
	case int:
		lowerThan(s.testingT, value, comparable)
	case int8:
		lowerThan(s.testingT, value, comparable)
	case int16:
		lowerThan(s.testingT, value, comparable)
	case int32:
		lowerThan(s.testingT, value, comparable)
	case int64:
		lowerThan(s.testingT, value, comparable)
	case uint:
		lowerThan(s.testingT, value, comparable)
	case uint8:
		lowerThan(s.testingT, value, comparable)
	case uint16:
		lowerThan(s.testingT, value, comparable)
	case uint32:
		lowerThan(s.testingT, value, comparable)
	case uint64:
		lowerThan(s.testingT, value, comparable)
	case uintptr:
		lowerThan(s.testingT, value, comparable)
	case float32:
		lowerThan(s.testingT, value, comparable)
	case float64:
		lowerThan(s.testingT, value, comparable)
	case complex64:
		// TODO : Add complex64 implementation
	case complex128:
		// TODO : Add complex128 implementation
	case string:
		// TODO : Add string implementation
	default:
		// TODO : Skipping unmapped type
		s.testingT.Skip()
	}

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}

func (s *Comparable[T]) BeLowerThanOrEqualTo(comparable T) IAdditional[T] {
	testable := (interface{})(s.value)

	switch value := testable.(type) {
	case int:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case int8:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case int16:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case int32:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case int64:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case uint:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case uint8:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case uint16:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case uint32:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case uint64:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case uintptr:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case float32:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case float64:
		lowerThanOrEqualTo(s.testingT, value, comparable)
	case complex64:
		// TODO : Add complex64 implementation
	case complex128:
		// TODO : Add complex128 implementation
	case string:
		// TODO : Add string implementation
	default:
		// TODO : Skipping unmapped type
		s.testingT.Skip()
	}

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}

func (s *Comparable[T]) BeGreaterThan(comparable T) IAdditional[T] {
	testable := (interface{})(s.value)

	switch value := testable.(type) {
	case int:
		greaterThan(s.testingT, value, comparable)
	case int8:
		greaterThan(s.testingT, value, comparable)
	case int16:
		greaterThan(s.testingT, value, comparable)
	case int32:
		greaterThan(s.testingT, value, comparable)
	case int64:
		greaterThan(s.testingT, value, comparable)
	case uint:
		greaterThan(s.testingT, value, comparable)
	case uint8:
		greaterThan(s.testingT, value, comparable)
	case uint16:
		greaterThan(s.testingT, value, comparable)
	case uint32:
		greaterThan(s.testingT, value, comparable)
	case uint64:
		greaterThan(s.testingT, value, comparable)
	case uintptr:
		greaterThan(s.testingT, value, comparable)
	case float32:
		greaterThan(s.testingT, value, comparable)
	case float64:
		greaterThan(s.testingT, value, comparable)
	case complex64:
		// TODO : Add complex64 implementation
	case complex128:
		// TODO : Add complex128 implementation
	case string:
		// TODO : Add string implementation
		// comparableValue, ok := (interface{})(comparable).(string)
		// if !ok {
		// 	s.testingT.Errorf("Comparison value is not of %t", new(T))
		// }
		// if value < comparableValue {
		// 	s.testingT.Errorf("Expected value [%v] is not greater than [%v].", value, comparableValue)
		// }
	default:
		s.testingT.Skipf("Skipping unmapped type.")
	}

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}

func (s *Comparable[T]) BeGreaterThanOrEqualTo(comparable T) IAdditional[T] {
	testable := (interface{})(s.value)

	switch value := testable.(type) {
	case int:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case int8:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case int16:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case int32:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case int64:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case uint:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case uint8:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case uint16:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case uint32:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case uint64:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case uintptr:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case float32:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case float64:
		greaterThanOrEqualTo(s.testingT, value, comparable)
	case complex64:
		// TODO : Add complex64 implementation
	case complex128:
		// TODO : Add complex128 implementation
	case string:
		// TODO : Add string implementation
		// comparableValue, ok := (interface{})(comparable).(string)
		// if !ok {
		// 	s.testingT.Errorf("Comparison value is not of %t", new(T))
		// }
		// if value < comparableValue {
		// 	s.testingT.Errorf("Expected value [%v] is not greater than [%v].", value, comparableValue)
		// }
	default:
		s.testingT.Skipf("Skipping unmapped type.")
	}

	additional := &Additional[T]{}
	additional.Comparable = s

	return additional
}
