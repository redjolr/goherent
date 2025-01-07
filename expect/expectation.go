package expect

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/redjolr/goherent/expect/internal/assertions"
	"github.com/redjolr/goherent/internal/utils"
	"github.com/redjolr/goherent/terminal/ansi_escape"
)

type expectation struct {
	t                       *testing.T
	checkExpectationAgainst any
}

func (e *expectation) ToEqual(actual any) {
	if err := assertions.ToEqual(e.checkExpectationAgainst, actual); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}

func (e *expectation) ToContain(containee any) {
	if err := assertions.ToContain(e.checkExpectationAgainst, containee); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}

func (e *expectation) ToBeError() {
	if err := assertions.ToBeError(e.checkExpectationAgainst); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}

func (e *expectation) NotToBeError() {
	if err := assertions.NotToBeError(e.checkExpectationAgainst); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}

func (e *expectation) ToBeTrue() {
	if err := assertions.ToBeTrue(e.checkExpectationAgainst); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}

func (e *expectation) ToBeFalse() {
	if err := assertions.ToBeFalse(e.checkExpectationAgainst); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}

func (e *expectation) ToBeNil() {
	if err := assertions.ToBeNil(e.checkExpectationAgainst); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}

func (e *expectation) NotToBeNil() {
	if err := assertions.NotToBeNil(e.checkExpectationAgainst); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}

func (e *expectation) ToBeOfSameTypeAs(compareVal any) {
	if err := assertions.ToBeOfSameTypeAs(e.checkExpectationAgainst, compareVal); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}

func (e *expectation) print(str string, leftPadWhitespace int) {
	lines := utils.SplitStringByNewLine(str)
	for _, line := range lines {
		fmt.Println(strings.Repeat(" ", leftPadWhitespace) + line)
	}
}

func (e *expectation) ToBeString() {
	if err := assertions.ToBeString(e.checkExpectationAgainst); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}

func (e *expectation) ToBeGreaterThan(actual any) {
	if err := assertions.ToBeGreaterThan(actual, e.checkExpectationAgainst); err != nil {
		_, file, line, _ := runtime.Caller(1)
		e.print(fmt.Sprintf(ansi_escape.YELLOW+"%s:%d"+ansi_escape.COLOR_RESET, file, line), 4)
		e.print(err.Error(), 6)
		e.t.Fail()
	}
}
