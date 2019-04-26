package pdu

// Binarizer converts value to bytes
type Binarizer interface {
	ToBinary() []byte
}
