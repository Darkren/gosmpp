package encoding

type ASCII struct {
}

func (c ASCII) Encode(str string) ([]byte, error) {
	return []byte(str), nil
}

func (c ASCII) Decode(data []byte) ([]byte, error) {
	return data, nil
}
