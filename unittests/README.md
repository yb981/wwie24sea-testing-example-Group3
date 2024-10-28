# Introduction to Unit Testing

## What is Unit Testing?

Unit testing involves testing individual components or functions of a software application in isolation to ensure they work as expected. These tests are typically automated and written by developers during the development process.

## General Use Cases for Unit Testing

- **Validation of Functionality**: Ensure that each function or method performs as intended.
- **Early Bug Detection**: Identify and fix bugs early in the development cycle.
- **Refactoring Confidence**: Make changes to the codebase with confidence that existing functionality remains intact.
- **Documentation**: Serve as documentation for how functions are supposed to behave.

## Differentiating Unit Testing from Integration Testing

- **Unit Testing**: Focuses on individual components in isolation.
- **Integration Testing**: Focuses on the interaction between multiple components to ensure they work together as expected.

## The Testing Pyramid

The testing pyramid is a concept that illustrates the different levels of testing and their relative quantities:

1. **Unit Tests**: Form the base of the pyramid. They are numerous and fast.
2. **Integration Tests**: Fewer than unit tests, they test the interaction between components.
3. **End-to-End Tests**: Fewest in number, they test the entire application flow.

## Do's and Don'ts for Unit Testing

### Do's
- **Write Tests for All Functions**: Ensure every function has corresponding unit tests.
- **Keep Tests Independent**: Each test should run independently of others.
- **Use Mocking**: Mock external dependencies to isolate the unit being tested.
- **Run Tests Frequently**: Integrate tests into the continuous integration pipeline.

### Don'ts
- **Avoid Testing Multiple Units**: Each test should focus on a single unit.
- **Don't Depend on External Systems**: Avoid dependencies on databases, APIs, or other external systems.
- **Avoid Complex Logic in Tests**: Tests should be simple and straightforward.

## Examples

### Python Example

```python
import unittest

def add(a, b):
    return a + b

class TestMathFunctions(unittest.TestCase):
    def test_add(self):
        self.assertEqual(add(1, 2), 3)
        self.assertEqual(add(-1, 1), 0)
        self.assertEqual(add(-1, -1), -2)

if __name__ == '__main__':
    unittest.main()
```

### Go Example

```go
package main

import "testing"

func Add(a int, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    if Add(1, 2) != 3 {
        t.Error("Expected 1 + 2 to equal 3")
    }
    if Add(-1, 1) != 0 {
        t.Error("Expected -1 + 1 to equal 0")
    }
    if Add(-1, -1) != -2 {
        t.Error("Expected -1 + -1 to equal -2")
    }
}
```

By following these guidelines and examples, you can effectively implement unit testing in your projects and ensure a robust and reliable codebase.