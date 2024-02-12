# Fluent Tests

A Golang test library to increase productivity by providing easy-to-use methods for you to write TDD and BDD-style tests using a fluent language, allowing you to easily organize your tests following the Arrange, Act, and Assert pattern. Plus, several assertions methods for types, contracts, strings, integers, floats, and more to speed up your SDLC.

## Get Started

```bash
go get github.com/sciensoft/fluenttests
```

### Usage

```go
func TestFluentContractsShouldHaveMember(t *testing.T) {
    // Arrange
    fluent := contracts.Fluent[any](t)
    contractType := reflect.TypeOf(struct { Name string }{})
    fieldType := reflect.TypeOf("")
    
    // Act
    author := repository.RetrieveAuthorById(1) // Returns: struct { Name string }{}

    // Assert
    fluent.It(author).
        Should().BeOfType(contractType).
        And().HaveField("Name").OfType(fieldType).WithValue("Benjamin Treynor")
}
```

For more samples, please visit [the tests] written to test this own library.

## Features

- Assertions to validate
  - Types,
  - Strings,
  - Integers,
  - and Floats
- Assertions to validate Structs, Interfaces/Objects contracts

### Roadmap

- Support to `complex64`, and `complex128` types
- Add support to `reflect.Kind` for the testing methods `.BeOfType()`, and `.OfType()`
- Add more auxiliary methods for structs, interfaces, strings, floats, and integers
- Add support asynchronous tests
- Add support for panic testing

## Contributions

Before start contributing, check our [CONTRIBUTING] guidelines out. Also, before making any significant change, look at the existing Issues and Pull Requests. One of them may be tackling the same thing.

## Issues

Please use the [Issues] tab to open an issue or suggest a new feature.

## License

Copyright 2023 Sciensoft

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at [http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0).

[CONTRIBUTING]: ./CONTRIBUTING.md
[Issues]: ./../../../issues
[the tests]: ./test
