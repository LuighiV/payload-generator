package converter

import (
	"encoding/binary"
)

func ConvertIntToBytes(value int) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(value))
	return bs
}
