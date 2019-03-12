// This file contains common functions of calling Eth-RPC.

package ueth

import (
	"context"
	"encoding/json"
	"math/big"
	"sort"
	"strconv"
	"time"

	"fmt"

	"github.com/updogliu/ugo/ulog"
	"github.com/updogliu/ugo/utime"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Returns the number of most recent block.
func GetBlockNumber(
	ethNodeUrl string, ctx context.Context) (int64, error) {

	cli, err := rpc.DialContext(ctx, ethNodeUrl)
	if err != nil {
		return 0, err
	}
	defer cli.Close()

	var result *hexutil.Big
	err = cli.CallContext(ctx, &result, "eth_blockNumber")
	if err != nil {
		return 0, err
	}
	return result.ToInt().Int64(), nil
}

// Receipt is a customized version of types.Receipt, such that Parity can be used as backend.
type Receipt struct {
	TxHash            common.Hash  `json:"transactionHash"`
	TxIndex           *hexutil.Big `json:"transactionIndex"`
	BlockHash         common.Hash  `json:"blockHash"`
	BlockNumber       *hexutil.Big `json:"blockNumber"`
	CumulativeGasUsed *hexutil.Big `json:"cumulativeGasUsed"`
	GasUsed           *hexutil.Big `json:"gasUsed"`
	StatusBig         *hexutil.Big `json:"status"`
	Status            uint64       `json:"-"`
}

func (receipt *Receipt) Successful() bool {
	return receipt.Status == types.ReceiptStatusSuccessful
}

// Returns an error if the tx is not found, or failed to parse, or not confirmed yet.
//
// Note:
// - Parity and Geth disagree on whether to return a receipt for pending tx for rpc
//   `eth_getTransactionReceipt`. See https://github.com/ethereum/go-ethereum/issues/3323
// - The response of "eth_getTransactionReceipt" of parity cannot be unmarshaled by the geth native
//   type, i.e. go-ethereum/core/types.Receipt because the element of the `Logs` field does not have
//   the `transactionHash` field, which is required by geth unmarshaler.
func GetTransactionReceipt(
	ethNodeUrl string, ctx context.Context, txHash common.Hash) (*Receipt, error) {

	cli, err := rpc.DialContext(ctx, ethNodeUrl)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	var result *Receipt
	err = cli.CallContext(ctx, &result, "eth_getTransactionReceipt", txHash)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, ethereum.NotFound
	}
	if result.TxIndex == nil {
		return nil, fmt.Errorf("Missing Receipt TxIndex: %v", result)
	}
	if result.StatusBig == nil {
		return nil, fmt.Errorf("Missing Receipt Status: %v", result)
	}
	if result.BlockNumber == nil {
		return nil, fmt.Errorf("Transaction has not been confirmed yet: %v", txHash.Hex())
	}
	result.Status = result.StatusBig.ToInt().Uint64()
	return result, nil
}

// hash: DATA, 32 Bytes - hash of the transaction.
// nonce: QUANTITY - the number of transactions made by the sender prior to this one.
// *blockHash: DATA, 32 Bytes - hash of the block where this tx was in or null if it is pending.
// *blockNumber: QUANTITY - block number where this tx was in or null if it is pending.
type Transaction struct {
	Hash        common.Hash     `json:"hash"`
	BlockHash   *common.Hash    `json:"blockHash"`
	BlockNumber *hexutil.Big    `json:"blockNumber"`
	From        common.Address  `json:"from"`
	To          *common.Address `json:"to"`
	Input       *hexutil.Bytes  `json:"input"`
	NonceBig    *hexutil.Big    `json:"nonce"`
	TxIndexBig  *hexutil.Big    `json:"transactionIndex"`
	Value       *hexutil.Big    `json:"value"` // in Wei

	// Remember to fill these with `NonceBig` and `TxIndexBig` when make an instance NOT from
	// Unmarshaling JSON.
	Nonce   int64 `json:"-"`
	TxIndex int64 `json:"-"`

	// Add more fields on demand.
	// Spec: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionbyhash
}

