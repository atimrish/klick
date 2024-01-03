package helpers

func ArrayReverse[T any](array *[]T) []T {

	length := len(*array)

	tmpArray := make([]T, length)

	for i := 0; i < length; i++ {
		tmpArray[i] = (*array)[length-i-1]
	}

	return tmpArray
}

func ArrayContains[T string](array *[]T, value T) bool {
	length := len(*array)

	for i := 0; i < length; i++ {
		if (*array)[i] == value {
			return true
		}
	}

	return false
}

