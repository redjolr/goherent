package internal

import (
	"regexp"
	"strconv"
	"strings"
)

type Sequence struct {
	value       string
	isPrintable bool
}

func NewSequence(value string, isPrintable bool) Sequence {
	return Sequence{
		value:       value,
		isPrintable: isPrintable,
	}
}

func (s *Sequence) IsPrintable() bool {
	return s.isPrintable
}

func (s *Sequence) Value() string {
	return s.value
}

func (s *Sequence) MoveLeftCount() int {
	moveCursorLeftRegex, _ := regexp.Compile("\033\\[[0-9]{1,}D")
	moveCursorLeftSeqLoc := moveCursorLeftRegex.FindStringIndex(s.value)
	if moveCursorLeftSeqLoc != nil && moveCursorLeftSeqLoc[0] == 0 {
		moveCursorLeftSeq := s.value[0:moveCursorLeftSeqLoc[1]]

		moveLeftCountAsStr, _ := strings.CutPrefix(moveCursorLeftSeq, "\033[")
		moveLeftCountAsStr, _ = strings.CutSuffix(moveLeftCountAsStr, "D")
		moveLeftCount, err := strconv.Atoi(moveLeftCountAsStr)
		if err != nil {
			panic("Cannot determine the number steps to move left.")
		}
		return moveLeftCount
	}
	return 0
}

func (s *Sequence) MoveRightCount() int {
	moveCursorRightRegex, _ := regexp.Compile("\033\\[[0-9]{1,}C")
	moveCursorRightSeqLoc := moveCursorRightRegex.FindStringIndex(s.value)
	if moveCursorRightSeqLoc != nil && moveCursorRightSeqLoc[0] == 0 {
		moveCursorRightSeq := s.value[0:moveCursorRightSeqLoc[1]]

		moveRightCountAsStr, _ := strings.CutPrefix(moveCursorRightSeq, "\033[")
		moveRightCountAsStr, _ = strings.CutSuffix(moveRightCountAsStr, "C")
		moveRightCount, err := strconv.Atoi(moveRightCountAsStr)
		if err != nil {
			panic("Cannot determine the number steps to move right.")
		}
		return moveRightCount
	}
	return 0
}

func (s *Sequence) MoveUpCount() int {
	moveCursorUpRegex, _ := regexp.Compile("\033\\[[0-9]{1,}A")
	moveCursorUpSeqLoc := moveCursorUpRegex.FindStringIndex(s.value)
	if moveCursorUpSeqLoc != nil && moveCursorUpSeqLoc[0] == 0 {
		moveCursorUpSeq := s.value[0:moveCursorUpSeqLoc[1]]

		moveUpCountAsStr, _ := strings.CutPrefix(moveCursorUpSeq, "\033[")
		moveUpCountAsStr, _ = strings.CutSuffix(moveUpCountAsStr, "A")
		moveUpCount, err := strconv.Atoi(moveUpCountAsStr)
		if err != nil {
			panic("Cannot determine the number steps to move up.")
		}
		return moveUpCount
	}
	return 0
}

func (s *Sequence) MoveDownCount() int {
	moveCursorDownRegex, _ := regexp.Compile("\033\\[[0-9]{1,}B")
	moveCursorDownSeqLoc := moveCursorDownRegex.FindStringIndex(s.value)
	if moveCursorDownSeqLoc != nil && moveCursorDownSeqLoc[0] == 0 {
		moveCursorDownSeq := s.value[0:moveCursorDownSeqLoc[1]]

		moveUpCountAsStr, _ := strings.CutPrefix(moveCursorDownSeq, "\033[")
		moveUpCountAsStr, _ = strings.CutSuffix(moveUpCountAsStr, "B")
		moveDownCount, err := strconv.Atoi(moveUpCountAsStr)
		if err != nil {
			panic("Cannot determine the number steps to move down.")
		}
		return moveDownCount
	}
	return 0
}

func (s *Sequence) Equals(val string) bool {
	return s.value == val
}

func (s *Sequence) Matches(matchRegex string) bool {
	reg, _ := regexp.Compile(matchRegex)
	moveCursorUpSeqLoc := reg.FindStringIndex(s.value)
	if moveCursorUpSeqLoc != nil && moveCursorUpSeqLoc[0] == 0 {
		return true
	}
	return false
}
