package concurrent_events

import (
	"fmt"
	"math"
	"strings"

	"github.com/redjolr/goherent/cmd/concurrent_events/templates"
	"github.com/redjolr/goherent/cmd/ctests_tracker"
	"github.com/redjolr/goherent/internal/utils"
	"github.com/redjolr/goherent/terminal"
	"github.com/redjolr/goherent/terminal/ansi_escape"
	"github.com/redjolr/goherent/terminal/spinner"
)

type Presenter struct {
	terminal     terminal.Terminal
	spinnerFrame int
}

func NewPresenter(term terminal.Terminal) Presenter {
	return Presenter{
		terminal: term,
	}
}

// AdvanceSpinner moves the running-package indicator to its next animation frame.
func (p *Presenter) AdvanceSpinner() {
	p.spinnerFrame++
}

func (p *Presenter) spinnerIcon() string {
	return spinner.Frame(p.spinnerFrame)
}

func (p *Presenter) IsViewPortLarge() bool {
	return p.terminal.Height() > utils.StrLinesCount(templates.RunningTestsSummary("", "", ""))+2
}

func (p *Presenter) TestingStarted() {
	if p.terminal.Height() < math.MaxInt {
		p.terminal.Print(strings.Repeat("\n", p.terminal.Height()))
	} else {
		p.terminal.Print(strings.Repeat("\n", 30))

	}
	p.terminal.Print("\n🚀 Starting...")
}

func (p *Presenter) TestingFinishedSummaryLabel() {
	p.terminal.Print("📋 Tests summary:\n\n")
}

func (p *Presenter) DisplayFinishedPackages(packages []*ctests_tracker.PackageUnderTest) {
	for i, packageUt := range packages {
		if i != 0 {
			p.terminal.Print("\n")
		}
		if packageUt.HasPassed() {
			p.terminal.Print("✅ " + packageUt.Name())
		} else if packageUt.HasBuildFailure() {
			p.terminal.Print("❌ " + packageUt.Name() + "  " + ansi_escape.RED + "[build failed]" + ansi_escape.COLOR_RESET)
			if packageUt.BuildOutput() != "" {
				p.terminal.Print("\n\n  " + packageUt.BuildOutput())
			}
			p.terminal.Print("\n")
		} else if packageUt.HasAtLeastOneFailedTest() {
			p.terminal.Print("❌ " + packageUt.Name())
			for _, ctest := range packageUt.FailedCtests() {
				p.terminal.Print("\n\n")
				p.terminal.Print("  " + ansi_escape.RED + "● " + ctest.Name() + ansi_escape.COLOR_RESET)
				if ctest.ContainsOutput() {
					p.terminal.Print("\n\n")
					p.terminal.Print("  " + ctest.Output())
				}
			}

			if packageUt.HasOutputOfParentTests() {
				p.terminal.Print("\n\n" + packageUt.ParentTestsOutput())
			}

			p.terminal.Print("\n")

		} else if packageUt.IsSkipped() {
			p.terminal.Print("⏩ " + packageUt.Name())
		}
	}
}

func (p *Presenter) DisplayPackages(
	runningPackages []*ctests_tracker.PackageUnderTest,
	finishedPackages []*ctests_tracker.PackageUnderTest,
) {
	if p.IsViewPortLarge() {
		p.displayPackagesInLargeTerminal(runningPackages, finishedPackages)
	} else {
		p.displayPackagesInSmallTerminal(runningPackages, finishedPackages)
	}
}

func (p *Presenter) displayPackagesInLargeTerminal(
	runningPackages []*ctests_tracker.PackageUnderTest,
	finishedPackages []*ctests_tracker.PackageUnderTest,
) {
	packagesThatFitInTerminalCount := p.terminal.Height() - utils.StrLinesCount(templates.RunningTestsSummary("", "", "")) - 2
	runningPackagesThatFitInTerminal := runningPackages[0:min(len(runningPackages), packagesThatFitInTerminalCount)]

	if len(runningPackages) < packagesThatFitInTerminalCount && len(finishedPackages) > 0 {
		showFinishedPackagesCount := min(len(finishedPackages), packagesThatFitInTerminalCount-len(runningPackages))
		latestFinishedPackages := finishedPackages[len(finishedPackages)-showFinishedPackagesCount:]
		for i, packageUt := range latestFinishedPackages {
			if i != 0 {
				p.terminal.Print("\n")
			}
			if packageUt.HasPassed() {
				p.terminal.Print("✅ " + packageUt.Name())
			} else if packageUt.HasAtLeastOneFailedTest() || packageUt.HasBuildFailure() {
				p.terminal.Print("❌ " + packageUt.Name())
			} else if packageUt.IsSkipped() {
				p.terminal.Print("⏩ " + packageUt.Name())
			}
		}
		if len(runningPackages) > 0 {
			p.terminal.Print("\n")
		}
	}

	for i, packageUt := range runningPackagesThatFitInTerminal {
		if i != 0 {
			p.terminal.Print("\n")
		}
		p.terminal.Print(p.spinnerIcon() + " " + packageUt.Name())
	}
}

