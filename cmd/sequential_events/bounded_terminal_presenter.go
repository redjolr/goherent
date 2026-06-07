package sequential_events

import (
	"fmt"
	"strings"

	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/internal/utils"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
)

type BoundedTerminalPresenter struct {
	terminal terminal.Terminal
}

func NewBoundedTerminalPresenter(term terminal.Terminal) BoundedTerminalPresenter {
	return BoundedTerminalPresenter{
		terminal: term,
	}
}

func (tp BoundedTerminalPresenter) TestingStarted() {
	tp.terminal.Print("\n🚀 Starting...")
}

func (tp BoundedTerminalPresenter) PackageTestsStartedRunning(packageName string) {
	tp.terminal.Print(fmt.Sprintf("\n\n📦 %s\n", packageName))
}

func (tp BoundedTerminalPresenter) CtestStartedRunning(ctest *ctests_tracker.Ctest) {
	if utils.StrLinesCount(ctest.Name()) > tp.terminal.Height() {
		printableName := strings.Join(utils.SplitStringByNewLine(ctest.Name())[0:tp.terminal.Height()], "\n")
		tp.terminal.Print(fmt.Sprintf("\n⏳ %s...", printableName))
	} else {
		tp.terminal.Print(fmt.Sprintf("\n⏳ %s", ctest.Name()))
	}
}

func (tp BoundedTerminalPresenter) CtestPassed(ctest *ctests_tracker.Ctest, duration float64) {
	hourGlassAndSpaceLength := utils.DisplayWidth("⏳ ")
	testNameLineCount := utils.StrLinesCount(ctest.Name())
	threeDotsLineCount := len("...")

	if utils.StrLinesCount(ctest.Name()) > tp.terminal.Height() {
		printedName := strings.Join(utils.SplitStringByNewLine(ctest.Name())[0:tp.terminal.Height()], "\n")
		unprintedName := strings.Join(utils.SplitStringByNewLine(ctest.Name())[tp.terminal.Height():], "\n")
		printedNameLines := utils.SplitStringByNewLine(printedName)
		lastLine := printedNameLines[len(printedNameLines)-1]
		lastLineLength := utils.DisplayWidth(lastLine)
		tp.terminal.MoveLeft(threeDotsLineCount + lastLineLength + hourGlassAndSpaceLength)
		if testNameLineCount > 1 {
			tp.terminal.MoveUp(tp.terminal.Height())
		}
		tp.terminal.Print("✅ ")
		tp.terminal.Print(printedNameLines[0] + formatDurationLabel(duration))
		if len(printedNameLines) > 1 {
			tp.terminal.Print("\n" + strings.Join(printedNameLines[1:], "\n"))
		}
		tp.terminal.Print("   ")
		tp.terminal.Print("\n")
		tp.terminal.Print(unprintedName)
	} else {
		nameLines := utils.SplitStringByNewLine(ctest.Name())
		lastLine := nameLines[len(nameLines)-1]
		lastLineLength := utils.DisplayWidth(lastLine)
		tp.terminal.MoveLeft(lastLineLength + hourGlassAndSpaceLength)
		if testNameLineCount > 1 {
			tp.terminal.MoveUp(testNameLineCount - 1)
		}
		tp.terminal.Print("✅")
		tp.printDurationAfterNameHead(nameLines[0], duration)
		if testNameLineCount > 1 {
			tp.terminal.MoveDown(testNameLineCount - 1)
		}
	}
}

func (tp BoundedTerminalPresenter) Print(output string) {
	tp.terminal.Print(output)
}

// printDurationAfterNameHead writes the duration label on the test's first line,
// right after the name head, with the cursor positioned just past the status
// icon on that line. The label goes here — not after the name's last line —
// because a multi-line BDD name's last line is often long enough that an
// appended label would wrap off the right edge of the terminal. The first line
// (the test function name) is short, so there is reliably room.
//
// The "+1" steps over the single space that separates the icon from the name
// ("✅ <name>"), landing the cursor on the empty cell just past the name head;
// the label's own leading space then forms the gap before "(…)".
func (tp BoundedTerminalPresenter) printDurationAfterNameHead(firstLine string, duration float64) {
	label := formatDurationLabel(duration)
	if label == "" {
		return
	}
	tp.terminal.MoveRight(utils.DisplayWidth(firstLine) + 1)
	tp.terminal.Print(label)
}

func (tp BoundedTerminalPresenter) CtestFailed(ctest *ctests_tracker.Ctest, duration float64) {
	hourGlassAndSpaceLength := utils.DisplayWidth("⏳ ")
	testNameLineCount := utils.StrLinesCount(ctest.Name())
	threeDotsLineCount := len("...")

	if utils.StrLinesCount(ctest.Name()) > tp.terminal.Height() {
		printedName := strings.Join(utils.SplitStringByNewLine(ctest.Name())[0:tp.terminal.Height()], "\n")
		unprintedName := strings.Join(utils.SplitStringByNewLine(ctest.Name())[tp.terminal.Height():], "\n")
		printedNameLines := utils.SplitStringByNewLine(printedName)
		lastLine := printedNameLines[len(printedNameLines)-1]
		lastLineLength := utils.DisplayWidth(lastLine)
		tp.terminal.MoveLeft(threeDotsLineCount + lastLineLength + hourGlassAndSpaceLength)
		if testNameLineCount > 1 {
			tp.terminal.MoveUp(tp.terminal.Height())
		}
		tp.terminal.Print("❌ ")
		tp.terminal.Print(printedNameLines[0] + formatDurationLabel(duration))
		if len(printedNameLines) > 1 {
			tp.terminal.Print("\n" + strings.Join(printedNameLines[1:], "\n"))
		}
		tp.terminal.Print("   ")
		tp.terminal.Print("\n")
		tp.terminal.Print(unprintedName)
	} else {
		nameLines := utils.SplitStringByNewLine(ctest.Name())
		lastLine := nameLines[len(nameLines)-1]
		lastLineLength := utils.DisplayWidth(lastLine)
		tp.terminal.MoveLeft(lastLineLength + hourGlassAndSpaceLength)
		if testNameLineCount > 1 {
			tp.terminal.MoveUp(testNameLineCount - 1)
		}
		tp.terminal.Print("❌")
		tp.printDurationAfterNameHead(nameLines[0], duration)
		if testNameLineCount > 1 {
			tp.terminal.MoveDown(testNameLineCount - 1)
		}
	}
}

