package encoding

import "golang.org/x/text/encoding/charmap"

type Latin1 struct{}

func (c Latin1) Encode(str string) ([]byte, error) {
	return encode(str, charmap.ISO8859_1.NewEncoder())
}

func (c Latin1) Decode(data []byte) ([]byte, error) {
	return decode(data, charmap.ISO8859_1.NewDecoder())
}
