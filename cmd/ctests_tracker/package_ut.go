package ctests_tracker

import (
	"slices"

	"github.com/redjolr/goherent/cmd/events/ctest_failed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_passed_event"
	"github.com/redjolr/goherent/cmd/events/ctest_paused_event"
	"github.com/redjolr/goherent/cmd/events/ctest_ran_event"
)

type PackageUnderTest struct {
	name   string
	ctests []Ctest
}

func NewPackageUnderTest(name string) PackageUnderTest {
	newPack := PackageUnderTest{
		name:   name,
		ctests: []Ctest{},
	}
	return newPack
}

func (packageUt *PackageUnderTest) insertCtestIfNew(ctest Ctest) *Ctest {
	if !packageUt.HasCtest(ctest) {
		packageUt.ctests = append(packageUt.ctests, ctest)
	}
	return packageUt.Ctest(ctest.name)
}

func (packageUt *PackageUnderTest) NewTestRanEvent(evt ctest_ran_event.CtestRanEvent) {
	ctest := packageUt.insertCtestIfNew(NewCtest(evt.Message()))
	ctest.NewRanEvent(evt)
}

func (packageUt *PackageUnderTest) NewTestPausedEvent(evt ctest_paused_event.CtestPausedEvent) {
	ctest := packageUt.insertCtestIfNew(NewCtest(evt.Message()))
	ctest.NewPausedEvent(evt)
}

func (packageUt *PackageUnderTest) NewTestPassedEvent(evt ctest_passed_event.CtestPassedEvent) {
	ctest := packageUt.insertCtestIfNew(NewCtest(evt.Message()))
	ctest.NewPassedEvent(evt)
}

func (packageUt *PackageUnderTest) NewTestFailedEvent(evt ctest_failed_event.CtestFailedEvent) {
	ctest := packageUt.insertCtestIfNew(NewCtest(evt.Message()))
	ctest.NewFailedEvent(evt)
}

func (packageUt *PackageUnderTest) HasCtest(ctest Ctest) bool {
	indexOfCtestWithName := slices.IndexFunc(packageUt.ctests, func(ctest Ctest) bool {
		return ctest.HasName(ctest.name)
	})
	return indexOfCtestWithName != -1
}

func (packageUt *PackageUnderTest) Ctest(name string) *Ctest {
	if packageUt.HasCtest(NewCtest(name)) {
		indexOfCtestWithName := slices.IndexFunc(packageUt.ctests, func(ctest Ctest) bool {
			return ctest.HasName(name)
		})
		return &packageUt.ctests[indexOfCtestWithName]
	}
	panic("Ctest does not exist. Check if it exists, before trying to get it.")
}

func (packageUt *PackageUnderTest) TestCount() int {
	return len(packageUt.ctests)
}

func (packageUt *PackageUnderTest) HasName(name string) bool {
	return packageUt.name == name
}
