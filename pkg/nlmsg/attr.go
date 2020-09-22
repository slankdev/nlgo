package nlmsg

import (
	"encoding/binary"
	"io"
)

type Attr struct {
	Len  uint16
	Type uint16
	Data []byte
}

func ReadAttr(r io.Reader) (*Attr, error) {
	var err error
	var atr Attr

	err = binary.Read(r, binary.LittleEndian, &atr.Len)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &atr.Type)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("to-read: %d\n", int(atr.Len))
	b := make([]byte, int(atr.Len)-4, int(atr.Len)-4)
	err = binary.Read(r, binary.LittleEndian, b)
	if err != nil {
		return nil, err
	}
	atr.Data = b

	//ReadDummy Buffer
	bl := 4 - (atr.Len % 4)
	if bl == 4 {
		bl = 0
	}
	// fmt.Printf("amari: %d\n", bl)
	bb := make([]byte, bl)
	err = binary.Read(r, binary.LittleEndian, bb)
	if err != nil {
		return nil, err
	}

	return &atr, nil
}
