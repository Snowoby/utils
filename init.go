package utils

import (
	"crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
)

func init() {
	b := make([]byte, 16)
	_, err:=rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}
