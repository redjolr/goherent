# goherent

**Coherent tests for Go** — a Jest-inspired assertion API and a clean test report, built as a thin wrapper around the standard `go test` runner.

```
✅ TestAdd/it adds two numbers (12ms)
✅ TestAdd/it is commutative (1.20s)

✓ All tests passed
Packages: 1 passed, 1 total
Tests:    2 passed, 2 total (100% passed)
Time:     1.232s
Ran all tests.
```

---

## Why goherent?

Go's testing tooling is excellent, but two things get awkward as a suite grows:

1. **Assertions are verbose.** Every check is a hand-written `if` plus a `t.Errorf` with a format string. The intent of the test gets buried in boilerplate.
2. **The default report is hard to scan.** `go test -v` produces a flat `=== RUN / --- PASS` stream that's noisy and offers no at-a-glance verdict.

goherent fixes both:

1. **A nicer assertion API, inspired by Jest** — `Expect(value).ToEqual(expected)`, `Expect(xs).ToContainElement(3)`, `Expect(fn).ToPanic()`, and a uniform `Not()` for negation. Test cases also get **real, multi-line descriptions** (Given/When/Then), which Go normally doesn't allow.
2. **A nice report** of the results from the standard `go test` output — colorized, with per-test durations, a pass-rate headline, and the slowest tests.

> **It's a wrapper, not a reinvention.** goherent does **not** replace the Go compiler or test runner — your tests compile and execute through the exact same `go test` toolchain, **so they run just as fast as plain `go test`.** goherent only changes how you *write* assertions and how results are *displayed*.

---

## A look at the output

### Passing run

```
🚀 Starting...

📦 github.com/you/project/math

✅ TestAdd/it adds two numbers (12ms)

✅ TestAdd/it is commutative (1.20s)

✓ All tests passed
Packages: 1 passed, 1 total
Tests:    2 passed, 2 total (100% passed)
Time:     1.232s
Ran all tests.

🐢 2 slowest tests:
  (1.20s) TestAdd/it is commutative
  (12ms) TestAdd/it adds two numbers
```

Durations are dimmed, and anything ≥ 1s is highlighted so slow tests stand out.

### Failing run

When an assertion fails, goherent shows the source location and a readable diff, then a red verdict and the list of failing tests:

```
🚀 Starting...

📦 github.com/you/project/math

❌ TestDivide/it errors on divide by zero (1ms)
    /you/project/math/divide_test.go:18
      not equal:
      expected: 6
      actual  : 7

Failed tests:

❌ github.com/you/project/math
  ● TestDivide/it errors on divide by zero

✗ Tests failed
Packages: 1 failed, 1 total
Tests:    0 passed, 1 total (0% passed)
Time:     0.004s
Ran all tests.
```

---

## Installation

goherent has two parts: a **runner** that executes your tests and shows the report, and a **library** (`test` + `expect`) you import in your test files.

**Requirements:** Go 1.22+.

Add goherent to your module:

```bash
go get github.com/redjolr/goherent
```

You can then run the test runner straight from the module — no separate install step needed:

```bash
go run github.com/redjolr/goherent ./...
```

If you'd rather have a `goherent` command on your `PATH`, install it once:

```bash
go install github.com/redjolr/goherent@latest
```

