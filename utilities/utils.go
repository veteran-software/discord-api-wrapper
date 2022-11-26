package utilities

func Contains[T comparable](slice []T, e T) bool {
	for _, v := range slice {
		if v == e {
			return true
		}
	}

	return false
}