func (p *Presenter) displayPackagesInSmallTerminal(
	runningPackages []*ctests_tracker.PackageUnderTest,
	finishedPackages []*ctests_tracker.PackageUnderTest,
) {
	runningPackagesThatFitInTerminal := runningPackages[0:min(len(runningPackages), p.terminal.Height())]

	if len(runningPackages) < p.terminal.Height() && len(finishedPackages) > 0 {
		showFinishedPackagesCount := min(len(finishedPackages), p.terminal.Height()-len(runningPackages))
		latestFinishedPackages := finishedPackages[len(finishedPackages)-showFinishedPackagesCount:]
		for i, packageUt := range latestFinishedPackages {
			if i != 0 {
				p.terminal.Print("\n")
			}
			if packageUt.HasPassed() {
				p.terminal.Print("✅ " + packageUt.Name())
			} else if packageUt.HasAtLeastOneFailedTest() || packageUt.HasBuildFailure() {
				p.terminal.Print("❌ " + packageUt.Name())
			} else if packageUt.IsSkipped() {
				p.terminal.Print("⏩ " + packageUt.Name())
			}
		}
		if len(runningPackages) > 0 {
			p.terminal.Print("\n")
		}
	}

	for i, packageut := range runningPackagesThatFitInTerminal {
		if i != 0 {
			p.terminal.Print("\n")
		}
		p.terminal.Print(p.spinnerIcon() + " " + packageut.Name())
	}
}

func (p *Presenter) RunningTestsSummary(testingSummary ctests_tracker.TestingSummary) {
	packagesSummary := ansi_escape.BOLD + "Packages:" + ansi_escape.RESET_BOLD + " "
	testsSummary := ansi_escape.BOLD + "Tests:" + ansi_escape.RESET_BOLD + "    "
	timeSummary := fmt.Sprintf(ansi_escape.BOLD+"Time:"+ansi_escape.RESET_BOLD+"     %.3fs", testingSummary.DurationS)
	runningPackagesCount := testingSummary.RunningPackagesCount
	passedPackagesCount := testingSummary.PassedPackagesCount
	failedPackagesCount := testingSummary.FailedPackagesCount
	skippedPackagesCount := testingSummary.SkippedPackagesCount
	passedTestsCount := testingSummary.PassedTestsCount
	failedTestsCount := testingSummary.FailedTestsCount
	skippedTestsCount := testingSummary.SkippedTestsCount

	packagesSummary += fmt.Sprintf("%d running", runningPackagesCount)

	if failedPackagesCount > 0 {
		packagesSummary += ", " + ansi_escape.RED + fmt.Sprintf("%d failed", failedPackagesCount) + ansi_escape.COLOR_RESET
	}
	if skippedPackagesCount > 0 {
		packagesSummary += ", " + ansi_escape.YELLOW + fmt.Sprintf("%d skipped", skippedPackagesCount) + ansi_escape.COLOR_RESET
	}
	if passedPackagesCount > 0 {
		packagesSummary += ", " + ansi_escape.GREEN + fmt.Sprintf("%d passed", passedPackagesCount) + ansi_escape.COLOR_RESET
	}
	if failedTestsCount > 0 {
		testsSummary += ansi_escape.RED + fmt.Sprintf("%d failed", failedTestsCount) + ansi_escape.COLOR_RESET + ", "
	}
	if skippedTestsCount > 0 {
		testsSummary += ansi_escape.YELLOW + fmt.Sprintf("%d skipped", skippedTestsCount) + ansi_escape.COLOR_RESET + ", "
	}
	if passedTestsCount > 0 {
		testsSummary += ansi_escape.GREEN + fmt.Sprintf("%d passed", passedTestsCount) + ansi_escape.COLOR_RESET + ", "
	}

	testsSummary, _ = strings.CutSuffix(testsSummary, ", ")

	p.terminal.Print(templates.RunningTestsSummary(packagesSummary, testsSummary, timeSummary))
}

func (p *Presenter) TestingFinishedSummary(summary ctests_tracker.TestingSummary) {
	packagesSummary := ansi_escape.BOLD + "Packages:" + ansi_escape.RESET_BOLD + " "
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
	testsSummary += fmt.Sprintf("%d total", summary.TestsCount)

	p.terminal.Print(
		templates.TestingFinishedSummary(packagesSummary, testsSummary, timeSummary),
	)
}

// slowTestThresholdS is the elapsed time (in seconds) at or above which a test's
// duration is highlighted in yellow in the slowest-tests report.
const slowTestThresholdS = 1.0

func (p *Presenter) SlowestTests(tests []*ctests_tracker.Ctest) {
	if len(tests) == 0 {
		return
	}
	p.terminal.Print("\n\n" + buildSlowestTestsReport(tests) + "\n\n")
}

// buildSlowestTestsReport renders the "N slowest tests" block shown at the end
// of a run: a header followed by one line per test with its (colored) duration
// and the first line of its name.
func buildSlowestTestsReport(tests []*ctests_tracker.Ctest) string {
	noun := "tests"
	if len(tests) == 1 {
		noun = "test"
	}
	report := fmt.Sprintf("🐢 %d slowest %s:", len(tests), noun)
	for _, ctest := range tests {
		report += "\n  " + slowestDurationLabel(ctest.DurationS()) + " " + firstLine(ctest.Name())
	}
	return report
}

func slowestDurationLabel(seconds float64) string {
	color := ansi_escape.DIM
	if seconds >= slowTestThresholdS {
		color = ansi_escape.YELLOW
	}
	return color + "(" + utils.FormatDuration(seconds) + ")" + ansi_escape.COLOR_RESET
}

func firstLine(s string) string {
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		return s[:i]
	}
	return s
}

func (p *Presenter) EraseScreen() {
	p.terminal.Print(ansi_escape.CURSOR_TO_HOME)
	p.terminal.Print(ansi_escape.ERASE_SCREEN)
}

func (p *Presenter) Error() {
	p.terminal.Print("\n\n❗ Error.")
}
