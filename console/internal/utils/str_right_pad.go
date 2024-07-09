package utils

func StrRightPad(s string, padStr string, totalLength int) string {
	for len(s) < totalLength {
		s += padStr
	}
	return s
}
