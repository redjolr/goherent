package ctests_tracker

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/redjolr/goherent/cmd/events"
)

type orderOutputEvt struct {
	originalOrder int
	event         events.CtestOutputEvent
}

type outputEventsSlice struct {
	orderedOutputEvts []orderOutputEvt
}

func New_outputEventsSlice(outputEvts []events.CtestOutputEvent) outputEventsSlice {
	slice := outputEventsSlice{
		orderedOutputEvts: []orderOutputEvt{},
	}
	for i, evt := range outputEvts {
		slice.orderedOutputEvts = append(slice.orderedOutputEvts, orderOutputEvt{
			originalOrder: i,
			event:         evt,
		})
	}
	return slice
}

func (ovs *outputEventsSlice) Contains(str string) bool {
	consecutiveEvtsOutput := ""
	for i := 0; i < len(ovs.orderedOutputEvts); i++ {
		fmt.Println("\nOUTPUT: ", ovs.orderedOutputEvts[i].event.Output)
		consecutiveEvtsOutput += ovs.orderedOutputEvts[i].event.Output
		if strings.Contains(consecutiveEvtsOutput, str) {
			return true
		}
	}
	return false
}

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
	if first > len(ovs.orderedOutputEvts)-1 {
		return outputEventsSlice{}, errors.New("First index has to be less or equal to len of output events -1.")
	}
	if last > len(ovs.orderedOutputEvts)-1 {
		return outputEventsSlice{}, errors.New("Last index has to be less or equal to len of output events -1.")
	}

	orderedOutputEvts := make([]orderOutputEvt, last-first)
	copy(orderedOutputEvts, ovs.orderedOutputEvts[first:last+1])
	fmt.Println("\n\n COPY", orderedOutputEvts)
	return outputEventsSlice{
		orderedOutputEvts: orderedOutputEvts,
	}, nil
}

func (ovs *outputEventsSlice) NarrowDownRangeStartingFromBeginning(str string, first, last int) (int, int) {
	if str == "" || len(ovs.orderedOutputEvts) == 0 {
		return -1, -1
	}
	if !ovs.Contains(str) {
		return -1, -1
	}

	subSliceWithoutFirst, _ := ovs.CopyOfRange(first+1, last)
	if !subSliceWithoutFirst.Contains(str) {
		return ovs.orderedOutputEvts[first].originalOrder, ovs.orderedOutputEvts[last].originalOrder
	}
	return subSliceWithoutFirst.NarrowDownRangeStartingFromBeginning(str, first+1, last)

	// subSliceWithoutLast, _ := ovs.CopyOfRange(first, last-1)
	// if !subSliceWithoutLast.Contains(str) {
	// 	// ovs.orderedOutputEvts = ovs.orderedOutputEvts[0 : len(ovs.orderedOutputEvts)-1]
	// 	// return ovs.FindRangeWithinSubslice(str, first, last-1)
	// 	return ovs.orderedOutputEvts[first].originalOrder, ovs.orderedOutputEvts[last].originalOrder

	// }

	// for i := first; i < last; i++ {
	// 	for j := i + 1; j < last; j++ {
	// 		if ovs.Contains(str) {
	// 			return ovs.orderedOutputEvts[i].originalOrder, ovs.orderedOutputEvts[j].originalOrder
	// 		}
	// 	}
	// }
	// return -1, -1
}

func (ovs *outputEventsSlice) NarrowDownRangeStartingFromEnd(str string, first, last int) (int, int) {
	if str == "" || len(ovs.orderedOutputEvts) == 0 {
		return -1, -1
	}
	if !ovs.Contains(str) {
		return -1, -1
	}

	subSliceWithoutLast, _ := ovs.CopyOfRange(first, last-1)
	if !subSliceWithoutLast.Contains(str) {
		return ovs.orderedOutputEvts[first].originalOrder, ovs.orderedOutputEvts[last].originalOrder

	}
	return subSliceWithoutLast.NarrowDownRangeStartingFromEnd(str, first, last-1)
}

func (ovs *outputEventsSlice) RemoveOriginalOrderRange(first, last int) {
	fmt.Println("\n\n LENGTH", len(ovs.orderedOutputEvts))
	for i, evt := range ovs.orderedOutputEvts {
		if evt.originalOrder >= first && evt.originalOrder <= last {
			if i < len(ovs.orderedOutputEvts)-1 {
				ovs.orderedOutputEvts = slices.Concat(
					ovs.orderedOutputEvts[0:i],
					ovs.orderedOutputEvts[i+1:],
				)
			} else if i == len(ovs.orderedOutputEvts)-1 {
				ovs.orderedOutputEvts = ovs.orderedOutputEvts[0:i]
			}
		}
	}
}

func (ovs *outputEventsSlice) Output() string {
	output := ""
	for i, evt := range ovs.orderedOutputEvts {
		output += evt.event.Output
		if i < len(ovs.orderedOutputEvts)-1 {
			output += "\n"
		}
	}
	return output
}

// func (ovs *outputEventsSlice) FindRangeFromEnd(str string) (int, int) {
// 	if str == "" || len(ovs.orderedOutputEvts) == 0 {
// 		return -1, -1
// 	}
// 	if !ovs.Contains(str) {
// 		return -1, -1
// 	}

// 	subSliceWithoutFirst, _ := ovs.CopyOfRange(1, len(ovs.orderedOutputEvts))
// 	if subSliceWithoutFirst.Contains(str) {
// 		ovs.orderedOutputEvts = ovs.orderedOutputEvts[1:]
// 		return ovs.FindRangeFromBeginning(str)
// 	}

// 	subSliceWithoutLast, _ := ovs.CopyOfRange(0, len(ovs.orderedOutputEvts)-1)
// 	if subSliceWithoutLast.Contains(str) {
// 		ovs.orderedOutputEvts = ovs.orderedOutputEvts[0 : len(ovs.orderedOutputEvts)-1]
// 		return ovs.FindRangeFromBeginning(str)
// 	}

// 	for i := 0; i < len(ovs.orderedOutputEvts); i++ {
// 		consecutiveEvtsOutput := ""
// 		for j := i + 1; j < len(ovs.orderedOutputEvts); j++ {
// 			consecutiveEvtsOutput += ovs.orderedOutputEvts[j].event.Output
// 			if strings.Contains(consecutiveEvtsOutput, str) {
// 				return i, j
// 			}
// 		}
// 	}
// 	return -1, -1
// }
