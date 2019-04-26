package pdu

type DeliverSMResp struct {
	Header
	MessageID string
}

func (d *DeliverSMResp) ToBinary() []byte {
	d.CommandLength = uint32(16 + len(d.MessageID) + 1)

	headerBytes := d.Header.ToBinary()

	data := make([]byte, d.CommandLength)

	bp := copy(data, headerBytes)

	bp += copy(data, d.MessageID)
	data[bp] = 0

	return data
}
