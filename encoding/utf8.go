package encoding

type UTF8 struct {
}

func (c UTF8) Encode(str string) ([]byte, error) {
	return []byte(str), nil
}

func (c UTF8) Decode(data []byte) ([]byte, error) {
	return data, nil
}
