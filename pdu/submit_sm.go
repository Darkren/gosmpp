package pdu

type SubmitSM struct {
	Header
	ServiceType          string
	SourceAddrTON        TON
	SourceAddrNPI        NPI
	SourceAddr           string
	DestAddrTON          TON
	DestAddrNPI          NPI
	DestinationAddr      string
	ESMClass             uint8
	ProtocolID           uint8
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

func (s *SubmitSM) ToBinary() []byte {
	tlvsLen := 0
	tlvsBytes := make([][]byte, 0, len(s.OptionalParameters))
	for i, tlv := range s.OptionalParameters {
		tlvsBytes = append(tlvsBytes, tlv.ToBinary())
		tlvsLen += len(tlvsBytes[i])
	}

	s.CommandLength = uint32(16 + len(s.ServiceType) + 1 + 1 + 1 + len(s.SourceAddr) + 1 + 1 + 1 +
		len(s.DestinationAddr) + 1 + 1 + 1 + 1 + len(s.ScheduleDeliveryTime) + 1 + len(s.ValidityPeriod) + 1 + 1 +
		1 + 1 + 1 + 1 + len(s.ShortMessage) + tlvsLen)

	headerBytes := s.Header.ToBinary()

	data := make([]byte, s.CommandLength)

	bp := copy(data, headerBytes)

	bp += copy(data[bp:], s.ServiceType)
	data[bp] = 0
	bp++

	data[bp] = byte(s.SourceAddrTON)
	bp++

	data[bp] = byte(s.SourceAddrNPI)
	bp++

	bp += copy(data[bp:], s.SourceAddr)
	data[bp] = 0
	bp++

	data[bp] = byte(s.DestAddrTON)
	bp++

	data[bp] = byte(s.DestAddrNPI)
	bp++

	bp += copy(data[bp:], s.DestinationAddr)
	data[bp] = 0
	bp++

	data[bp] = s.ESMClass
	bp++

	data[bp] = s.ProtocolID
	bp++

	data[bp] = s.PriorityFlag
	bp++

	bp += copy(data[bp:], s.ScheduleDeliveryTime)
	data[bp] = 0
	bp++

	bp += copy(data[bp:], s.ValidityPeriod)
	data[bp] = 0
	bp++

	data[bp] = s.RegisteredDelivery
	bp++

	data[bp] = s.ReplaceIfPresentFlag
	bp++

	data[bp] = byte(s.DataCoding)
	bp++

	data[bp] = s.SMDefaultMsgId
	bp++

	data[bp] = s.SMLength
	bp++

	bp += copy(data[bp:], s.ShortMessage)

	for _, tlv := range tlvsBytes {
		bp += copy(data[bp:], tlv)
	}

	return data
}
