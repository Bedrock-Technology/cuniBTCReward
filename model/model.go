package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Cursor struct {
	gorm.Model
	// ChainId used for evm chain
	ChainId   uint   `gorm:"not null;default:0;index:t_chainid_name,unique"`
	ChainName string `gorm:"size:255;index:t_chainid_name"`
	// BlockNumber when we scan from_blockNUmber to to_blockNumber, has processed
	BlockNumber uint64 `gorm:"default:0"`
}

// EvmTransaction
// if someone deposit and withdrawal in one transaction, confliction will happen.
type EvmTransaction struct {
	gorm.Model
	// effect address, such as from or to
	Address     string `gorm:"size:128;index:t_address_chainid_hash,unique"`
	ChainId     uint   `gorm:"not null;default:0;index:t_address_chainid_hash"`
	Hash        string `gorm:"size:128;index:t_address_chainid_hash"`
	IndexInHash uint   `gorm:"default:0;index:t_address_chainid_hash"`
	// Contract is the contract address of this transaction, such as cuniBTC or redeemRouter, each stategy.
	Contract       string `gorm:"size:128;index:t_address_chainid_hash"`
	Token          string `gorm:"size:255"`
	BlockNumber    uint64 `gorm:"default:0;index:t_blocknumber"`
	BlockTimestamp uint64 `gorm:"default:0;index:t_blocktimestamp"`
	// + for income, - for outcome
	Amount    decimal.Decimal `gorm:"type:decimal(38);default:0"`
	LogMethod string          `gorm:"size:255"`
	Memo      string          `gorm:"size:255"`
}
type DelayRedeemRecord struct {
	gorm.Model
	ChainId           uint            `gorm:"not null;default:0;index:t_chainid_address,unique"`
	Address           string          `gorm:"size:128;index:t_chainid_address"`
	Contract          string          `gorm:"size:128;index:t_chainid_address"`
	CreateHash        string          `gorm:"size:128;index:t_chainid_address"`
	IndexInHash       uint            `gorm:"default:0;index:t_chainid_address"`
	Token             string          `gorm:"size:255"`
	Amount            decimal.Decimal `gorm:"type:decimal(38);default:0"`
	Fee               decimal.Decimal `gorm:"type:decimal(38);default:0"`
	Index             uint64          `gorm:"default:0"`
	CreateBlockNumber uint64          `gorm:"default:0;index:t_blocknumber"`
	CreateBlockTime   time.Time       `gorm:"default:0"`
	Claimed           bool            `gorm:"default:false"`
	ClaimTx           string          `gorm:"size:255"`
	ClaimAt           time.Time       `gorm:"default:0"`
}

type Strategy struct {
	gorm.Model
	ChainId           uint   `gorm:"not null;default:0;index:t_chainid_name,unique"`
	ChainName         string `gorm:"size:255;index:t_chainid_name"`
	Name              string `gorm:"size:255"`
	Symbol            string `gorm:"size:255;index:t_chainid_name"`
	CuniBTC           string `gorm:"size:255"`
	Vault             string `gorm:"size:255"`
	DelayRedeemRouter string `gorm:"size:255"`
	Airdrop           string `gorm:"size:255"`
}

type AirDropRecord struct {
	gorm.Model
	ChainId  uint            `gorm:"not null;default:0;index:t_chainid_address,unique"`
	Contract string          `gorm:"size:255;index:t_chainid_address"`
	Epoch    uint64          `gorm:"default:0;index:t_chainid_address"`
	Address  string          `gorm:"size:255;index:t_chainid_address"`
	Shares   decimal.Decimal `gorm:"type:decimal(38);default:0"`
	Amount   decimal.Decimal `gorm:"type:decimal(38);default:0"`
	Claimed  bool            `gorm:"default:false"`
	ClaimTx  string          `gorm:"size:255"`
	ClaimAt  time.Time       `gorm:"default:0"`
	Proof    string          `gorm:"type:longblob"`
}

type AirDropEpoch struct {
	gorm.Model
	ChainId   uint   `gorm:"not null;default:0;index:t_chainid_epoch,unique"`
	Contract  string `gorm:"size:255;index:t_chainid_epoch"`
	Epoch     uint64 `gorm:"default:0;index:t_chainid_epoch"`
	Token     string `gorm:"size:255;default:''"`
	Root      string `gorm:"size:255"`
	ValidTime uint64
	ActiveAt  time.Time
	Disabled  bool `gorm:"default:false"`
}

type Epoch struct {
	gorm.Model
	ChainId               uint   `gorm:"not null;default:0;index:t_chainid_epoch,unique"`
	Epoch                 uint64 `gorm:"default:0;index:t_chainid_epoch"`
	Contract              string `gorm:"size:255;index:t_chainid_epoch"`
	OperateStart          uint64 `gorm:"default:0"`
	OperateStartTimestamp uint64 `gorm:"default:0"`
	LockupStart           uint64 `gorm:"default:0"`
	LockupStartTimestamp  uint64 `gorm:"default:0"`
	StartGenesis          uint64 `gorm:"default:0"`
	StartGenesisTimestamp uint64 `gorm:"default:0"`
	OperatePeriod         uint64 `gorm:"default:0"`
	LockupPeriod          uint64 `gorm:"default:0"`
}

func GetCursor(database *gorm.DB, chainID uint) (*Cursor, error) {
	var cursor Cursor
	err := database.Model(&Cursor{}).Where("chain_id = ?", chainID).First(&cursor).Error
	return &cursor, err
}

func GetStrategy(database *gorm.DB, chainID uint) ([]Strategy, error) {
	var strategy []Strategy
	err := database.Model(&Strategy{}).Where("chain_id = ?", chainID).Find(&strategy).Error
	return strategy, err
}
