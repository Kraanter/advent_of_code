package util

func AbsoluteInt(number int) int {
	if number < 0 {
		return -number
	}

	return number
}
