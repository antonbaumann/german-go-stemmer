package util

func HasSuffix(a, suff []rune) bool {
	offset := len(a) - len(suff)
	if offset < 0 {
		return false
	}
	for i := 0; i < len(suff); i++ {
		if suff[i] != a[offset+i] {
			return false
		}
	}
	return true
}
