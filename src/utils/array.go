package utils

func Contains[T comparable](arr []T, val T) (bool, int) {
	for idx, v := range arr {
		if v == val {
			return true, idx
		}
	}
	return false, -1
}
