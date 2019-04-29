package encoding

import "golang.org/x/text/encoding"

type Encoding interface {
	Encode(str string) ([]byte, error)
	Decode([]byte) (string, error)
}

func encode(str string, encoder *encoding.Encoder) ([]byte, error) {
	return encoder.Bytes([]byte(str))
}

func decode(data []byte, decoder *encoding.Decoder) ([]byte, error) {
	tmp, err := decoder.Bytes(data)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}
