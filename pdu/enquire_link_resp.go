package pdu

type EnquireLinkResp struct {
	Header
}

func (r *EnquireLinkResp) ToBinary() []byte {
	r.CommandLength = 16

	return r.Header.ToBinary()
}
