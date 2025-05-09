package ctests_tracker

import (
	"slices"
	"time"

	"github.com/redjolr/goherent/cmd/events"
)

type CtestsTracker struct {
	packagesUnderTest []*PackageUnderTest
	testingStartedAt  time.Time
	testingFinishedAt time.Time
}

func NewCtestsTracker() CtestsTracker {
	return CtestsTracker{
		packagesUnderTest: []*PackageUnderTest{},
		testingStartedAt:  time.Time{},
		testingFinishedAt: time.Time{},
	}
}

func (tracker *CtestsTracker) InsertCtest(ctest Ctest) Ctest {
	if tracker.ContainsPackageUtWithName(ctest.packageName) {
		existingPackUt := tracker.PackageUnderTest(ctest.packageName)
		existingPackUt.insertCtest(ctest)
		tracker.replacePackageWith(ctest.packageName, existingPackUt)
	} else {
		packUt := NewPackageUnderTest(ctest.packageName)
		packUt.insertCtest(ctest)
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, &packUt)
	}

	return ctest
}

func (tracker *CtestsTracker) Packages() []*PackageUnderTest {
	return tracker.packagesUnderTest
}

func (tracker *CtestsTracker) RunningPackages() []*PackageUnderTest {
	runningPackages := []*PackageUnderTest{}
	for _, pack := range tracker.packagesUnderTest {
		if pack.TestsAreRunning() {
			runningPackages = append(runningPackages, pack)
		}
	}
	return runningPackages
}

func (tracker *CtestsTracker) FinishedPackages() []*PackageUnderTest {
	runningPackages := []*PackageUnderTest{}
	for _, pack := range tracker.packagesUnderTest {
		if !pack.TestsAreRunning() {
			runningPackages = append(runningPackages, pack)
		}
	}
	return runningPackages
}

func (tracker *CtestsTracker) PassedPackages() []*PackageUnderTest {
	runningPackages := []*PackageUnderTest{}
	for _, pack := range tracker.packagesUnderTest {
		if pack.HasPassed() {
			runningPackages = append(runningPackages, pack)
		}
	}
	return runningPackages
}

func (tracker *CtestsTracker) FinishedFailedPackages() []*PackageUnderTest {
	runningPackages := []*PackageUnderTest{}
	for _, pack := range tracker.packagesUnderTest {
		if pack.HasAtLeastOneFailedTest() && !pack.TestsAreRunning() {
			runningPackages = append(runningPackages, pack)
		}
	}
	return runningPackages
}

func (tracker *CtestsTracker) InsertPackageUt(name string) PackageUnderTest {
	existingPackageUt := tracker.FindPackageWithName(name)
	if existingPackageUt != nil {
		return *existingPackageUt
	}
	packUt := NewPackageUnderTest(name)
	tracker.packagesUnderTest = append(tracker.packagesUnderTest, &packUt)
	return packUt
}

func (tracker *CtestsTracker) DeletePackage(packageUt *PackageUnderTest) {
	packInd := slices.Index(tracker.packagesUnderTest, packageUt)
	if packInd != -1 {
		if packInd == len(tracker.packagesUnderTest)-1 {
			tracker.packagesUnderTest = tracker.packagesUnderTest[0:packInd]
		} else {
			tracker.packagesUnderTest = slices.Concat(
				tracker.packagesUnderTest[0:packInd],
				tracker.packagesUnderTest[packInd+1:],
			)
		}

	}
}

func (tracker *CtestsTracker) HandleCtestPassedEvent(evt events.CtestPassedEvent) {
	if !tracker.ContainsPackageUtWithName(evt.PackageName) {
		packUt := NewPackageUnderTest(evt.PackageName)
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, &packUt)
	}
	packUt := tracker.FindPackageWithName(evt.PackageName)

	if !packUt.containsCtest(evt.TestName) {
		packUt.insertCtest(NewPassedCtest(evt))
		return
	}
	ctest := packUt.ctestByName(evt.TestName)
	ctest.MarkAsPassed(evt)
}

