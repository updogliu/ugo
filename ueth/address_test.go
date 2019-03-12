package ueth

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestNormalizeAddress(t *testing.T) {
	addr1 := ""
	_, err := NormalizeHexAddress(addr1)
	if err != ErrInvalidHexAddress {
		t.FailNow()
	}

	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md
	addr2 := "0xFc571a5fA85fd82393FbD3Ff9d74583a000C174c"
	formattedAddr2, err := NormalizeHexAddress(addr2)
	if err != nil {
		t.FailNow()
	}

	if addr2 != formattedAddr2 {
		t.FailNow()
	}

	addr3 := strings.ToLower(addr2)
	formattedAddr3, err := NormalizeHexAddress(addr3)
	if err != nil {
		t.FailNow()
	}
	if addr3 == formattedAddr3 || addr2 != formattedAddr3 {
		t.FailNow()
	}
}

func TestPrintSomeNormalizedAddr(t *testing.T) {
	fmt.Println("Print some normalized addresses:")
	for i := 0; i < 10; i++ {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		fmt.Print(ECDSAPrivKeyToHex(key), ": ", AddrOfPrivKey(key).Hex(), "\n")
	}
}

func TestMustHexToAddr(t *testing.T) {
	hex := "0xFc571a5fA85fd82393FbD3Ff9d74583a000C174c"
	require.Equal(t, common.HexToAddress(hex), MustHexToAddr(hex))

	lowercase := "0xa4e8c3ec456107ea67d3075bf9e3df3a75823db0"
	require.Equal(t, common.HexToAddress(lowercase), MustHexToAddr(lowercase))

	require.Panics(t, func() { MustHexToAddr("Fc571a5fA85fd82393FbD3Ff9d74583a000C174c") })
	require.Panics(t, func() { MustHexToAddr("0xFc571a5fA85fd82393FbD3Ff9d74583a000C174c00") })
	require.Panics(t, func() { MustHexToAddr("0xFc571a5fA85fd82393FbD3Ff9d74583a000C174g") })
}
