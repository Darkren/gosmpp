package pdu

type EnquireLink struct {
	Header
}

func (l *EnquireLink) ToBinary() []byte {
	l.CommandLength = 16

	return l.Header.ToBinary()
}
