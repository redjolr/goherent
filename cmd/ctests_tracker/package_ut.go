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

func (packageUt *PackageUnderTest) RunningCtestsCount() int {
	count := 0
	for _, ctest := range packageUt.ctests {
		if ctest.isRunning {
			count++
		}
	}
	return count
}

func (packageUt *PackageUnderTest) isCtestTheFirstOne(ctest Ctest) bool {
	if len(packageUt.ctests) == 0 {
		return false
	}
	return packageUt.ctests[0].HasName(ctest.name)
}

func (packageUt *PackageUnderTest) insertCtest(ctest Ctest) Ctest {
	if !packageUt.containsCtest(ctest) {
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

func (packageUt *PackageUnderTest) containsCtest(ctest Ctest) bool {
	indexOfCtestWithName := slices.IndexFunc(packageUt.ctests, func(aCtest Ctest) bool {
		return ctest.Equals(aCtest)
	})
	return indexOfCtestWithName != -1
}
