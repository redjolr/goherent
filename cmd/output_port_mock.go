package cmd

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type OutputPortMock struct {
	mock.Mock
}

func NewOutputPortMock() *OutputPortMock {
	return &OutputPortMock{}
}

func (outputMock *OutputPortMock) FirstCtestOfPackageStartedRunning(testName string, packageName string) {
	outputMock.Called(testName, packageName)

}

func (outputMock *OutputPortMock) FirstCtestOfPackagePassed(testName string, packageName string, testDuration float64) {
	outputMock.Called(testName, packageName, testDuration)
}

func (outputMock *OutputPortMock) CtestPassed(testName string, timeElapsed float64) {
	outputMock.Called(testName, timeElapsed)
}

func (outputMock *OutputPortMock) CtestStartedRunning(testName string) {
	outputMock.Called(testName)
}

func (outputPort *OutputPortMock) CtestFailed(testName string, timeElapsed float64) {
	fmt.Printf("❌ %s\n\n %f\n", testName, timeElapsed)
}
