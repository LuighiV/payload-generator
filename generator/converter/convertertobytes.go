package converter

import (
	"encoding/binary"
	"math"
)

func ConvertIntToBytes(value int) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(value))
	return bs
}

func ConvertFloatToBytes(value float32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, math.Float32bits(value))
	return bs
}
