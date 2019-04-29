package parser

import (
	"bytes"
	"errors"

	"github.com/Darkren/gosmpp/converter"
	"github.com/Darkren/gosmpp/pdu"
)

var ErrMailformedPackage = errors.New("the package is mailformed")

var (
	cStringDelim = []byte{0x00}
)

type Parser struct {
	buffer bytes.Buffer
}

func (p *Parser) Append(data []byte) {
	p.buffer.Write(data)
}

func (p *Parser) Next() *pdu.Package {
	if p.buffer.Len() < 16 {
		return nil
	}

	var commandLengthBytes [4]byte
	if _, err := p.buffer.Read(commandLengthBytes[:]); err != nil {
		return nil
	}
	commandLength := converter.BytesToUint32(commandLengthBytes)
	if p.buffer.Len() < int(commandLength)-4 {
		return nil
	}

	pduBytes := make([]byte, commandLength-4)
	if _, err := p.buffer.Read(pduBytes); err != nil {
		return nil
	}

	header := pdu.Header{
		CommandLength: commandLength,
		CommandID: pdu.CommandID(converter.BytesToUint32([4]byte{
			pduBytes[0], pduBytes[1], pduBytes[2], pduBytes[3],
		})),
		CommandStatus: pdu.CommandStatus(converter.BytesToUint32([4]byte{
			pduBytes[4], pduBytes[5], pduBytes[6], pduBytes[7],
		})),
		SequenceNumber: converter.BytesToUint32([4]byte{
			pduBytes[8], pduBytes[9], pduBytes[10], pduBytes[11],
		}),
	}

	body := pduBytes[12:]

	return &pdu.Package{
		Header: header,
		Body:   body,
	}
}

func nextCString(data []byte) (string, int) {
	delimInd := bytes.Index(data, cStringDelim)
	if delimInd == -1 {
		return "", -1
	}

	return string(data[:delimInd]), delimInd + 1
}

func ToBindTransceiver(pack *pdu.Package) (*pdu.BindTransceiver, error) {
	fieldStart := 0

	systemID, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	password, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	systemType, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	interfaceVersion := pack.Body[fieldStart]
	fieldStart++

	addrTON := pdu.TON(pack.Body[fieldStart])
	fieldStart++

	addrNPI := pdu.NPI(pack.Body[fieldStart])
	fieldStart++

	addressRange := string(pack.Body[fieldStart:])

	return &pdu.BindTransceiver{
		Header:           pack.Header,
		SystemID:         systemID,
		Password:         password,
		SystemType:       systemType,
		InterfaceVersion: interfaceVersion,
		AddrTON:          addrTON,
		AddrNPI:          addrNPI,
		AddressRange:     addressRange,
	}, nil
}

func ToBindTransceiverResp(pack *pdu.Package) (*pdu.BindTransceiverResp, error) {
	fieldStart := 0

	systemID, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	tlvs := getTLVs(pack.Body[fieldStart:])

	return &pdu.BindTransceiverResp{
		Header:             pack.Header,
		SystemID:           systemID,
		OptionalParameters: tlvs,
	}, nil
}

func ToDeliverSM(pack *pdu.Package) (*pdu.DeliverSM, error) {
	fieldStart := 0

	serviceType, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	sourceAddrTON := pdu.TON(pack.Body[fieldStart])
	fieldStart++

	sourceAddrNPI := pdu.NPI(pack.Body[fieldStart])
	fieldStart++

	sourceAddr, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	destAddrTON := pdu.TON(pack.Body[fieldStart])
	fieldStart++

	destAddrNPI := pdu.NPI(pack.Body[fieldStart])
	fieldStart++

	destinationAddr, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	esmClass := pack.Body[fieldStart]
	fieldStart++

	protocolID := pack.Body[fieldStart]
	fieldStart++

	priorityFlag := pack.Body[fieldStart]
	fieldStart++

	scheduleDeliveryTime := ""
	fieldStart++

	validityPeriod := ""
	fieldStart++

	registeredDelivery := pack.Body[fieldStart]
	fieldStart++

	var replaceIfPresentFlag uint8 = 0
	fieldStart++

	dataCoding := pdu.DataCoding(pack.Body[fieldStart])
	fieldStart++

	var smDefaultMsgID uint8 = 0
	fieldStart++

	smLength := pack.Body[fieldStart]
	fieldStart++

	shortMessage := string(pack.Body[fieldStart : fieldStart+int(smLength)])
	fieldStart += int(smLength)

	tlvs := getTLVs(pack.Body[fieldStart:])

	return &pdu.DeliverSM{
		Header:               pack.Header,
		ServiceType:          serviceType,
		SourceAddrTON:        sourceAddrTON,
		SourceAddrNPI:        sourceAddrNPI,
		SourceAddr:           sourceAddr,
		DestAddrTON:          destAddrTON,
		DestAddrNPI:          destAddrNPI,
		DestinationAddr:      destinationAddr,
		ESMClass:             esmClass,
		ProtocolId:           protocolID,
		PriorityFlag:         priorityFlag,
		ScheduleDeliveryTime: scheduleDeliveryTime,
		ValidityPeriod:       validityPeriod,
		RegisteredDelivery:   registeredDelivery,
		ReplaceIfPresentFlag: replaceIfPresentFlag,
		DataCoding:           dataCoding,
		SMDefaultMsgId:       smDefaultMsgID,
		SMLength:             smLength,
		ShortMessage:         shortMessage,
		OptionalParameters:   tlvs,
	}, nil
}

