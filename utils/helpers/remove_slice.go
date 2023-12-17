package helpers

func RemoveValue(slice []int, value int) []int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
