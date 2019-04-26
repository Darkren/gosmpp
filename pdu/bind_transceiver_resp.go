package pdu

type BindTransceiverResp struct {
	Header
	SystemID           string
	OptionalParameters []TLV
}

func (r *BindTransceiverResp) ToBinary() []byte {
	tlvsLen := 0
	tlvsBytes := make([][]byte, 0, len(r.OptionalParameters))
	for i, tlv := range r.OptionalParameters {
		tlvsBytes = append(tlvsBytes, tlv.ToBinary())
		tlvsLen += len(tlvsBytes[i])
	}

	r.CommandLength = uint32(16 + len(r.SystemID) + 1 + tlvsLen)

	headerBytes := r.Header.ToBinary()

	data := make([]byte, r.CommandLength)

	bp := copy(data, headerBytes)

	bp += copy(data[bp:], r.SystemID)
	data[bp] = 0
	bp++

	for _, tlv := range tlvsBytes {
		bp += copy(data[bp:], tlv)
	}

	return data
}
