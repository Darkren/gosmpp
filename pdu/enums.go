package pdu

import (
	"strconv"
)

type CommandID uint32

const (
	CommandSubmitSm            CommandID = 0x00000004
	CommandSubmitSmResp        CommandID = 0x80000004
	CommandDeliverSm           CommandID = 0x00000005
	CommandDeliverSmResp       CommandID = 0x80000005
	CommandUnbind              CommandID = 0x00000006
	CommandUnbindResp          CommandID = 0x80000006
	CommandBindTransceiver     CommandID = 0x00000009
	CommandBindTransceiverResp CommandID = 0x80000009
	CommandEnquireLink         CommandID = 0x00000015
	CommandEnquireLinkResp     CommandID = 0x80000015
)

type CommandStatus uint32

const (
	CommandStatus_ESME_ROCK             CommandStatus = 0x00000000
	CommandStatus_ESME_RINVMSGLEN       CommandStatus = 0x00000001
	CommandStatus_ESME_RINVCMDLEN       CommandStatus = 0x00000002
	CommandStatus_ESME_RINVCMDID        CommandStatus = 0x00000003
	CommandStatus_ESME_RINVBNDSTS       CommandStatus = 0x00000004
	CommandStatus_ESME_RALYBND          CommandStatus = 0x00000005
	CommandStatus_ESME_RINVPRTFLG       CommandStatus = 0x00000006
	CommandStatus_ESME_RINVREGDLVFLG    CommandStatus = 0x00000007
	CommandStatus_ESME_RSYSERR          CommandStatus = 0x00000008
	CommandStatus_ESME_RINVSRCADR       CommandStatus = 0x0000000A
	CommandStatus_ESME_RINVDSTADR       CommandStatus = 0x0000000B
	CommandStatus_ESME_RINVMSGID        CommandStatus = 0x0000000C
	CommandStatus_ESME_RBINDFAIL        CommandStatus = 0x0000000D
	CommandStatus_ESME_RINVPASWD        CommandStatus = 0x0000000E
	CommandStatus_ESME_RINVSYSID        CommandStatus = 0x0000000F
	CommandStatus_ESME_RCANCELFAIL      CommandStatus = 0x00000011
	CommandStatus_ESME_RREPLACEFAIL     CommandStatus = 0x00000013
	CommandStatus_ESME_RMSGQFUL         CommandStatus = 0x00000014
	CommandStatus_ESME_RINVSERTYP       CommandStatus = 0x00000015
	CommandStatus_ESME_RINVNUMDESTS     CommandStatus = 0x00000033
	CommandStatus_ESME_RINVDLNAME       CommandStatus = 0x00000034
	CommandStatus_ESME_RINVDESTFLAG     CommandStatus = 0x00000040
	CommandStatus_ESME_RINVSUBREP       CommandStatus = 0x00000042
	CommandStatus_ESME_RINVESMCLASS     CommandStatus = 0x00000043
	CommandStatus_ESME_RCNTSUBDL        CommandStatus = 0x00000044
	CommandStatus_ESME_RSUBMITFAIL      CommandStatus = 0x00000045
	CommandStatus_ESME_RINVSRCTON       CommandStatus = 0x00000048
	CommandStatus_ESME_RINVSRCNPI       CommandStatus = 0x00000049
	CommandStatus_ESME_RINVDSTTON       CommandStatus = 0x00000050
	CommandStatus_ESME_RINVDSTNPI       CommandStatus = 0x00000051
	CommandStatus_ESME_RINVSYSTYP       CommandStatus = 0x00000053
	CommandStatus_ESME_RINVREPFLAG      CommandStatus = 0x00000054
	CommandStatus_ESME_RINVNUMMSGS      CommandStatus = 0x00000055
	CommandStatus_ESME_RTHROTTLED       CommandStatus = 0x00000058
	CommandStatus_ESME_RINVSCHED        CommandStatus = 0x00000061
	CommandStatus_ESME_RINVEXPIRY       CommandStatus = 0x00000062
	CommandStatus_ESME_RINVDFTMSGID     CommandStatus = 0x00000063
	CommandStatus_ESME_RX_T_APPN        CommandStatus = 0x00000064
	CommandStatus_ESME_RX_P_APPN        CommandStatus = 0x00000065
	CommandStatus_ESME_RX_R_APPN        CommandStatus = 0x00000066
	CommandStatus_ESME_RQUERYFAIL       CommandStatus = 0x00000067
	CommandStatus_ESME_RINVOPTPARSTREAM CommandStatus = 0x000000C0
	CommandStatus_ESME_ROPTPARNOTALLWD  CommandStatus = 0x000000C1
	CommandStatus_ESME_RINVPARLEN       CommandStatus = 0x000000C2
	CommandStatus_ESME_RMISSINGOPTPARAM CommandStatus = 0x000000C3
	CommandStatus_ESME_RINVOPTPARAMVAL  CommandStatus = 0x000000C4
	CommandStatus_ESME_RDELIVERYFAILURE CommandStatus = 0x000000FE
	CommandStatus_ESME_RUNKNOWNERR      CommandStatus = 0x000000FF
)

type TON uint8

const (
	TONUnknown          TON = 0x00
	TONInternational    TON = 0x01
	TONNational         TON = 0x02
	TONNetworkSpecific  TON = 0x03
	TONSubscriberNumber TON = 0x04
	TONAlphanumeric     TON = 0x05
	TONAbbreviated      TON = 0x06
)

