package ctests_tracker

import (
	"errors"
	"slices"
	"strings"

	"github.com/redjolr/goherent/cmd/events"
	"github.com/redjolr/goherent/internal"
)

type outputEventsSlice struct {
	outputEvts []events.CtestOutputEvent
}

func New_outputEventsSlice(outputEvts []events.CtestOutputEvent) outputEventsSlice {
	return outputEventsSlice{
		outputEvts: outputEvts,
	}

}

func (ovs *outputEventsSlice) Contains(str string) bool {
	consecutiveEvtsOutput := ""
	for i := 0; i < len(ovs.outputEvts); i++ {
		consecutiveEvtsOutput += ovs.outputEvts[i].Output
		if strings.Contains(internal.DecodeGoherentTestName(consecutiveEvtsOutput), str) {
			return true
		}
	}
	return false
}

// Both inclusive
func (ovs *outputEventsSlice) CopyOfRange(first, last int) (outputEventsSlice, error) {
	if first > last {
		return outputEventsSlice{}, errors.New("First index has to be less or equal than last index.")
	}
	if first < 0 {
		return outputEventsSlice{}, errors.New("First index has to be greater or equal than 0.")
	}
	if last < 0 {
		return outputEventsSlice{}, errors.New("Last index has to be greater or equal than 0.")
	}
	if first > len(ovs.outputEvts)-1 {
		return outputEventsSlice{}, errors.New("First index has to be less or equal to len of output events -1.")
	}
	if last > len(ovs.outputEvts)-1 {
		return outputEventsSlice{}, errors.New("Last index has to be less or equal to len of output events -1.")
	}

	outputEvts := make([]events.CtestOutputEvent, last-first+1)
	copy(outputEvts, ovs.outputEvts[first:last+1])
	return outputEventsSlice{
		outputEvts: outputEvts,
	}, nil
}

func (ovs *outputEventsSlice) NarrowDownRange(str string, first, last int) (int, int) {
	subSliceWithoutFirstEvt, _ := ovs.CopyOfRange(first+1, last)
	subSliceWithoutLast, _ := ovs.CopyOfRange(first, last-1)
	if !subSliceWithoutLast.Contains(str) && !subSliceWithoutFirstEvt.Contains(str) {
		return first, last
	}
	if !subSliceWithoutLast.Contains(str) && subSliceWithoutFirstEvt.Contains(str) {
		return ovs.NarrowDownRange(str, first+1, last)
	}
	return ovs.NarrowDownRange(str, first, last-1)
}

func (ovs *outputEventsSlice) RemoveOrderRange(first, last int) {
	if first > last {
		return
	}
	if last < len(ovs.outputEvts)-1 {
		ovs.outputEvts = slices.Concat(
			ovs.outputEvts[0:first],
			ovs.outputEvts[last+1:],
		)
	} else if last >= len(ovs.outputEvts)-1 {
		ovs.outputEvts = ovs.outputEvts[0:first]
	}
}

func (ovs *outputEventsSlice) Output() string {
	output := ""
	for _, evt := range ovs.outputEvts {
		output += evt.Output

	}
	return output
}