func (tp BoundedTerminalPresenter) CtestSkipped(ctest *ctests_tracker.Ctest) {
	hourGlassAndSpaceLength := utils.DisplayWidth("⏳ ")
	testNameLineCount := utils.StrLinesCount(ctest.Name())
	threeDotsLineCount := len("...")

	if utils.StrLinesCount(ctest.Name()) > tp.terminal.Height() {
		printedName := strings.Join(utils.SplitStringByNewLine(ctest.Name())[0:tp.terminal.Height()], "\n")
		unprintedName := strings.Join(utils.SplitStringByNewLine(ctest.Name())[tp.terminal.Height():], "\n")
		printedNameLines := utils.SplitStringByNewLine(printedName)
		lastLine := printedNameLines[len(printedNameLines)-1]
		lastLineLength := utils.DisplayWidth(lastLine)
		tp.terminal.MoveLeft(threeDotsLineCount + lastLineLength + hourGlassAndSpaceLength)
		if testNameLineCount > 1 {
			tp.terminal.MoveUp(tp.terminal.Height())
		}
		tp.terminal.Print("⏩ ")

		tp.terminal.Print(printedName + "   ")
		tp.terminal.Print("\n")
		tp.terminal.Print(unprintedName)
	} else {
		nameLines := utils.SplitStringByNewLine(ctest.Name())
		lastLine := nameLines[len(nameLines)-1]
		lastLineLength := utils.DisplayWidth(lastLine)
		tp.terminal.MoveLeft(lastLineLength + hourGlassAndSpaceLength)
		if testNameLineCount > 1 {
			tp.terminal.MoveUp(testNameLineCount - 1)
		}
		tp.terminal.Print("⏩")
		if testNameLineCount > 1 {
			tp.terminal.MoveDown(testNameLineCount - 1)
		}
		tp.terminal.MoveRight(lastLineLength)
	}
}

func (tp BoundedTerminalPresenter) CtestOutput(ctest *ctests_tracker.Ctest) {
	tp.terminal.Print("\n" + ctest.Output())
}

func (tp BoundedTerminalPresenter) FailedTestsList(failedPackages []*ctests_tracker.PackageUnderTest) {
	tp.terminal.Print("\n\nFailed tests:")
	for _, packageUt := range failedPackages {
		tp.terminal.Print("\n\n❌ " + ansi_escape.BOLD + ansi_escape.RED + packageUt.Name() + ansi_escape.COLOR_RESET)
		for _, ctest := range packageUt.FailedCtests() {
			tp.terminal.Print("\n\n  " + ansi_escape.RED + "● " + ctest.Name() + ansi_escape.COLOR_RESET)
			if ctest.ContainsOutput() {
				tp.terminal.Print("\n\n  " + ctest.Output())
			}
		}
		if packageUt.HasOutputOfParentTests() {
			tp.terminal.Print("\n\n" + packageUt.ParentTestsOutput())
		}
	}
	tp.terminal.Print("\n")
}

func (tp BoundedTerminalPresenter) Error() {
	tp.terminal.Print("\n\n❗ Error.")
}

func (tp BoundedTerminalPresenter) TestingFinishedSummary(summary ctests_tracker.TestingSummary) {

	packagesSummary := ansi_escape.BOLD + "\nPackages:" + ansi_escape.RESET_BOLD + " "
	testsSummary := ansi_escape.BOLD + "Tests:" + ansi_escape.RESET_BOLD + "    "
	timeSummary := fmt.Sprintf(ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     %.3fs", summary.DurationS)

	if summary.FailedPackagesCount > 0 {
		packagesSummary += ansi_escape.RED +
			fmt.Sprintf("%d failed", summary.FailedPackagesCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.SkippedPackagesCount > 0 {
		packagesSummary += ansi_escape.YELLOW +
			fmt.Sprintf("%d skipped", summary.SkippedPackagesCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.PassedPackagesCount > 0 {
		packagesSummary += ansi_escape.GREEN +
			fmt.Sprintf("%d passed", summary.PassedPackagesCount) +
			ansi_escape.COLOR_RESET + ", "
	}

	if summary.FailedTestsCount > 0 {
		testsSummary += ansi_escape.RED +
			fmt.Sprintf("%d failed", summary.FailedTestsCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.SkippedTestsCount > 0 {
		testsSummary += ansi_escape.YELLOW +
			fmt.Sprintf("%d skipped", summary.SkippedTestsCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	if summary.PassedTestsCount > 0 {
		testsSummary += ansi_escape.GREEN +
			fmt.Sprintf("%d passed", summary.PassedTestsCount) +
			ansi_escape.COLOR_RESET + ", "
	}
	packagesSummary += fmt.Sprintf("%d total", summary.PackagesCount)
	testsSummary += fmt.Sprintf("%d total", summary.TestsCount) + passRateLabel(summary)

	tp.terminal.Print("\n" + testingVerdictHeadline(summary) + "\n")
	tp.terminal.Print(
		packagesSummary + "\n" +
			testsSummary + "\n" +
			timeSummary + "\n" +
			"Ran all tests.",
	)
}
