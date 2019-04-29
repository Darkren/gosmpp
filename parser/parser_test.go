package parser

import (
	"bytes"
	"encoding/binary"
	"testing"
	"unsafe"
)

func BenchmarkNext(b *testing.B) {
	buf := []byte{0, 0, 1, 0}
	for i := 0; i < b.N; i++ {
		var y int32
		_ = binary.Read(bytes.NewReader(buf), binary.BigEndian, &y)
	}
}

func BenchmarkNextUnsafe(b *testing.B) {
	var x uint32 = 0x01020304
	bigEndian := true
	if *(*byte)(unsafe.Pointer(&x)) == 0x04 {
		bigEndian = false
	}

	buf := []byte{1, 1, 0, 0}
	for i := 0; i < b.N; i++ {
		if !bigEndian {
			buf2 := []byte{buf[3], buf[2], buf[1], buf[0]}
			_ = *(*uint32)(unsafe.Pointer(&buf2[0]))
		} else {
			_ = *(*uint32)(unsafe.Pointer(&buf[0]))
		}
	}
}
