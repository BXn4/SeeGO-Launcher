package utils

func IsBool(value any) bool {
	switch value {
	case true:
		return true
	case false:
		return true
	}

	return false
}
