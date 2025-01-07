package assertions

func ToBeGreaterThanOrEqualTo(isGreaterOrEqualCandidate, checkIfGreaterOrEqualAgainst any) error {
	return compareTwoValues(
		isGreaterOrEqualCandidate,
		checkIfGreaterOrEqualAgainst,
		[]compareResult{compareGreater, compareEqual}, "\"%v\" is not greater than or equal to \"%v\"",
	)
}
