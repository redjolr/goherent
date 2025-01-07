package assertions

// ToBeGreaterThan asserts that the first element is greater than the second
//
//	assert.Greater(t, 2, 1)
//	assert.Greater(t, float64(2), float64(1))
//	assert.Greater(t, "b", "a")
func ToBeGreaterThan(isGreaterCandidate, checkIfGreaterAgainst any) error {
	return compareTwoValues(
		isGreaterCandidate,
		checkIfGreaterAgainst,
		[]compareResult{compareGreater}, "\"%v\" is not greater than \"%v\"",
	)
}
