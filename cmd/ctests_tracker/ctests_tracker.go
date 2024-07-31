package ctests_tracker

import (
	"slices"

	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
)

type CtestsTracker struct {
	packagesUnderTest []PackageUnderTest
}

func NewCtestsTracker() CtestsTracker {
	return CtestsTracker{
		packagesUnderTest: []PackageUnderTest{},
	}
}

func (tracker *CtestsTracker) InsertCtest(ctest Ctest) Ctest {
	var packUt PackageUnderTest
	if tracker.ContainsPackageUtWithName(ctest.packageName) {
		packUt = tracker.PackageUnderTest(ctest.packageName)
		packUt.insertCtest(ctest)
		tracker.replacePackageWith(ctest.packageName, packUt)
	} else {
		packUt = NewPackageUnderTest(ctest.packageName)
		packUt.insertCtest(ctest)
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, packUt)
	}

	return ctest
}

func (tracker *CtestsTracker) NewCtestRanEvent(evt ctest_ran_event.CtestRanEvent) {
	if !tracker.ContainsPackageUtWithName(evt.PackageName()) {
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, NewPackageUnderTest(evt.PackageName()))
	}
}

func (tracker *CtestsTracker) IsCtestFirstOfItsPackage(ctest Ctest) bool {
	if !tracker.ContainsPackageUtWithName(ctest.packageName) {
		return false
	}
	packageUnderTest := tracker.PackageUnderTest(ctest.packageName)
	return packageUnderTest.isCtestTheFirstOne(ctest)
}

func (tracker *CtestsTracker) ContainsPackageUtWithName(name string) bool {
	indexOfPackUttWithName := slices.IndexFunc(tracker.packagesUnderTest, func(packUt PackageUnderTest) bool {
		return packUt.name == name
	})
	return indexOfPackUttWithName != -1
}

func (tracker *CtestsTracker) PackagesCount() int {
	return len(tracker.packagesUnderTest)
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
		if packageUt.HasFailedTests() {
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

func (tracker *CtestsTracker) PackageUnderTest(name string) PackageUnderTest {
	if tracker.ContainsPackageUtWithName(name) {
		indexOfPackUtWithName := slices.IndexFunc(tracker.packagesUnderTest, func(packUt PackageUnderTest) bool {
			return packUt.name == name
		})
		return tracker.packagesUnderTest[indexOfPackUtWithName]
	}
	panic("Ctest does not exist. Check if it exists, before trying to get it.")
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

func (tracker *CtestsTracker) insertPackageUnderTestIfNew(packUt PackageUnderTest) {
	if !tracker.ContainsPackageUtWithName(packUt.name) {
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, packUt)
	}
}

func (tracker *CtestsTracker) replacePackageWith(packageName string, replacement PackageUnderTest) {
	packageIndex := slices.IndexFunc(tracker.packagesUnderTest, func(packUt PackageUnderTest) bool {
		return packUt.name == packageName
	})
	if packageIndex == -1 {
		return
	}
	slices.Replace(tracker.packagesUnderTest, packageIndex, packageIndex+1, replacement)
}
