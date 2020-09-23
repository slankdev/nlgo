package util

import (
	"encoding/binary"
	"fmt"
)

func ByteToString(b []byte) string {
	return string(b)
}

func ByteToU64(b []byte) string {
	val := binary.LittleEndian.Uint64(b)
	return fmt.Sprintf("%d", val)
}

func ByteToU32(b []byte) string {
	val := binary.LittleEndian.Uint32(b)
	return fmt.Sprintf("%d", val)
}

func ByteToU16(b []byte) string {
	val := binary.LittleEndian.Uint16(b)
	return fmt.Sprintf("%d", val)
}

func ByteToU8(b []byte) string {
	return fmt.Sprintf("%d", b[0])
}
