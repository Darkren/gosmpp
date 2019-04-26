package pdu

import (
	"bitbucket.org/Darkren/smpp-client/smpp/converter"
)

// Header is a header of any valid SMPP package
type Header struct {
	CommandLength  uint32
	CommandID      CommandID
	CommandStatus  CommandStatus
	SequenceNumber uint32
}

func (h Header) ToBinary() []byte {
	commandLengthBytes := converter.Uint32ToBytes(h.CommandLength)
	commandIDBytes := converter.Uint32ToBytes(uint32(h.CommandID))
	commandStatusBytes := converter.Uint32ToBytes(uint32(h.CommandStatus))
	sequenceNumberBytes := converter.Uint32ToBytes(h.SequenceNumber)

	return []byte{
		commandLengthBytes[0],
		commandLengthBytes[1],
		commandLengthBytes[2],
		commandLengthBytes[3],
		commandIDBytes[0],
		commandIDBytes[1],
		commandIDBytes[2],
		commandIDBytes[3],
		commandStatusBytes[0],
		commandStatusBytes[1],
		commandStatusBytes[2],
		commandStatusBytes[3],
		sequenceNumberBytes[0],
		sequenceNumberBytes[1],
		sequenceNumberBytes[2],
		sequenceNumberBytes[3],
	}
}