func (tracker *CtestsTracker) HandleCtestSkippedEvent(evt events.CtestSkippedEvent) {
	if !tracker.ContainsPackageUtWithName(evt.PackageName) {
		packUt := NewPackageUnderTest(evt.PackageName)
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, &packUt)
	}
	packUt := tracker.FindPackageWithName(evt.PackageName)

	if !packUt.containsCtest(evt.TestName) {
		packUt.insertCtest(NewSkippedCtest(evt))
		return
	}
	ctest := packUt.ctestByName(evt.TestName)
	ctest.MarkAsSkipped(evt)
}

func (tracker *CtestsTracker) HandleCtestFailedEvent(evt events.CtestFailedEvent) {
	if evt.IsEventOfAParentTest() {
		return
	}
	if !tracker.ContainsPackageUtWithName(evt.PackageName) {
		packUt := NewPackageUnderTest(evt.PackageName)
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, &packUt)
	}
	packUt := tracker.FindPackageWithName(evt.PackageName)

	if !packUt.containsCtest(evt.TestName) {
		packUt.insertCtest(NewFailedCtest(evt))
		return
	}
	ctest := packUt.ctestByName(evt.TestName)
	ctest.MarkAsFailed(evt)
}

func (tracker *CtestsTracker) HandleCtestRanEvent(evt events.CtestRanEvent) {
	if !tracker.ContainsPackageUtWithName(evt.PackageName) {
		packUt := NewPackageUnderTest(evt.PackageName)
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, &packUt)
	}

	packUt := tracker.FindPackageWithName(evt.PackageName)
	if !packUt.containsCtest(evt.TestName) {
		packUt.insertCtest(NewRunningCtest(evt))
	}
}

func (tracker *CtestsTracker) HandleCtestOutputEvent(evt events.CtestOutputEvent) {
	if !tracker.ContainsPackageUtWithName(evt.PackageName) {
		packUt := NewPackageUnderTest(evt.PackageName)
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, &packUt)
	}
	packUt := tracker.FindPackageWithName(evt.PackageName)

	if evt.IsEventOfAParentTest() && !evt.IsAGenericRunPassFailOutput() {
		packUt.RecordOutputEvtOfParentTest(evt)
		return
	}
	if evt.IsEventOfAParentTest() {
		return
	}

	if !packUt.containsCtest(evt.TestName) {
		packUt.insertCtest(NewCtest(evt.TestName, evt.PackageName))
	}
	ctest := tracker.FindCtestWithNameInPackage(evt.TestName, evt.PackageName)
	ctest.RecordOutputEvt(evt)
}

func (tracker *CtestsTracker) IsCtestFirstOfItsPackage(ctest Ctest) bool {
	if !tracker.ContainsPackageUtWithName(ctest.packageName) {
		return false
	}
	packageUnderTest := tracker.PackageUnderTest(ctest.packageName)
	return packageUnderTest.isCtestTheFirstOne(ctest)
}

func (tracker *CtestsTracker) ContainsPackageUtWithName(name string) bool {
	indexOfPackUttWithName := slices.IndexFunc(tracker.packagesUnderTest, func(packUt *PackageUnderTest) bool {
		return packUt.name == name
	})
	return indexOfPackUttWithName != -1
}

func (tracker *CtestsTracker) TestingStarted(testingStartedEvt events.TestingStartedEvent) {
	tracker.testingStartedAt = testingStartedEvt.Timestamp
}

func (tracker *CtestsTracker) TestingFinished(testingFinishedEvt events.TestingFinishedEvent) {
	for _, packageUt := range tracker.packagesUnderTest {
		packageUt.MarkAsFinished()
	}
	tracker.testingFinishedAt = testingFinishedEvt.Timestamp
}

func (tracker *CtestsTracker) FindPackageWithName(packageName string) *PackageUnderTest {
	for _, packUt := range tracker.packagesUnderTest {
		if packUt.name == packageName {
			return packUt
		}
	}
	return nil
}

func (tracker *CtestsTracker) PackagesCount() int {
	return len(tracker.packagesUnderTest)
}

func (tracker *CtestsTracker) HasPackages() bool {
	return tracker.PackagesCount() > 0
}