func (tx *Transaction) UnmarshalJSON(bs []byte) error {
	type txFieldsUnmarshal Transaction
	if err := json.Unmarshal(bs, (*txFieldsUnmarshal)(tx)); err != nil {
		return err
	}

	if tx.NonceBig == nil {
		return fmt.Errorf("Missing Transaction Nonce")
	}
	tx.Nonce = tx.NonceBig.ToInt().Int64()

	if tx.TxIndexBig == nil {
		return fmt.Errorf("Missing Transaction Index")
	}
	tx.TxIndex = tx.TxIndexBig.ToInt().Int64()
	return nil
}

func (tx *Transaction) String() string {
	return fmt.Sprint("{",
		"Hash: ", tx.Hash.Hex(),
		", BlockHash: ", tx.BlockHash.Hex(),
		", BlockNumber: ", tx.BlockNumber.ToInt(),
		", From: ", tx.From.Hex(),
		", To: ", tx.To.Hex(),
		", Input: ", tx.Input.String(),
		", Nonce: ", tx.Nonce,
		", TxIndex: ", tx.TxIndex,
		", Value: ", tx.Value.ToInt(),
		"}")
}

// Returns an error if the tx is not found or failed to parse.
// Returns a tx with nil BlockHash and BlockNumber if it has not been confirmed.
func GetTransactionByHash(
	ethNodeUrl string, ctx context.Context, txHash common.Hash) (*Transaction, error) {

	cli, err := rpc.DialContext(ctx, ethNodeUrl)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	var result *Transaction
	err = cli.CallContext(ctx, &result, "eth_getTransactionByHash", txHash)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, ethereum.NotFound
	}

	if result.BlockHash != nil {
		if result.BlockHash.Big().BitLen() == 0 {
			result.BlockHash = nil
		}
	}
	if (result.BlockHash == nil) != (result.BlockNumber == nil) {
		return nil, fmt.Errorf(
			"BlockHash and BlockNumber should be both nil or both non-nil. %+v", result)
	}
	return result, nil
}

// Gets all transactions in the block specified by `blockHash`.
func GetTxsOfBlock(ethNodeUrl string, ctx context.Context, blockHash common.Hash) (
	txs []*Transaction, retErr error) {

	cli, err := rpc.DialContext(ctx, ethNodeUrl)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	type Wrapper struct {
		Txs json.RawMessage `json:"transactions"`
	}
	var wrapper *Wrapper
	err = cli.CallContext(ctx, &wrapper, "eth_getBlockByHash", blockHash, true)
	if err != nil {
		return nil, err
	}
	if wrapper == nil {
		return nil, ethereum.NotFound
	}

	if err := json.Unmarshal(wrapper.Txs, &txs); err != nil {
		return nil, fmt.Errorf("Error in decoding txs: %v. Raw: %v", err, string(wrapper.Txs))
	}
	for _, tx := range txs {
		if tx == nil {
			return nil, fmt.Errorf("Nil tx decoded from %v", string(wrapper.Txs))
		}
		if tx.NonceBig == nil {
			return nil, fmt.Errorf("Missing Transaction Nonce")
		}
		tx.Nonce = tx.NonceBig.ToInt().Int64()
	}
	return txs, nil
}

type BlockHeader struct {
	ParentHash  common.Hash  `json:"parentHash"`
	UncleHash   common.Hash  `json:"sha3Uncles"`
	Root        common.Hash  `json:"stateRoot"`
	Hash        common.Hash  `json:"hash"`
	ReceiptHash common.Hash  `json:"receiptsRoot"`
	NumberBig   *hexutil.Big `json:"number"`
	Number      int64        `json:"-"`
	TimeBig     *hexutil.Big `json:"timestamp"`
	TimeSec     int64        `json:"-"`
}

// Gets block header by number. Gets header of the latest block if `blockNumber < 0`.
// Returns `ethereum.NotFound` if the block is not found.
// Returns `ethereum.NotFound` for older blocks if Ethereum client is not a full node.
//
// Note:
//   We cannot use `ethclient.HeaderByNumber` because that function requires that the `mixHash`
// field in the JSON-RPC response, which is used in the calculation of `types.Header.Hash()`.
// But parity EthRPC node does not returns the `mixHash` field.
func GetBlockHeaderByNumber(
	ethNodeUrl string, ctx context.Context, blockNumber int64) (*BlockHeader, error) {

	cli, err := rpc.DialContext(ctx, ethNodeUrl)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	var header *BlockHeader
	var blockNumberArg string
	if blockNumber < 0 {
		blockNumberArg = "latest"
	} else {
		blockNumberArg = "0x" + strconv.FormatInt(blockNumber, 16)
	}
	// If true it returns the full transaction objects, if false only the hashes of the transactions.
	returnAllTxs := false
	err = cli.CallContext(ctx, &header, "eth_getBlockByNumber", blockNumberArg, returnAllTxs)
	if err != nil {
		return nil, err
	}
	if header == nil {
		return nil, ethereum.NotFound
	}
	header.Number = header.NumberBig.ToInt().Int64()
	header.TimeSec = header.TimeBig.ToInt().Int64()
	return header, nil
}

