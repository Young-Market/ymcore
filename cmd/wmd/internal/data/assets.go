package data

import (
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

var (
	WavesID            = crypto.Digest{}
	WavesIssuerAddress = proto.WavesAddress{}
	WavesAssetInfo     = AssetInfo{ID: WavesID, Name: "YM", IssuerAddress: WavesIssuerAddress, Decimals: 8, Reissuable: false, Supply: 50000000000000000}
)

type AssetInfo struct {
	ID            crypto.Digest
	Name          string
	IssuerAddress proto.WavesAddress
	Decimals      uint8
	Reissuable    bool
	Supply        uint64
}
