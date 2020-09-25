package nlconst

import "fmt"

type Bit32 uint32

func (b Bit32) String() string {
	return fmt.Sprintf("0x%x", int(b))
}
