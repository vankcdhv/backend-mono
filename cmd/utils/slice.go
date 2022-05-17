package utils

func StringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Int64SliceContains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Int32SliceContains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IntSliceContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
