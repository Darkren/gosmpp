package pdu

type SubmitSMResp struct {
	Header
	MessageId string
}

func (r *SubmitSMResp) ToBinary() []byte {
	messageIDBytes := []byte(r.MessageId)
	messageIDBytes = append(messageIDBytes, 0)

	r.CommandLength = uint32(16 + len(r.MessageId) + 1)

	headerBytes := r.Header.ToBinary()

	data := make([]byte, r.CommandLength)

	bp := copy(data, headerBytes)

	bp += copy(data[bp:], r.MessageId)
	data[bp] = 0

	return data
}
