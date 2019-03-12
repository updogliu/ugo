package ueth

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

var ErrInvalidHexAddress = errors.New("Invalid Address")

// Returns whether `s` is an Eth address in hex form with "0x" prefix.
// e.g. IsHexAddr("0x1B38C16a23269Ef3D776A16B46b0f2f160Fcf7FC") returns true.
func IsHexAddr(s string) bool {
	if len(s) != 2+2*AddrByteLen {
		return false
	}
	if !strings.HasPrefix(s, "0x") {
		return false
	}
	for i := 2; i < len(s); i++ {
		if !('0' <= s[i] && s[i] <= '9' || 'A' <= s[i] && s[i] <= 'Z' || 'a' <= s[i] && s[i] <= 'z') {
			return false
		}
	}
	return true
}

// Returns whether addr is a EIP55-compliant mixed-case checksum address.
func IsNormalizedHexAddr(addr string) bool {
	normalized, err := NormalizeHexAddress(addr)
	if err != nil {
		return false
	}
	return addr == normalized
}

// NormalizeHexAddress formats `addr` as an EIP55-compliant mixed-case checksum address.
// It returns error if `addr` is not a valid hex address.
func NormalizeHexAddress(addr string) (string, error) {
	if !IsHexAddr(addr) {
		return "", ErrInvalidHexAddress
	}
	hex := common.HexToAddress(addr)
	return hex.Hex(), nil
}

// Panics if `s` is not an Eth address in hex form with "0x" prefix.
func MustHexToAddr(s string) common.Address {
	bytes, err := hexutil.Decode(s)
	if err != nil || len(bytes) != AddrByteLen {
		panic("Invalid hex address: " + s)
	}

	var addr common.Address
	addr.SetBytes(bytes)
	return addr
}
