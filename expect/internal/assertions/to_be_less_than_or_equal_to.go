package assertions

func ToBeLessThanOrEqualTo(isLessOrEqualCandidate, checkIfLessOrEqualAgainst any) error {
	return compareTwoValues(
		isLessOrEqualCandidate,
		checkIfLessOrEqualAgainst,
		[]compareResult{compareLess, compareEqual},
		"\"%v\" is not less than or equal to \"%v\"",
	)
}
