package hyperdeck

func IsAsync(payload []byte) bool {
	return payload[0] == '5'
}

func IsSuccess(payload []byte) bool {
	return payload[0] == '2'
}

func IsError(payload []byte) bool {
	return !IsSuccess(payload) && !IsAsync(payload)
}
