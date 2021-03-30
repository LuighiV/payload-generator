package converter

import (
	"encoding/binary"
)

func ConvertIntToBytes(value int) []byte {
	bs := make([]byte, 4)
	binary.PutVarint(bs, int64(value))
	return bs
}
