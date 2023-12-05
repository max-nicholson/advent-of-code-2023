package lib

func Min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
