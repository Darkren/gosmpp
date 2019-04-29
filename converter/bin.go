package converter

import "unsafe"

var (
	bigEndian = true
)

func init() {
	var x uint32 = 0x01020304
	if *(*byte)(unsafe.Pointer(&x)) == 0x04 {
		bigEndian = false
	}
}

func BytesToUint32(data [4]byte) uint32 {
	if bigEndian {
		return *(*uint32)(unsafe.Pointer(&data[0]))
	} else {
		dataReversed := [4]byte{data[3], data[2], data[1], data[0]}
		return *(*uint32)(unsafe.Pointer(&dataReversed[0]))
	}
}

func Uint32ToBytes(val uint32) [4]byte {
	data := *(*[4]byte)(unsafe.Pointer(&val))

	if bigEndian {
		return data
	} else {
		dataReversed := [4]byte{data[3], data[2], data[1], data[0]}
		return dataReversed
	}
}

func Uint64ToBytes(val uint64) [8]byte {
	data := *(*[8]byte)(unsafe.Pointer(&val))

	if bigEndian {
		return data
	} else {
		dataReversed := [8]byte{data[7], data[6], data[5], data[4], data[3], data[2], data[1], data[0]}
		return dataReversed
	}
}

func Uint16ToBytes(val uint16) [2]byte {
	data := *(*[2]byte)(unsafe.Pointer(&val))

	if bigEndian {
		return data
	} else {
		dataReversed := [2]byte{data[1], data[0]}
		return dataReversed
	}
}

func BytesToUint16(data [2]byte) uint16 {
	if bigEndian {
		return *(*uint16)(unsafe.Pointer(&data[0]))
	} else {
		dataReversed := [2]byte{data[1], data[0]}
		return *(*uint16)(unsafe.Pointer(&dataReversed[0]))
	}
}