type NPI uint8

const (
	NPIUnknown     NPI = 0x00
	NPI_ISDN       NPI = 0x01
	NPIData        NPI = 0x03
	NPITelex       NPI = 0x04
	NPILandMobile  NPI = 0x06
	NPINational    NPI = 0x08
	NPIPrivate     NPI = 0x09
	NPI_ERMES      NPI = 0x0A
	NPIInternet    NPI = 0x0E
	NPI_WAPClienId NPI = 0x12
)

type Tag uint16

const (
	TagDestAddrSubunit          Tag = 0x0005
	TagDestNetworkType          Tag = 0x0006
	TagDestBearerType           Tag = 0x0007
	TagDestTelematicsId         Tag = 0x0008
	TagSourceAddrSubunit        Tag = 0x000D
	TagSourceNetworkType        Tag = 0x000E
	TagSourceBearerType         Tag = 0x000F
	TagSourceTelematicsId       Tag = 0x0010
	TagQOSTimeToLive            Tag = 0x0017
	TagPayloadType              Tag = 0x0019
	TagAdditionalStatusInfoText Tag = 0x001D
	TagReceiptedMessageId       Tag = 0x001E
	TagMsMsgWaitFacilities      Tag = 0x0030
	TagPrivacyIndicator         Tag = 0x0201
	TagSourceSubaddress         Tag = 0x0202
	TagDestSubaddress           Tag = 0x0203
	TagUserMessageReference     Tag = 0x0204
	TagUserResponseCode         Tag = 0x0205
	TagSourcePort               Tag = 0x020A
	TagDestinationPort          Tag = 0x020B
	TagSARMsgRefNum             Tag = 0x020C
	TagLanguageIndicator        Tag = 0x020D
	TagSARTotalSegments         Tag = 0x020E
	TagSARSegmentSeqnum         Tag = 0x020F
	TagSCInterfaceVersion       Tag = 0x0210
	TagCallbackNumPresInd       Tag = 0x0302
	TagCallbackNumAtag          Tag = 0x0303
	TagNumberOfMessages         Tag = 0x0304
	TagCallbackNum              Tag = 0x0381
	TagDPFResult                Tag = 0x0420
	TagSetDPF                   Tag = 0x0421
	TagMSAvailabilityStatus     Tag = 0x0422
	TagNetworkErrorCode         Tag = 0x0423
	TagMessagePayload           Tag = 0x0424
	TagDeliveryFailureReason    Tag = 0x0425
	TagMoreMessagesToSend       Tag = 0x0426
	TagMessageState             Tag = 0x0427
	TagUssdServiceOP            Tag = 0x0501
	TagDisplayTime              Tag = 0x1201
	TagSMSSignal                Tag = 0x1203
	TagMSValidity               Tag = 0x1204
	TagAlertOnMessageDelivery   Tag = 0x130C
	TagITSReplyType             Tag = 0x1380
	TagITSSessionInfo           Tag = 0x1383
)

type DataCoding uint8

const (
	DataCodingSMSCDefaultAlphabet DataCoding = 0x00
	DataCodingASCII               DataCoding = 0x01
	DataCodingOctetUnspecified    DataCoding = 0x02
	DataCodingLatin1              DataCoding = 0x03
	DataCodingOctetUnspecified2   DataCoding = 0x04
	DataCodingJIS                 DataCoding = 0x05
	DataCodingCyrillic            DataCoding = 0x06
	DataCodingLatinHebrew         DataCoding = 0x07
	DataCodingUCS2                DataCoding = 0x08
	DataCodingPictogramEncoding   DataCoding = 0x09
	DataCodingISO2022JP           DataCoding = 0x0A
	DataCodingExtendedKanjiJIS    DataCoding = 0x0D
	DataCodingKS_C_5601           DataCoding = 0x0E
)

func (c DataCoding) String() string {
	switch c {
	case DataCodingSMSCDefaultAlphabet, DataCodingASCII, DataCodingOctetUnspecified, DataCodingLatin1,
		DataCodingOctetUnspecified2, DataCodingJIS, DataCodingCyrillic, DataCodingLatinHebrew, DataCodingUCS2,
		DataCodingPictogramEncoding, DataCodingISO2022JP, DataCodingExtendedKanjiJIS, DataCodingKS_C_5601:
		return strconv.Itoa(int(c))
	default:
		return ""
	}
}

type MessageState uint8

const (
	MessageStateENROUTE       MessageState = 0x01
	MessageStateDELIVERED     MessageState = 0x02
	MessageStateEXPIRED       MessageState = 0x03
	MessageStateDELETED       MessageState = 0x04
	MessageStateUNDELIVERABLE MessageState = 0x05
	MessageStateACCEPTED      MessageState = 0x06
	MessageStateUNKNOWN       MessageState = 0x07
	MessageStateREJECTED      MessageState = 0x08
)

func (s MessageState) String() string {
	switch s {
	case MessageStateENROUTE:
		return "ENROUTE"
	case MessageStateDELIVERED:
		return "DELIVERED"
	case MessageStateEXPIRED:
		return "EXPIRED"
	case MessageStateDELETED:
		return "DELETED"
	case MessageStateUNDELIVERABLE:
		return "UNDELIVERABLE"
	case MessageStateACCEPTED:
		return "ACCEPTED"
	case MessageStateUNKNOWN:
		return "UNKNOWN"
	case MessageStateREJECTED:
		return "REJECTED"
	default:
		return ""
	}
}
