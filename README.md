# Fluent Tests

A library with many helpful methods for you to write TDD and BDD-style tests using a fluent language, allowing you to easily organize your tests following the AAA (Arrange, Act, and Assert).

## Get Started

```bash
go get github.com/AlexzSouz/sciensoft-fluenttest
```

### Usage

```golang
func TestFluentContractsShouldHaveMember(t *testing.T) {
    // Arrange
    fluent := contracts.Fluent[any](t)

    // Act
    model := struct{
        Name string 
    }{
        Name: "Robert Griesemer",
    }

    // Assert
    fluent.It(model).
        Should().HaveMember("Name")
}
```

For more samples, please visit [the tests] written to test this own library.

## Features

- Several methods to valid strings, floats, and integers
- Methods to validate structs and interfaces members

### Roadmap

- Support to `complex64`, and `complex128` types
- Add more auxiliary methods for strings, floats, and integers
- Add asynchronous test support
- Support for test panic

## Contributions

Before start contributing, check our [CONTRIBUTING] guidelines out. Also, before making any significant change, look at the existing Issues and Pull Requests. One of them may be tackling the same thing.

## Issues

Please use the [Issues] tab to open an issue or suggest a new feature.

## License

Copyright 2019 Sciensoft

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at [http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0).

[CONTRIBUTING]: ./CONTRIBUTING.md
[Issues]: ./../../../issues
[the tests]: ./test
