package ubig

import (
	"fmt"
	"math/big"
	"strings"
)

func U64(v uint64) *big.Int {
	return new(big.Int).SetUint64(v)
}

func I64(v int64) *big.Int {
	return big.NewInt(v)
}

func Clone(x *big.Int) *big.Int {
	return new(big.Int).Set(x)
}

func Add(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

func Sub(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func Mul(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

func MulU64(x, y uint64) *big.Int {
	return Mul(U64(x), U64(y))
}

func Quo(x, y *big.Int) *big.Int {
	return new(big.Int).Quo(x, y)
}

func Lsh(x *big.Int, n uint) *big.Int {
	return new(big.Int).Lsh(x, n)
}

func Rsh(x *big.Int, n uint) *big.Int {
	return new(big.Int).Rsh(x, n)
}

func And(x, y *big.Int) *big.Int {
	return new(big.Int).And(x, y)
}

// Returns x**y
func Exp(x, y *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, nil)
}

// It accepts the formats 'b' (binary), 'o' (octal), 'd' (decimal), 'x' (lowercase hexadecimal), and
// 'X' (uppercase hexadecimal)
//
// Panics on error.
func MustScan(str string) *big.Int {
	v := new(big.Int)
	_, err := fmt.Sscan(str, v)
	if err != nil {
		panic(err)
	}
	return v
}

// Hex to big.Int. It is lenient about "0x" prefix.
func MustHex(hex string) *big.Int {
	if !strings.HasPrefix(hex, "0x") {
		hex = "0x" + hex
	}
	return MustScan(hex)
}

func IsBetween(x, min, max *big.Int) bool {
	return x.Cmp(min) >= 0 && x.Cmp(max) <= 0
}