func (tracker *CtestsTracker) CtestsCount() int {
	count := 0
	for _, packageUt := range tracker.packagesUnderTest {
		count += packageUt.CtestsCount()
	}
	return count
}

func (tracker *CtestsTracker) PassedCtestsCount() int {
	count := 0
	for _, packageUt := range tracker.packagesUnderTest {
		count += packageUt.PassedCtestsCount()
	}
	return count
}

func (tracker *CtestsTracker) FailedCtestsCount() int {
	count := 0
	for _, packageUt := range tracker.packagesUnderTest {
		count += packageUt.FailedCtestsCount()
	}
	return count
}

func (tracker *CtestsTracker) SkippedCtestsCount() int {
	count := 0
	for _, packageUt := range tracker.packagesUnderTest {
		count += packageUt.SkippedCtestsCount()
	}
	return count
}

func (tracker *CtestsTracker) PassedPackagesCount() int {
	count := 0
	for _, packageUt := range tracker.packagesUnderTest {
		if packageUt.HasPassed() {
			count += 1
		}
	}
	return count
}

func (tracker *CtestsTracker) FailedPackagesCount() int {
	count := 0
	for _, packageUt := range tracker.packagesUnderTest {
		if packageUt.HasAtLeastOneFailedTest() {
			count += 1
		}
	}
	return count
}

func (tracker *CtestsTracker) SkippedPackagesCount() int {
	count := 0
	for _, packageUt := range tracker.packagesUnderTest {
		if packageUt.IsSkipped() {
			count += 1
		}
	}
	return count
}

func (tracker *CtestsTracker) RunningPackagesCount() int {
	count := 0
	for _, pack := range tracker.packagesUnderTest {
		if pack.TestsAreRunning() {
			count++
		}
	}
	return count
}

func (tracker *CtestsTracker) PackageUnderTest(name string) *PackageUnderTest {
	if tracker.ContainsPackageUtWithName(name) {
		indexOfPackUtWithName := slices.IndexFunc(tracker.packagesUnderTest, func(packUt *PackageUnderTest) bool {
			return packUt.name == name
		})
		return tracker.packagesUnderTest[indexOfPackUtWithName]
	}
	return nil
}

func (tracker *CtestsTracker) FindCtestWithNameInPackage(ctestName string, packageName string) *Ctest {
	for _, packUt := range tracker.packagesUnderTest {
		if packUt.name == packageName {
			ctest := packUt.ctestByName(ctestName)
			if ctest != nil {
				return ctest
			}
		}
	}
	return nil
}

func (tracker *CtestsTracker) RunningCtestsCount() int {
	count := 0
	for _, packageUt := range tracker.packagesUnderTest {
		count += packageUt.RunningCtestsCount()
	}
	return count
}

func (tracker *CtestsTracker) TestingSummary() TestingSummary {
	return TestingSummary{
		PackagesCount:        tracker.PackagesCount(),
		PassedPackagesCount:  tracker.PassedPackagesCount(),
		FailedPackagesCount:  tracker.FailedPackagesCount(),
		SkippedPackagesCount: tracker.SkippedPackagesCount(),
		RunningPackagesCount: tracker.RunningPackagesCount(),

		TestsCount:        tracker.CtestsCount(),
		PassedTestsCount:  tracker.PassedCtestsCount(),
		FailedTestsCount:  tracker.FailedCtestsCount(),
		SkippedTestsCount: tracker.SkippedCtestsCount(),
		RunningTestsCount: tracker.RunningCtestsCount(),

		DurationS: float32(tracker.testingFinishedAt.Sub(tracker.testingStartedAt).Seconds()),
	}
}

func (tracker *CtestsTracker) replacePackageWith(packageName string, replacement *PackageUnderTest) {
	packageIndex := slices.IndexFunc(tracker.packagesUnderTest, func(packUt *PackageUnderTest) bool {
		return packUt.name == packageName
	})
	if packageIndex == -1 {
		return
	}
	slices.Replace(tracker.packagesUnderTest, packageIndex, packageIndex+1, replacement)
}
