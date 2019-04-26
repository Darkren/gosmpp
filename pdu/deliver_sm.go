package pdu

type DeliverSM struct {
	Header
	ServiceType          string
	SourceAddrTON        TON
	SourceAddrNPI        NPI
	SourceAddr           string
	DestAddrTON          TON
	DestAddrNPI          NPI
	DestinationAddr      string
	ESMClass             uint8
	ProtocolId           uint8
	PriorityFlag         uint8
	ScheduleDeliveryTime string
	ValidityPeriod       string
	RegisteredDelivery   uint8
	ReplaceIfPresentFlag uint8
	DataCoding           DataCoding
	SMDefaultMsgId       uint8
	SMLength             uint8
	ShortMessage         string
	OptionalParameters   []TLV
}

func (d *DeliverSM) ToBinary() []byte {
	tlvsLen := 0
	tlvsBytes := make([][]byte, 0, len(d.OptionalParameters))
	for i, tlv := range d.OptionalParameters {
		tlvsBytes = append(tlvsBytes, tlv.ToBinary())
		tlvsLen += len(tlvsBytes[i])
	}

	d.CommandLength = uint32(16 + len(d.ServiceType) + 1 + 1 + 1 + len(d.SourceAddr) + 1 + 1 + 1 +
		len(d.DestinationAddr) + 1 + 1 + 1 + 1 + len(d.ScheduleDeliveryTime) + 1 + len(d.ValidityPeriod) + 1 +
		1 + 1 + 1 + 1 + 1 + len(d.ShortMessage) + tlvsLen)

	headerBytes := d.Header.ToBinary()

	data := make([]byte, d.CommandLength)

	bp := copy(data, headerBytes)

	bp += copy(data[bp:], d.ServiceType)
	data[bp] = 0
	bp++

	data[bp] = byte(d.SourceAddrTON)
	bp++

	data[bp] = byte(d.SourceAddrNPI)
	bp++

	bp += copy(data[bp:], d.SourceAddr)
	data[bp] = 0
	bp++

	data[bp] = byte(d.DestAddrTON)
	bp++

	data[bp] = byte(d.DestAddrNPI)
	bp++

	bp += copy(data[bp:], d.DestinationAddr)
	data[bp] = 0
	bp++

	data[bp] = d.ESMClass
	bp++

	data[bp] = d.ProtocolId
	bp++

	data[bp] = d.PriorityFlag
	bp++

	bp += copy(data[bp:], d.ScheduleDeliveryTime)
	data[bp] = 0
	bp++

	bp += copy(data[bp:], d.ValidityPeriod)
	data[bp] = 0
	bp++

	data[bp] = d.RegisteredDelivery
	bp++

	data[bp] = d.ReplaceIfPresentFlag
	bp++

	data[bp] = byte(d.DataCoding)
	bp++

	data[bp] = d.SMDefaultMsgId
	bp++

	data[bp] = d.SMLength
	bp++

	bp += copy(data[bp:], d.ShortMessage)

	for _, tlv := range tlvsBytes {
		bp += copy(data[bp:], tlv)
	}

	return data
}
