package pdu

type Unbind struct {
	Header
}

func (u *Unbind) ToBinary() []byte {
	u.CommandLength = 16

	return u.Header.ToBinary()
}
