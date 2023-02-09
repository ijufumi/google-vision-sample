package utils

func MinInArray(array ...float64) float64 {
	value := array[0]
	for _, v := range array {
		if value > v {
			value = v
		}
	}
	return value
}

func MaxInArray(array ...float64) float64 {
	value := array[0]
	for _, v := range array {
		if value < v {
			value = v
		}
	}
	return value
}
