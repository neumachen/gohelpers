package paratils

// SliceContainsStr ...
func SliceContainsStr(strSlice []string, searchString string) bool {
	for _, value := range strSlice {
		if value == searchString {
			return true
		}
	}
	return false
}
