package ctests_tracker

import (
	"slices"
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

func (packageUt *PackageUnderTest) insertCtest(ctest Ctest) Ctest {
	if !packageUt.containsCtestWithName(ctest.name) {
		packageUt.ctests = append(packageUt.ctests, ctest)
		return ctest
	}
	panic("Ctest already exists")
}

func (packageUt *PackageUnderTest) containsCtestWithName(ctestName string) bool {
	indexOfCtestWithName := slices.IndexFunc(packageUt.ctests, func(ctest Ctest) bool {
		return ctest.HasName(ctestName)
	})
	return indexOfCtestWithName != -1
}
