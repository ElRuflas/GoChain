package utils

func CheckHash(hash []byte, difficulty int) bool {
	for i := 0; i < difficulty; i++ {
		if hash[i] != 0x00 {
			return false
		}
	}
	return true
}
