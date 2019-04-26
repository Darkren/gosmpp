package pdu

type BindTransceiver struct {
	Header
	SystemID         string
	Password         string
	SystemType       string
	InterfaceVersion uint8
	AddrTON          TON
	AddrNPI          NPI
	AddressRange     string
}

func (b *BindTransceiver) ToBinary() []byte {
	b.CommandLength = uint32(16 + len(b.SystemID) + 1 + len(b.Password) + 1 + len(b.SystemType) + 1 +
		1 + 1 + 1 + len(b.AddressRange) + 1)

	headerBytes := b.Header.ToBinary()

	data := make([]byte, b.CommandLength)

	bp := copy(data, headerBytes)

	bp += copy(data[bp:], b.SystemID)
	data[bp] = 0
	bp++

	bp += copy(data[bp:], b.Password)
	data[bp] = 0
	bp++
	bp += copy(data[bp:], b.SystemType)
	data[bp] = 0
	bp++

	data[bp] = b.InterfaceVersion
	bp++

	data[bp] = byte(b.AddrTON)
	bp++

	data[bp] = byte(b.AddrNPI)
	bp++

	bp += copy(data[bp:], b.AddressRange)
	data[bp] = 0

	return data
}
