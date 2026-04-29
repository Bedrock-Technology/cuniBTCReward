package airdrop

import (
	"cuniBTCReward/model"
	"testing"

	"github.com/shopspring/decimal"
)

func TestCalcAvgAmount(t *testing.T) {
	startAmount := decimal.NewFromUint64(3000000000) //30
	startBlock := uint64(100)
	endBlock := uint64(200)
	transactions := []*model.EvmTransaction{
		{
			BlockNumber: 110,
			Amount:      decimal.NewFromInt(-1000000000), //-10
		},
		{
			BlockNumber: 120,
			Amount:      decimal.NewFromInt(500000000), //+5
		},
		{
			BlockNumber: 150,
			Amount:      decimal.NewFromInt(-2000000000), //-20
		},
	}

	expectedAvgAmount := decimal.NewFromInt(1500000000)
	actualAvgAmount, err := CalcAvgAmount(startAmount, transactions, startBlock, endBlock)
	if err != nil {
		t.Fatalf("CalcAvgAmount returned error: %v", err)
	}
	if !actualAvgAmount.Equal(expectedAvgAmount) {
		t.Errorf("CalcAvgAmount returned %s, expected %s", actualAvgAmount.String(), expectedAvgAmount.String())
	}
}
func TestCalcAvgAmount2(t *testing.T) {
	startAmount := decimal.NewFromUint64(3000000000) //30
	startBlock := uint64(100)
	endBlock := uint64(200)
	transactions := []*model.EvmTransaction{
		{
			BlockNumber: 110,
			Amount:      decimal.NewFromInt(-1000000000), //-10
		},
		{
			BlockNumber: 120,
			Amount:      decimal.NewFromInt(500000000), //+5
		},
		{
			BlockNumber: 200,
			Amount:      decimal.NewFromInt(-2000000000), //-20
		},
	}

	expectedAvgAmount := decimal.NewFromInt(2500000000)
	actualAvgAmount, err := CalcAvgAmount(startAmount, transactions, startBlock, endBlock)
	if err != nil {
		t.Fatalf("CalcAvgAmount returned error: %v", err)
	}
	if !actualAvgAmount.Equal(expectedAvgAmount) {
		t.Errorf("CalcAvgAmount returned %s, expected %s", actualAvgAmount.String(), expectedAvgAmount.String())
	}
}
func TestCalcAvgAmount3(t *testing.T) {
	startAmount := decimal.NewFromUint64(3000000000) //30
	startBlock := uint64(100)
	endBlock := uint64(200)
	transactions := []*model.EvmTransaction{
		{
			BlockNumber: 100,
			Amount:      decimal.NewFromInt(-1000000000), //-10
		},
		{
			BlockNumber: 120,
			Amount:      decimal.NewFromInt(500000000), //+5
		},
		{
			BlockNumber: 200,
			Amount:      decimal.NewFromInt(-2000000000), //-20
		},
	}

	expectedAvgAmount := decimal.NewFromInt(2400000000)
	actualAvgAmount, err := CalcAvgAmount(startAmount, transactions, startBlock, endBlock)
	if err != nil {
		t.Fatalf("CalcAvgAmount returned error: %v", err)
	}
	if !actualAvgAmount.Equal(expectedAvgAmount) {
		t.Errorf("CalcAvgAmount returned %s, expected %s", actualAvgAmount.String(), expectedAvgAmount.String())
	}
}
