package ctests_tracker

import (
	"slices"

	"github.com/redjolr/goherent/cmd/events"
)

type PackageUnderTest struct {
	name                    string
	ctests                  []Ctest
	outputEvtsOfParentTests []events.CtestOutputEvent
	testingFinished         bool
}

func NewPackageUnderTest(name string) PackageUnderTest {
	newPack := PackageUnderTest{
		name:                    name,
		ctests:                  []Ctest{},
		outputEvtsOfParentTests: []events.CtestOutputEvent{},
		testingFinished:         false,
	}
	return newPack
}

func (packageUt *PackageUnderTest) Name() string {
	return packageUt.name
}

func (packageUt *PackageUnderTest) TestsAreRunning() bool {
	return !packageUt.testingFinished
}

func (packageUt *PackageUnderTest) MarkAsFinished() {
	packageUt.testingFinished = true
}

func (packageUt *PackageUnderTest) RunningCtestsCount() int {
	count := 0
	for _, ctest := range packageUt.ctests {
		if ctest.isRunning {
			count++
		}
	}
	return count
}

func (packageUt *PackageUnderTest) PassedCtestsCount() int {
	count := 0
	for _, ctest := range packageUt.ctests {
		if ctest.hasPassed {
			count++
		}
	}
	return count
}

func (packageUt *PackageUnderTest) FailedCtestsCount() int {
	count := 0
	for _, ctest := range packageUt.ctests {
		if ctest.hasFailed {
			count++
		}
	}
	return count
}

func (packageUt *PackageUnderTest) FailedCtests() []Ctest {
	failedCtestCount := packageUt.FailedCtestsCount()
	failedCtests := make([]Ctest, failedCtestCount, failedCtestCount)
	failedCtestInd := 0
	for _, ctest := range packageUt.ctests {
		if ctest.hasFailed {
			failedCtests[failedCtestInd] = ctest
			failedCtestInd++
		}
	}
	return failedCtests
}

func (packageUt *PackageUnderTest) SkippedCtestsCount() int {
	count := 0
	for _, ctest := range packageUt.ctests {
		if ctest.isSkipped {
			count++
		}
	}
	return count
}

func (packageUt *PackageUnderTest) HasAtLeastOneTest() bool {
	return len(packageUt.ctests) != 0
}

func (packageUt *PackageUnderTest) HasAtLeastOneFailedTest() bool {
	for _, ctest := range packageUt.ctests {
		if ctest.hasFailed {
			return true
		}
	}
	return false
}

func (packageUt *PackageUnderTest) HasAtLeastOnePassedTest() bool {
	for _, ctest := range packageUt.ctests {
		if ctest.hasPassed {
			return true
		}
	}
	return false
}

func (packageUt *PackageUnderTest) HasAtLeastOneSkippedTest() bool {
	for _, ctest := range packageUt.ctests {
		if ctest.isSkipped {
			return true
		}
	}
	return false
}

func (packageUt *PackageUnderTest) HasPassed() bool {
	return !packageUt.TestsAreRunning() && packageUt.PassedCtestsCount() > 0 &&
		packageUt.PassedCtestsCount()+packageUt.SkippedCtestsCount() == len(packageUt.ctests)
}

func (packageUt *PackageUnderTest) IsSkipped() bool {
	return !packageUt.TestsAreRunning() && packageUt.SkippedCtestsCount() == len(packageUt.ctests)
}

func (packageUt *PackageUnderTest) CtestsCount() int {
	return len(packageUt.ctests)
}

func (p *PackageUnderTest) RecordOutputEvtOfParentTest(evt events.CtestOutputEvent) {
	p.outputEvtsOfParentTests = append(p.outputEvtsOfParentTests, evt)
}

func (p *PackageUnderTest) ParentTestsOutput() string {
	output := ""
	for _, outputEvt := range p.outputEvtsOfParentTests {
		output += outputEvt.Output
	}
	return output
}

func (p *PackageUnderTest) HasOutputOfParentTests() bool {
	return p.ParentTestsOutput() != ""
}

func (packageUt *PackageUnderTest) isCtestTheFirstOne(ctest Ctest) bool {
	if len(packageUt.ctests) == 0 {
		return false
	}
	return packageUt.ctests[0].HasName(ctest.name)
}

func (packageUt *PackageUnderTest) insertCtest(ctest Ctest) Ctest {
	if !packageUt.containsCtest(ctest.name) {
		packageUt.ctests = append(packageUt.ctests, ctest)
		return ctest
	}
	panic("Ctest already exists")
}

func (packageUt *PackageUnderTest) ctestByName(ctestName string) *Ctest {
	indexOfCtestWithName := slices.IndexFunc(packageUt.ctests, func(aCtest Ctest) bool {
		return aCtest.HasName(ctestName)
	})
	if indexOfCtestWithName != -1 {
		return &packageUt.ctests[indexOfCtestWithName]
	}
	return nil
}

func (packageUt *PackageUnderTest) containsCtest(ctestName string) bool {
	indexOfCtestWithName := slices.IndexFunc(packageUt.ctests, func(aCtest Ctest) bool {
		return aCtest.HasName(ctestName)
	})
	return indexOfCtestWithName != -1
}
