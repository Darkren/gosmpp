package pdu

import "github.com/Darkren/gosmpp/converter"

type TLV struct {
	Tag    Tag
	Length uint16
	Value  []byte
}

func (t TLV) ToBinary() []byte {
	tagBytes := converter.Uint16ToBytes(uint16(t.Tag))
	lengthBytes := converter.Uint16ToBytes(t.Length)

	data := make([]byte, 4+t.Length)

	data[0] = tagBytes[0]
	data[1] = tagBytes[1]
	data[2] = lengthBytes[0]
	data[3] = lengthBytes[1]

	copy(data[4:], t.Value)

	return data
}
