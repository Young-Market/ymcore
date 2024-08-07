package genesis_generator

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func TestGenerate(t *testing.T) {
	a, err := proto.MustKeyPair([]byte("test")).Addr(proto.MainNetScheme)
	require.NoError(t, err)
	block, err := GenerateGenesisBlock(proto.MainNetScheme, []GenesisTransactionInfo{{Address: a, Amount: 9000000000000000, Timestamp: 1558516864282}}, 153722867, 1558516864282)
	fmt.Println(block)
	require.NoError(t, err)
	require.Equal(t, 1, block.TransactionCount)
	ok, err := block.VerifySignature(proto.MainNetScheme)
	require.NoError(t, err)
	assert.True(t, ok)
}

func TestGenerateMainNet(t *testing.T) {
	timestamp := uint64(time.Now().UnixNano())
	txs := []GenesisTransactionInfo{
		{Address: proto.MustAddressFromString("2yHPenxgb4PBWajHkkf8meLNL514NDrHqqrB"), Amount: 8750000000000000, Timestamp: timestamp},
		{Address: proto.MustAddressFromString("2yHihYAyGPNmRZY2uRDvgfEMVo2ZLaykjMYG"), Amount: 7500000000000000, Timestamp: timestamp},
		{Address: proto.MustAddressFromString("2yHgtQ5v1erNhcjm6NyX3vNYRaSjoZWjTU8q"), Amount: 3750000000000000, Timestamp: timestamp},
		{Address: proto.MustAddressFromString("2yHQ8XBqs5gV6B9g7FQZV5mbZpGN8eZpX8Ry"), Amount: 5000000000000000, Timestamp: timestamp},
	}
	blocks, err := GenerateGenesisBlock(proto.MainNetScheme, txs, 153722867, timestamp)
	require.NoError(t, err)
	fmt.Println("Genesis Block Info:", blocks)
	sig := crypto.MustSignatureFromBase58(blocks.ID.String())
	block, err := RecreateGenesisBlock(proto.MainNetScheme, txs, 153722867, timestamp, sig)
	require.NoError(t, err)
	bb, err := block.MarshalBinary(proto.MainNetScheme)
	require.NoError(t, err)
	assert.Equal(t, 406, len(bb))
	assert.Equal(t, 4, block.TransactionCount)
	assert.Equal(t, 189, int(block.TransactionBlockLength))
	assert.Equal(t, sig.Bytes(), block.BlockSignature.Bytes())
	ok, err := block.VerifySignature(proto.MainNetScheme)
	require.NoError(t, err)
	assert.True(t, ok)
	txid11, err := blocks.Transactions[3].GetID(proto.MainNetScheme)
	ss, err := crypto.NewSignatureFromBytes(txid11)
	fmt.Println("Signature:", ss)
	txID1, err := block.Transactions[0].GetID(proto.MainNetScheme)
	require.NoError(t, err)
	assert.Equal(t,
		crypto.MustSignatureFromBase58("5kKxScNbxMwqBf5x1HhG8RNMLJv5Wa439LNk1e6799fVrxfXa6CANdqRBwu5pa281fNj34KD1Yp1QGqhc5HELJbh").Bytes(), txID1)

	txID2, err := block.Transactions[1].GetID(proto.MainNetScheme)
	require.NoError(t, err)
	assert.Equal(t,
		crypto.MustSignatureFromBase58("4aNFL6N3XBQNCBoEj8nSWn3Tu82AfKCZJSxRY5KMbcYLc2MpJPpwoJHefE6P51oVbiXRaHEtnpC536QQfAoR41fX").Bytes(), txID2)

	txID3, err := block.Transactions[2].GetID(proto.MainNetScheme)
	require.NoError(t, err)
	assert.Equal(t,
		crypto.MustSignatureFromBase58("2pWaeZVSZ3vMbZwZSiJQkChSQXzpfy7S4VPGfpDYLdkYLxtrJxpHsd7UX5qb7NJYpNrw4LE8DQgfURn6Lk4j7DhF").Bytes(), txID3)

	txID4, err := block.Transactions[3].GetID(proto.MainNetScheme)
	require.NoError(t, err)
	assert.Equal(t,
		crypto.MustSignatureFromBase58("5d3ftTpNMqdUQSEYXh1FVtqeWEeBqYuCD1JwBqYivDT4mhPqZxTRxg9NxBdNooCVzup31hNGgwLUU6EVZ1oJdxtn").Bytes(), txID4)

}

func TestGenerateDevNet(t *testing.T) {
	txs := []GenesisTransactionInfo{
		{Address: proto.MustAddressFromString("3FgScYB6MNdnN8m4xXddQe1Bjkwmd3U7YtM"), Amount: 6130000000000000, Timestamp: 1597073607702},
		{Address: proto.MustAddressFromString("3FWXhvWq2r8m54MmCEZ3YZkLg2qUdGWbU3V"), Amount: 15000000000000, Timestamp: 1597073607702},
		{Address: proto.MustAddressFromString("3FcSgww3tKZ7feQVmcnPFmRxsjqBodYz63x"), Amount: 25000000000000, Timestamp: 1597073607702},
		{Address: proto.MustAddressFromString("3FS5TnwA7xEXQ8LFRBdNk1MwqFR5SGz8vPn"), Amount: 25000000000000, Timestamp: 1597073607702},
		{Address: proto.MustAddressFromString("3FPzy3a12ccLUXTVTz5vhvkmVYXTXdVTKqF"), Amount: 40000000000000, Timestamp: 1597073607702},
		{Address: proto.MustAddressFromString("3FdEAz6F8xj37XVSUVTiqu8YfKBvtzWZZtn"), Amount: 45000000000000, Timestamp: 1597073607702},
		{Address: proto.MustAddressFromString("3FWMHWBXf5qzDenTFhUhT2tuqaoGnYHr6PM"), Amount: 50000000000000, Timestamp: 1597073607702},
		{Address: proto.MustAddressFromString("3FQntwq5KiXxEb8k2xLM6VGcZbBoTEroCsB"), Amount: 70000000000000, Timestamp: 1597073607702},
	}
	sig := crypto.MustSignatureFromBase58("5rDxRRzc9CM21j8XuAE1qp39svEr1BeLLF38HnchZd579ATdAPHqWxkt42AtoAV52GkVLU6F3TC2CWp2nzRKHpj8")
	block, err := RecreateGenesisBlock('D', txs, 5000, 1597073607702, sig)
	require.NoError(t, err)
	assert.Equal(t, 8, block.TransactionCount)
	ok, err := block.VerifySignature('D')
	require.NoError(t, err)
	assert.True(t, ok)
}
