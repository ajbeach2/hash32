package hash32

import (
	"crypto/md5"
	"encoding/base32"
	"strconv"
	"strings"
)

var base32NoPaddingEncoder = base32.StdEncoding.WithPadding(base32.NoPadding)

// Hash32 does an md5 hash of the input, and encodes to base32
func Hash32(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return strings.ToLower(string(
		base32NoPaddingEncoder.EncodeToString(
			hasher.Sum(nil),
		)))
}

// HashID a numberic inteteger to unique base32 string
func HashID(id int) string {
	bs := []byte(strconv.Itoa(id))
	result := strings.ToLower(string(base32NoPaddingEncoder.EncodeToString(bs)))
	return result
}
