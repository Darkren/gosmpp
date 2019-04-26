package pdu

type UnbindResp struct {
	Header
}

func (r *UnbindResp) ToBinary() []byte {
	r.CommandLength = 16

	return r.Header.ToBinary()
}
