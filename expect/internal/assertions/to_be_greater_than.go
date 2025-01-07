package assertions

func ToBeGreaterThan(isGreaterCandidate, checkIfGreaterAgainst any) error {
	return compareTwoValues(
		isGreaterCandidate,
		checkIfGreaterAgainst,
		[]compareResult{compareGreater}, "\"%v\" is not greater than \"%v\"",
	)
}
