package tests_tracker

import (
	"slices"

	"github.com/redjolr/goherent/cmd/events/test_paused_event"
	"github.com/redjolr/goherent/cmd/events/test_ran_event"
)

type PackageUnderTest struct {
	name   string
	ctests []Ctest
}

func NewPackageUnderTest(name string) PackageUnderTest {
	return PackageUnderTest{
		name:   name,
		ctests: []Ctest{},
	}
}

func (packageUt *PackageUnderTest) insertCtestIfNew(ctest Ctest) *Ctest {
	if !packageUt.HasCtest(ctest) {
		packageUt.ctests = append(packageUt.ctests, ctest)
	}
	return packageUt.Ctest(ctest.name)
}

func (packageUt *PackageUnderTest) NewTestRanEvent(evt test_ran_event.TestRanEvent) {
	ctest := packageUt.insertCtestIfNew(NewCtest(evt.Message()))
	ctest.NewRanEvent(evt)
}

func (packageUt *PackageUnderTest) NewTestPausedEvent(evt test_paused_event.TestPausedEvent) {
	ctest := packageUt.insertCtestIfNew(NewCtest(evt.Message()))
	ctest.NewPausedEvent(evt)
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
