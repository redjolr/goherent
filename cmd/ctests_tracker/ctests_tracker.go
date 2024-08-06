package ctests_tracker

import (
	"fmt"
	"slices"

	"github.com/redjolr/goherent/cmd/events"
)

type CtestsTracker struct {
	packagesUnderTest []*PackageUnderTest
}

func NewCtestsTracker() CtestsTracker {
	return CtestsTracker{
		packagesUnderTest: []*PackageUnderTest{},
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

func (tracker *CtestsTracker) NewCtestRanEvent(evt events.CtestRanEvent) {
	if !tracker.ContainsPackageUtWithName(evt.PackageName) {
		packUt := NewPackageUnderTest(evt.PackageName)
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, &packUt)
	}
	fmt.Println()
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
	fmt.Println("\n\n\n indexOfPackUttWithName", indexOfPackUttWithName)
	return indexOfPackUttWithName != -1
}

func (tracker *CtestsTracker) MarkAllPackagesAsFinished() {
	for _, packageUt := range tracker.packagesUnderTest {
		packageUt.MarkAsFinished()
	}
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

func (tracker *CtestsTracker) insertPackageUnderTestIfNew(packUt PackageUnderTest) {
	if !tracker.ContainsPackageUtWithName(packUt.name) {
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, &packUt)
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