This puts a `goherent` binary in your Go bin directory (usually `~/go/bin` — make sure it's on your `PATH`), after which you can run:

```bash
goherent ./...
```

---

## Quick start

Write tests with the `Test` function and the injected `Expect`:

```go
package math_test

import (
	"testing"

	. "github.com/redjolr/goherent/test" // dot-import so `Test` is unqualified
	"github.com/redjolr/goherent/expect" // for the `expect.F` parameter type

	"github.com/you/project/math"
)

func TestAdd(t *testing.T) {
	Test("it adds two numbers", func(Expect expect.F) {
		Expect(math.Add(2, 3)).ToEqual(5)
	}, t)

	// Descriptions can be multiline — great for Given/When/Then.
	Test(`
		Given two numbers a and b
		When we add them
		Then addition is commutative
	`, func(Expect expect.F) {
		Expect(math.Add(2, 3)).ToEqual(math.Add(3, 2))
	}, t)
}
```

Run the goherent runner instead of `go test`:

```bash
go run github.com/redjolr/goherent ./...
```

That's it — it accepts the same package/flag arguments you'd give `go test`.

> The examples below use `goherent ./...` for brevity. If you haven't installed the binary, substitute `go run github.com/redjolr/goherent ./...` anywhere you see `goherent`.

---

## Running your tests

**Every `go test` flag works** — goherent forwards your arguments straight through to `go test` unchanged, so packages, build flags, test flags, coverage, profiling, and the rest behave exactly as they do normally:

```bash
goherent ./...                          # whole module
goherent ./math/...                     # a subtree
goherent -run TestAdd ./math            # filter by name
goherent -count=1 ./...                 # disable test caching
goherent -race ./...                    # race detector
goherent -coverprofile=cover.out ./...  # coverage
goherent -timeout 30s -shuffle on ./... # any other go test flags
```

**Concurrent vs. sequential.** By default `go test` runs packages in parallel. Force sequential execution with `-p 1`:

```bash
goherent -p 1 ./...           # one package at a time
```

**Verbosity.** You don't need `-v`; goherent ignores it (the report is always descriptive).

**CI / non-TTY.** When the `CI` environment variable is `true`, goherent prints plain, readable output that stays clean in pipeline logs:

```bash
CI=true goherent ./...
```

---

## The test API

Import the package with a dot-import so the helpers read naturally:

```go
import . "github.com/redjolr/goherent/test"
```

### `Test(name string, body func(Expect expect.F), t *testing.T)`

Defines one test case. `name` is any string (including multiline). `body` receives `Expect`, the assertion entrypoint. Pass the real `*testing.T` as the third argument.

```go
func TestThing(t *testing.T) {
	Test("it works", func(Expect expect.F) {
		Expect(Thing()).ToBeTrue()
	}, t)
}
```

Each `Test` runs as a Go subtest (`t.Run`), so it's isolated and shows up individually in the report.

### `TestSkip(name string, body func(Expect expect.F), t *testing.T)`

Same signature as `Test`, but the case is skipped (`t.Skip()`). Handy for temporarily parking a case while keeping its description.

```go
TestSkip("it handles the edge case (TODO)", func(Expect expect.F) {
	Expect(Edge()).ToEqual(want)
}, t)
```

---

## Assertions

Every assertion starts with `Expect(actual)` and reads as a sentence. A failing assertion reports the file:line and a message; it does **not** stop the test, so multiple expectations can report in one run.

### Negation — `Not()`

Any matcher can be inverted with `Not()`. There's a single, uniform negation path, so every matcher — including ones added in the future — gets its inverse for free:

```go
Expect(2 + 2).Not().ToEqual(5)
Expect(users).Not().ToContainElement("banned")
Expect(m).Not().ToHaveKey("secret")
Expect(value).Not().ToBeNil()
```

`Not().Not()` is the positive matcher again. The aliases `NotToEqual`, `NotToBeError`, and `NotToBeNil` exist as shorthands for the common cases.

### Equality

| Matcher | Checks |
|---|---|
| `Expect(a).ToEqual(b)` | deep equality (handles structs, slices, maps, etc.) |
| `Expect(a).NotToEqual(b)` | `a` is not deeply equal to `b` (alias for `Not().ToEqual(b)`) |

```go
Expect(user).ToEqual(User{Name: "Ada", Age: 36})
Expect(got).NotToEqual(unwanted)
```

### Booleans & nil

| Matcher | Checks |
|---|---|
| `Expect(x).ToBeTrue()` | `x == true` |
| `Expect(x).ToBeFalse()` | `x == false` |
| `Expect(x).ToBeNil()` | `x` is nil (including typed nil pointers, slices, maps, …) |
| `Expect(x).NotToBeNil()` | `x` is not nil |

```go
Expect(cache.Has(key)).ToBeTrue()
Expect(result).ToBeNil()
```

### Errors

| Matcher | Checks |
|---|---|
| `Expect(err).ToBeError()` | the value implements `error` (and is non-nil) |
| `Expect(err).NotToBeError()` | the value is nil or not an `error` |

```go
_, err := Parse("bad")
Expect(err).ToBeError()

_, err = Parse("ok")
Expect(err).NotToBeError()
```

### Numbers & ordering

Works across Go's numeric kinds (ints, uints, floats), plus comparable types like strings, `time.Time`, and `[]byte` where ordering is defined.

| Matcher | Checks |
|---|---|
| `Expect(n).ToBeGreaterThan(m)` | `n > m` |
| `Expect(n).ToBeGreaterThanOrEqualTo(m)` | `n >= m` |
| `Expect(n).ToBeLessThan(m)` | `n < m` |
| `Expect(n).ToBeLessThanOrEqualTo(m)` | `n <= m` |
| `Expect(n).ToBePositive()` | `n > 0` |
| `Expect(n).ToBeNegative()` | `n < 0` |
| `Expect(n).ToBeCloseTo(target, tolerance)` | `|n - target| <= tolerance` (float tolerance) |

```go
Expect(score).ToBeGreaterThanOrEqualTo(60)
Expect(balance).ToBePositive()
Expect(3.14159).ToBeCloseTo(3.14, 0.01) // avoids float == pitfalls
```

### Strings & regex

| Matcher | Checks |
|---|---|
| `Expect(s).ToBeString()` | the value is a `string` |
| `Expect(s).ToContain(sub)` | string contains substring `sub` |
| `Expect(s).ToMatch(pattern)` | string matches the regular expression `pattern` |

```go
Expect(banner).ToContain("goherent")
Expect(version).ToMatch(`^v\d+\.\d+\.\d+$`)
```

### Collections & maps

| Matcher | Checks |
|---|---|
| `Expect(xs).ToContain(v)` | substring (string), **element** (slice/array), or **key** (map) |
| `Expect(xs).ToContainElement(v)` | slice/array/map **value** membership (not substrings) |
| `Expect(m).ToHaveKey(k)` | map contains key `k` |

`ToContain` is the flexible, do-what-I-mean matcher; `ToContainElement` (collection **values**) and `ToHaveKey` (map **keys**) are the precise ones.

```go
Expect([]int{1, 2, 3}).ToContainElement(2)
Expect("hello world").ToContain("world")
Expect(map[string]int{"a": 1}).ToHaveKey("a")
Expect(scores).Not().ToContainElement(0)
```

### Length

Works on strings, slices, arrays, maps, and channels.

| Matcher | Checks |
|---|---|
| `Expect(xs).ToHaveLength(n)` | length is exactly `n` |
| `Expect(xs).ToHaveLengthGreaterThan(n)` | length `> n` |
| `Expect(xs).ToHaveLengthLessThan(n)` | length `< n` |

```go
Expect(items).ToHaveLength(3)
Expect(results).ToHaveLengthGreaterThan(0)
```

### Types

| Matcher | Checks |
|---|---|
| `Expect(x).ToBeOfSameTypeAs(y)` | `x` and `y` have the same dynamic type |

```go
Expect(got).ToBeOfSameTypeAs(User{})
```

### Panics

The value under test must be a no-argument function; the assertion calls it and checks whether it panicked.

| Matcher | Checks |
|---|---|
| `Expect(fn).ToPanic()` | calling `fn()` panics |
| `Expect(fn).Not().ToPanic()` | calling `fn()` returns normally |

```go
Expect(func() { MustParse("nope") }).ToPanic()
Expect(func() { MustParse("ok") }).Not().ToPanic()
```

---

## FAQ

**Do I have to rewrite my existing tests?**
No. Plain `go test` tests still run under `goherent`. Adopt the `Test`/`Expect` API incrementally where it helps.

**Can I still use `t` directly inside a `Test`?**
The body receives `Expect`; if you need `t`, capture it from the enclosing function — but most checks read better through `Expect`.

**Does it work in CI?**
Yes. Set `CI=true` to get clean, plain output suited to pipeline logs.
