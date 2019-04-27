package hashish

import (
	"crypto/md5"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

var base32NoPaddingEncoder = base32.StdEncoding.WithPadding(base32.NoPadding)
var base64NoPaddingURLEncoder = base64.URLEncoding.WithPadding(base64.NoPadding)

// Hash32 does an md5 hash of the input, and encodes to base32
func Hash32(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return ToBase32(hasher.Sum(nil))
}

// ID32 a numberic inteteger to unique base32 string
func ID32(id int) string {
	bs := []byte(strconv.Itoa(id))
	return ToBase32(bs)
}

// ToBase32 converts input byte into base32 string
func ToBase32(input []byte) string {
	return strings.ToLower(string(base32NoPaddingEncoder.EncodeToString(input)))
}

// ToBase64 converts input byte into base64 string
func ToBase64(input []byte) string {
	return strings.ToLower(string(base64NoPaddingURLEncoder.EncodeToString(input)))
}

// StrToBase32 converts input string into base32 string
func StrToBase32(str string) string {
	input := []byte(str)
	return ToBase32(input)
}

// UUIDBinary returns a base10 int64 representation of a uuid
func UUIDBinary() uint64 {
	id, _ := uuid.New().MarshalBinary()
	return binary.BigEndian.Uint64(id)
}

// UUID32 returns a UUID encoded to base32
func UUID32() (string, error) {
	id, err := uuid.New().MarshalBinary()
	if err != nil {
		return "", err
	}
	return ToBase32(id), nil
}

// UUIDTo32 takes a uuid string, parses, and converts to base32
func UUIDTo32(input string) (string, error) {
	id, err := uuid.Parse(input)
	if err != nil {
		return "", err
	}

	bin, err := id.MarshalBinary()
	if err != nil {
		return "", err
	}

	return ToBase32(bin), nil
}

// UUIDTo64 takes a uuid string, parses, and converts to base32
func UUIDTo64(input string) (string, error) {
	id, err := uuid.Parse(input)
	if err != nil {
		return "", err
	}

	bin, err := id.MarshalBinary()
	if err != nil {
		return "", err
	}

	return ToBase64(bin), nil
}
