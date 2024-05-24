package tests_tracker

import (
	"slices"

	"github.com/redjolr/goherent/cmd/events/test_ran_event"
)

type PackageUnderTest struct {
	name  string
	cests [](*Cest)
}

func NewPackageUnderTest(name string) PackageUnderTest {
	return PackageUnderTest{
		name:  name,
		cests: []*Cest{},
	}
}

func (packageUt *PackageUnderTest) insertCestIfNew(cest *Cest) *Cest {
	if !packageUt.HasCest(cest) {
		packageUt.cests = append(packageUt.cests, cest)
	}
	return cest
}

func (packageUt *PackageUnderTest) NewTestRanEvent(evt test_ran_event.TestRanEvent) {
	cest := packageUt.insertCestIfNew(NewCest(evt.Message()))
	cest.NewRanEvent(evt)
}

func (packageUt *PackageUnderTest) HasCest(cest *Cest) bool {
	indexOfCestWithName := slices.IndexFunc(packageUt.cests, func(cest *Cest) bool {
		return cest.HasName(cest.name)
	})
	return indexOfCestWithName != -1
}

func (packageUt *PackageUnderTest) Cest(name string) *Cest {
	if packageUt.HasCest(NewCest(name)) {
		indexOfCestWithName := slices.IndexFunc(packageUt.cests, func(cest *Cest) bool {
			return cest.HasName(name)
		})
		return packageUt.cests[indexOfCestWithName]
	}
	panic("Cest does not exist. Check if it exists, before trying to get it.")
}
