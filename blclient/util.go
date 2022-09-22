package blclient

func isOk(data []byte) bool {
	if len(data) >= 2 {
		return (string(data[:2]) == "OK")
	}
	return false
}
