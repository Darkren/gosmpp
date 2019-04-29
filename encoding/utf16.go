package encoding

import "golang.org/x/text/encoding/unicode"

type UTF16 struct {
	Endiannes unicode.Endianness
	BOM       bool
}

func (c UTF16) Encode(str string) ([]byte, error) {
	if c.BOM {
		return encode(str, unicode.UTF16(c.Endiannes, unicode.ExpectBOM).NewEncoder())
	} else {
		return encode(str, unicode.UTF16(c.Endiannes, unicode.UseBOM).NewEncoder())
	}
}

func (c UTF16) Decode(data []byte) ([]byte, error) {
	if c.BOM {
		return decode(data, unicode.UTF16(c.Endiannes, unicode.ExpectBOM).NewDecoder())
	} else {
		return decode(data, unicode.UTF16(c.Endiannes, unicode.UseBOM).NewDecoder())
	}
}
