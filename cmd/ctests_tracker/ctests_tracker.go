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

func (tracker *CtestsTracker) NewCtestRanEvent(evt ctest_ran_event.CtestRanEvent) {
	if !tracker.ContainsPackageUtWithName(evt.PackageName()) {
		tracker.packagesUnderTest = append(tracker.packagesUnderTest, NewPackageUnderTest(evt.PackageName()))
	}
}

func (tracker *CtestsTracker) ContainsPackageUtWithName(name string) bool {
	indexOfPackUttWithName := slices.IndexFunc(tracker.packagesUnderTest, func(packUt PackageUnderTest) bool {
		return packUt.HasName(name)
	})
	return indexOfPackUttWithName != -1
}

func (tracker *CtestsTracker) PackageUnderTest(name string) *PackageUnderTest {
	if tracker.ContainsPackageUtWithName(name) {
		indexOfPackUtWithName := slices.IndexFunc(tracker.packagesUnderTest, func(packUt PackageUnderTest) bool {
			return packUt.HasName(name)
		})
		return &tracker.packagesUnderTest[indexOfPackUtWithName]
	}
	panic("Ctest does not exist. Check if it exists, before trying to get it.")
}
