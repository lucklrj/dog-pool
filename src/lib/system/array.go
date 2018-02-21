package system

func IndexOf(value interface{}, array []interface{}) (isExists bool) {
	for _, singleValue := range array {
		if singleValue == value {
			return true
		}
	}
	return false
}
