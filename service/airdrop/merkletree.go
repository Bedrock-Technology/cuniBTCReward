package airdrop

import (
	"cuniBTCReward/model"
	"encoding/json"
	"fmt"
	"time"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ethereum/go-ethereum/common"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

var leafEncodings = []string{
	smt.SOL_ADDRESS,
	smt.SOL_UINT256,
}

type TreeLeaf struct {
	Address string
	Amount  string
}

func GetMerkleRootAndProof(leaves []TreeLeaf) (*smt.StandardTree, error) {
	values := [][]interface{}{}
	for _, leaf := range leaves {
		values = append(values, []interface{}{smt.SolAddress(leaf.Address), smt.SolNumber(leaf.Amount)})
	}

	return smt.Of(values, leafEncodings)
}

func (a *Airdrop) CreateAirdropEpoch(chainId uint, contract string, epoch uint64, leaves []TreeLeaf, shares []string, dryRun bool) ([]byte, error) {
	if len(leaves) == 0 {
		return nil, fmt.Errorf("leaves is empty")
	}
	if len(leaves) != len(shares) {
		return nil, fmt.Errorf("leaves length does not match shares length")
	}

	tree, err := GetMerkleRootAndProof(leaves)
	if err != nil {
		return nil, err
	}
	proof, err := tree.DumpLeafProof()
	if err != nil {
		return nil, err
	}
	if len(proof.Proofs) != len(leaves) {
		return nil, fmt.Errorf("proof length does not match leaves length")
	}
	// AirDropRecord
	airdropRecords, err := lo.MapErr(proof.Proofs, func(proofData *smt.StandardMerkleLeafProofData, k int) (*model.AirDropRecord, error) {
		values := proofData.Value
		address, ok := values[0].(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert address to string")
		}
		amount, ok := values[1].(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert amount to string")
		}
		decimalAmount, err := decimal.NewFromString(amount)
		if err != nil {
			return nil, fmt.Errorf("failed to convert amount string to decimal: %v", err)
		}
		if decimalAmount.IsNegative() {
			return nil, fmt.Errorf("amount is negative for address[%s]", common.HexToAddress(address).String())
		}

		decimalShares, err := decimal.NewFromString(shares[k])
		if err != nil {
			return nil, fmt.Errorf("failed to convert shares string to decimal: %v", err)
		}
		if decimalShares.IsNegative() {
			return nil, fmt.Errorf("shares is negative for address[%s]", common.HexToAddress(address).String())
		}

		proofString, err := json.Marshal(proofData.Proof)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal proof: %v", err)
		}

		return &model.AirDropRecord{
			ChainId:  chainId,
			Contract: contract,
			Epoch:    epoch,
			Address:  common.HexToAddress(address).String(),
			Amount:   decimalAmount,
			Shares:   decimalShares,
			Claimed:  false,
			ClaimTx:  "",
			ClaimAt:  time.Unix(0, 0),
			Proof:    string(proofString),
		}, nil
	})
	if err != nil {
		return nil, err
	}

	if !dryRun {
		if err := a.database.CreateInBatches(airdropRecords, 500).Error; err != nil {
			return nil, err
		}
	}
	return tree.GetRoot(), nil
}
