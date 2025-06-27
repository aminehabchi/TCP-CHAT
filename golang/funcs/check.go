package funcs

func IsValidName(name string) bool {
	if len(name) < 2 || len(name) > 15 {
		return false
	}
	for i := 0; i < len(name); i++ {
		if name[i] < 33 {
			return false
		}
	}
	return true
}

func isValidMessage(message string) bool {
	message = message[:len(message)-1]
	// fmt.Println([]byte(message))
	if len(message) == 0 {
		return false
	}
	for i := 0; i < len(message); i++ {
		if message[i] < 32 || message[i] > 128 {
			return false
		}
	}
	return true
}
