package assertions

func ToBeLessThan(isLessCandidate, checkIfLessAgainst any) error {
	return compareTwoValues(
		isLessCandidate,
		checkIfLessAgainst,
		[]compareResult{compareLess},
		"\"%v\" is not less than \"%v\"",
	)
}
