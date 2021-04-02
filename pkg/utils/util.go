package utils

func InArray(val string, slice []string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}

	return -1, false
}

func Uint8ToString(bs ...uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = v
	}

	return string(b)
}

func Int8ToBoolean(value int8) bool {
	if value > 0 {
		return true
	} else {
		return false
	}
}
