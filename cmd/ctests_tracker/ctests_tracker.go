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

func (tracker *CtestsTracker) insertPackageUnderTestIfNew(packUt PackageUnderTest) {
	if !tracker.ContainsPackageUtWithName(packUt.name) {
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, packUt)
	}
}

func (tracker *CtestsTracker) InsertCtest(ctest Ctest) Ctest {
	packUt := NewPackageUnderTest(ctest.packageName)
	packUt.insertCtest(ctest)
	tracker.packagesUnderTest = append(tracker.packagesUnderTest, packUt)
	return ctest
}

func (tracker *CtestsTracker) NewCtestRanEvent(evt ctest_ran_event.CtestRanEvent) {
	if !tracker.ContainsPackageUtWithName(evt.PackageName()) {
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, NewPackageUnderTest(evt.PackageName()))
	}
}

func (tracker *CtestsTracker) ContainsPackageUtWithName(name string) bool {
	indexOfPackUttWithName := slices.IndexFunc(tracker.packagesUnderTest, func(packUt PackageUnderTest) bool {
		return packUt.name == name
	})
	return indexOfPackUttWithName != -1
}

func (tracker *CtestsTracker) PackageUnderTest(name string) *PackageUnderTest {
	if tracker.ContainsPackageUtWithName(name) {
		indexOfPackUtWithName := slices.IndexFunc(tracker.packagesUnderTest, func(packUt PackageUnderTest) bool {
			return packUt.name == name
		})
		return &tracker.packagesUnderTest[indexOfPackUtWithName]
	}
	panic("Ctest does not exist. Check if it exists, before trying to get it.")
}

func (tracker *CtestsTracker) ContainsCtestWithName(name string) bool {
	indexOfPackUttWithName := slices.IndexFunc(tracker.packagesUnderTest, func(packUt PackageUnderTest) bool {
		return packUt.containsCtestWithName(name)
	})
	return indexOfPackUttWithName != -1
}
