package encoding

import "golang.org/x/text/encoding/charmap"

type CP1252 struct{}

func (c CP1252) Encode(str string) ([]byte, error) {
	return encode(str, charmap.Windows1252.NewEncoder())
}

func (c CP1252) Decode(data []byte) ([]byte, error) {
	return decode(data, charmap.Windows1252.NewDecoder())
}
