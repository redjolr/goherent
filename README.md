# goherent

A Go package / command for coherent tests.

Goherent provides a cleaner way to write descriptive test cases and produces more readable test results.

In Go, tests are typically organized into standalone functions, as shown below, where each function prefixed with Test is a separate test case.

```go
package some_test

import (
	"testing"
)

func TestThatExpectedValueIsAsExpectedWhenYouDoSomething(t *testing.T) {
    setupVar := "someValue"
    output := service.DoSomething(setupVar)
    if output != "expectedValue" {
        t.Errorf("Expected value was incorrect. Received: %s, expected: %s.", output, "expectedValue")
    }
}

func TestThatExpectedValueIsAsOtherExpectedWhenYouDoSomethingElse(t *testing.T) {
    setupVar := "someOtherValue"
    output := service.DoSomethingElse(setupVar)
    if output != "otherExpectedValue" {
        t.Errorf("Expected value was incorrect. Received: %s, expected: %s.", output, "otherExpectedValue")
    }
}

```

Providing a test message for these test cases is not possible, making it difficult to create a clear description as the tests become more complex. As we will see, Goherent makes it easy to pass test messages to a test case, thus improving test readability.

To run the tests and print the results for all test cases, use the command `go test -v ./....`. This will output this:

```bash
$ go test -v ./...
=== RUN   TestThatExpectedValueIsAsExpectedWhenYouDoSomething
--- PASS: TestThatExpectedValueIsAsExpectedWhenYouDoSomething (0.00s)
=== RUN   TestThatExpectedValueIsAsOtherExpectedWhenYouDoSomethingElse
--- PASS: TestThatExpectedValueIsAsOtherExpectedWhenYouDoSomethingElse (0.00s)
PASS
ok      github.com/redjolr/go-iam/src/core/tests        0.167s
```

The test report displayed in your terminal is clunky and awkward.

## The package

The `goherent` package provides a simple `Test` function which accepts a test message and a function with a `func(t *testing.T)` signature, where you write your test (and also a third parameter, where you need to pass `t`).

The awkward tests from the example above will be rewritten to:

```go
package some_test

import (
	"testing"

    . "github.com/redjolr/goherent/pkg/goherent"
)


func TestSomething(t *testing.T) {
    Test("The expected value should be as expected, if you do something.", func(t *testing.T) {
        setupVar := "someValue"
        output := service.DoSomething(setupVar)
        if output != "expectedValue" {
            t.Errorf("Expected value was incorrect. Received: %s, expected: %s.", output, "expectedValue")
        }
    }, t)

    // Or even more explicit test messages
    Test(`
        Given that our setupVar is "someValue"
        When we call service.DoSomething()
        Then its output should be "expectedValue"
    `, func(t *testing.T) {
        setupVar := "someValue"
        output := service.DoSomething(setupVar)
        if output != "expectedValue" {
            t.Errorf("Expected value was incorrect. Received: %s, expected: %s.", output, "expectedValue")
        }
    }, t)
}
```

When we run the tests with the goherent command, we see this beatiful report:

```
✅ TestUuid/Doing something
(0.167s)

✅ TestUuid/
Given that we have a V4 uuidStr = "b78a53b5-87e2-417a-9a43-0bf7c58022a8"
When we call UuidFromString(uuidStr)
Then a Uuid object must be returned and no error
(0.167s)
```