func GetLatestBlockHeader(ethNodeUrl string, ctx context.Context) (*BlockHeader, error) {
	return GetBlockHeaderByNumber(ethNodeUrl, ctx, -1)
}

// Returns (client, chainId, nil) if the Eth RPC node is ready in time; otherwise returns an error.
func WaitEthNodeReady(ethNodeUrl string, timeout time.Duration) (
	*ethclient.Client, *big.Int, error) {

	var ethClient *ethclient.Client
	var err error
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	for round := 1; ctx.Err() == nil; round++ {
		var chainId *big.Int
		ethClient, err = ethclient.Dial(ethNodeUrl)
		if err == nil {
			chainId, err = ethClient.NetworkID(ctx) // try sending an rpc
		}

		if err != nil {
			if round%50 == 0 {
				ulog.Infof("Waiting for ETH node %v to be ready. Last error: %v", ethNodeUrl, err)
			}
			utime.SleepMs(100)
			continue
		}

		ulog.Infof("ETH node %v is ready", ethNodeUrl)
		return ethClient, chainId, nil
	}
	return nil, nil, fmt.Errorf("Wait ETH node %v timed out. Last error: %v", ethNodeUrl, err)
}

// Returns nil if and only if the tx has been confirmed enough times.
// Precondition: `minConfirmTimes` > 0
//
// Note: Usually the caller wants to check the tx receipt status after calling this function.
func WaitTxConfirmed(
	ethNodeUrl string, ctx context.Context, txHash common.Hash, minConfirmTimes int) error {

	if minConfirmTimes <= 0 {
		ulog.Panic("minConfirmTimes <= 0")
	}

	var latestHeader *BlockHeader
	var confirmedTimes int64
	for round := 1; ctx.Err() == nil; round++ {
		tx, err := GetTransactionByHash(ethNodeUrl, ctx, txHash)
		if err != nil || tx.BlockNumber == nil {
			goto Continue
		}

		if minConfirmTimes == 1 { // no need to check chain height
			return nil
		}

		latestHeader, err = GetLatestBlockHeader(ethNodeUrl, ctx)
		if err != nil {
			goto Continue
		}

		confirmedTimes = latestHeader.Number - tx.BlockNumber.ToInt().Int64() + 1
		ulog.Infof("Tx %v confirmed %v/%v times...", txHash.Hex(), confirmedTimes, minConfirmTimes)
		if confirmedTimes >= int64(minConfirmTimes) {
			return nil
		}

	Continue:
		if err != nil && round%20 == 0 {
			ulog.Error("Got an error during waiting tx confirmed: ", err)
		}
		if ctx.Err() == nil {
			time.Sleep(250 * time.Millisecond) // wait for a check interval
		}
	} // for loop

	return fmt.Errorf("WaitTxConfirmed timed out. Tx: %v", txHash.Hex())
}

// Wraps `FilterLogs` with post-sorting - just in case.
func QueryLogs(backend bind.ContractBackend, ctx context.Context, query ethereum.FilterQuery) (
	[]types.Log, error) {

	logs, err := backend.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	// Order by 1) BlockNumber, 2) TxIndex, 3) Log index in the receipt
	isBeforeFn := func(i, j int) bool {
		if logs[i].BlockNumber != logs[j].BlockNumber {
			return logs[i].BlockNumber < logs[j].BlockNumber
		}
		if logs[i].TxIndex != logs[j].TxIndex {
			return logs[i].TxIndex < logs[j].TxIndex
		}
		return logs[i].Index < logs[j].Index
	}

	if !sort.SliceIsSorted(logs, isBeforeFn) {
		ulog.Warning("Raw result from FilterLogs is NOT sorted.")
		sort.Slice(logs, isBeforeFn)
	}
	return logs, nil
}