func ToDeliverSMResp(pack *pdu.Package) *pdu.DeliverSMResp {
	return &pdu.DeliverSMResp{
		Header:    pack.Header,
		MessageID: "",
	}
}

func ToEnquireLink(pack *pdu.Package) *pdu.EnquireLink {
	return &pdu.EnquireLink{
		Header: pack.Header,
	}
}

func ToEnquireLinkResp(pack *pdu.Package) *pdu.EnquireLinkResp {
	return &pdu.EnquireLinkResp{
		Header: pack.Header,
	}
}

func ToSubmitSM(pack *pdu.Package) (*pdu.SubmitSM, error) {
	fieldStart := 0

	serviceType, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	sourceAddrTON := pdu.TON(pack.Body[fieldStart])
	fieldStart++

	sourceAddrNPI := pdu.NPI(pack.Body[fieldStart])
	fieldStart++

	sourceAddr, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	destAddrTON := pdu.TON(pack.Body[fieldStart])
	fieldStart++

	destAddrNPI := pdu.NPI(pack.Body[fieldStart])
	fieldStart++

	destinationAddr, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	esmClass := pack.Body[fieldStart]
	fieldStart++

	protocolID := pack.Body[fieldStart]
	fieldStart++

	priorityFlag := pack.Body[fieldStart]
	fieldStart++

	scheduleDeliveryTime, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	validityPeriod, fieldStart := nextCString(pack.Body[fieldStart:])
	if fieldStart == -1 {
		return nil, ErrMailformedPackage
	}

	registeredDelivery := pack.Body[fieldStart]
	fieldStart++

	replaceIfPresentFlag := pack.Body[fieldStart]
	fieldStart++

	dataCoding := pdu.DataCoding(pack.Body[fieldStart])
	fieldStart++

	smDefaultMsgID := pack.Body[fieldStart]
	fieldStart++

	smLength := pack.Body[fieldStart]
	fieldStart++

	shortMessage := string(pack.Body[fieldStart : fieldStart+int(smLength)])
	fieldStart += int(smLength)

	tlvs := getTLVs(pack.Body[fieldStart:])

	return &pdu.SubmitSM{
		Header:               pack.Header,
		ServiceType:          serviceType,
		SourceAddrTON:        sourceAddrTON,
		SourceAddrNPI:        sourceAddrNPI,
		SourceAddr:           sourceAddr,
		DestAddrTON:          destAddrTON,
		DestAddrNPI:          destAddrNPI,
		DestinationAddr:      destinationAddr,
		ESMClass:             esmClass,
		ProtocolID:           protocolID,
		PriorityFlag:         priorityFlag,
		ScheduleDeliveryTime: scheduleDeliveryTime,
		ValidityPeriod:       validityPeriod,
		RegisteredDelivery:   registeredDelivery,
		ReplaceIfPresentFlag: replaceIfPresentFlag,
		DataCoding:           dataCoding,
		SMDefaultMsgId:       smDefaultMsgID,
		SMLength:             smLength,
		ShortMessage:         shortMessage,
		OptionalParameters:   tlvs,
	}, nil
}

func ToSubmitSMResp(pack *pdu.Package) *pdu.SubmitSMResp {
	messageId := string(pack.Body[:len(pack.Body)-1])

	return &pdu.SubmitSMResp{
		Header:    pack.Header,
		MessageId: messageId,
	}
}

func ToUnbind(pack *pdu.Package) *pdu.Unbind {
	return &pdu.Unbind{
		Header: pack.Header,
	}
}

func ToUnbindResp(pack *pdu.Package) *pdu.UnbindResp {
	return &pdu.UnbindResp{
		Header: pack.Header,
	}
}

func getTLVs(data []byte) []pdu.TLV {
	var tlvs []pdu.TLV

	uint16Buffer := [2]byte{}
	i := 0
	for i < len(data) {
		// TODO: range check
		uint16Buffer[0] = data[i]
		uint16Buffer[1] = data[i+1]
		tag := pdu.Tag(converter.BytesToUint16(uint16Buffer))

		uint16Buffer[0] = data[i+2]
		uint16Buffer[1] = data[i+3]
		length := converter.BytesToUint16(uint16Buffer)

		value := data[i+4 : i+4+int(length)]

		tlvs = append(tlvs, pdu.TLV{
			Tag:    tag,
			Length: length,
			Value:  value,
		})

		i += 4 + int(length)
	}

	return tlvs
}
