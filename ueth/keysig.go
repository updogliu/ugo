package ueth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func HexToECDSAPrivKey(hexKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(strings.TrimPrefix(hexKey, "0x"))
}

func ECDSAPrivKeyToHex(privKey *ecdsa.PrivateKey) string {
	return "0x" + hex.EncodeToString(crypto.FromECDSA(privKey))
}

func ECDSAPubKeyToHex(pubKey *ecdsa.PublicKey) string {
	return "0x" + hex.EncodeToString(crypto.FromECDSAPub(pubKey))
}

func AddrOfPrivKey(privKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privKey.PublicKey)
}

func NewTransactorFromPrivKeyHex(privKeyhex string, chainId *big.Int) (*bind.TransactOpts, error) {
	privKey, err := HexToECDSAPrivKey(privKeyhex)
	if err != nil {
		return nil, err
	}
	return NewEIP155Transactor(privKey, chainId), nil
}

// DO NOT use go-ethereum's `bind.NewKeyedTransactor`. Use this one instead.
//
// `bind.NewKeyedTransactor` would cause that all bound contracts turn out to use HomesteadSigner.
// See https://github.com/ethereum/go-ethereum/issues/16484 for details.
//
// The implementaton is adapted from go-ethereum's `bind.NewKeyedTransactor`.
func NewEIP155Transactor(key *ecdsa.PrivateKey, chainId *big.Int) *bind.TransactOpts {
	keyAddr := AddrOfPrivKey(key)
	eip155Signer := types.NewEIP155Signer(chainId)

	return &bind.TransactOpts{
		From: keyAddr,
		Signer: func(ignored_ types.Signer, address common.Address, tx *types.Transaction) (
			*types.Transaction, error) {

			if address != keyAddr {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := crypto.Sign(eip155Signer.Hash(tx).Bytes(), key)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(eip155Signer, signature)
		},
	}
}

// Preliminary check of whether `sig` can possibly be a valid Ethereum Homestead signature.
func PreValidateSig(sig []byte) bool {
	if len(sig) != 65 {
		return false
	}
	r := new(big.Int).SetBytes(sig[0:32])
	s := new(big.Int).SetBytes(sig[32:64])
	return crypto.ValidateSignatureValues(sig[64], r, s, true /*homestead*/)
}

// Get the recovered address.
func GetRecoveredAddr(hash []byte, sig []byte) (common.Address, error) {
	pubKey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pubKey), nil
}

// Verify that `sig` is the result of `signerAddr` signing `signedBytes`.
func VerifySig(signedBytes []byte, sig []byte, signerAddr common.Address) bool {
	if !PreValidateSig(sig) {
		return false
	}

	hash := crypto.Keccak256(signedBytes)
	recoveredAddr, err := GetRecoveredAddr(hash, sig)
	if err != nil {
		return false
	}
	return recoveredAddr == signerAddr
}

// Verify that `sig` is the result of `signerAddr` signing `hash`.
func VerifySigByHash(hash []byte, sig []byte, signerAddr common.Address) bool {
	if !PreValidateSig(sig) {
		return false
	}

	recoveredAddr, err := GetRecoveredAddr(hash, sig)
	if err != nil {
		return false
	}
	return recoveredAddr == signerAddr
}

func VerifyRpcSignedMsg(msg string, sig []byte, signerAddr common.Address) bool {
	signedStr := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(msg), msg)
	return VerifySig([]byte(signedStr), sig, signerAddr)
}

// Generate private key in a deterministic way (and thus unsafe).
//
// Do NOT use it in production code. Use go-ethereum/crypto.GenerateKey() instead.
func GenerateDeterministicKey_TestOnly(nonce int64) (*ecdsa.PrivateKey, error) {
	source := rand.NewSource(nonce ^ 0x123456789ABCDEF)
	return ecdsa.GenerateKey(crypto.S256(), rand.New(source))
}
