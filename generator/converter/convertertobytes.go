// Package converter converts datatypes to array of bytes
package converter

import (
	"encoding/binary"
	"math"
)

// ConvertIntToBytes return a byte array which contains the representation for
// an integer value
func ConvertIntToBytes(value int) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(value))
	return bs
}

// ConvertFloatToBytes return a byte array which contains the representation for
// a float value
func ConvertFloatToBytes(value float32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, math.Float32bits(value))
	return bs
}
