package crontab

import (
	"cuniBTCReward/api/internal/config"
	"cuniBTCReward/model"

	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ExtraReward struct {
	db       *gorm.DB
	symbol   string
	queryUrl string
	config   config.Config
}

func NewExtraReward(queryUrl, symbolReward string, sqlDB *gorm.DB, c config.Config) *ExtraReward {
	return &ExtraReward{
		db:       sqlDB,
		symbol:   symbolReward,
		queryUrl: queryUrl,
		config:   c,
	}
}

func (e *ExtraReward) ExtraRewardCron() {

}

func distributeExtraReward(db *gorm.DB, symbol string, timestamp uint64, chainID int64, totalRewards decimal.Decimal, vault, redeem string) (err error) {
	var records []model.ExtraRewardDis
	rec := model.ExtraRewardRec{
		Symbol:      symbol,
		Timestamp:   timestamp,
		TotalAmount: totalRewards,
	}

	sql := `SELECT t.address,
       COALESCE(SUM(t.amount), 0) AS share,
	   MAX(CASE WHEN t.contract = ? THEN t.block_timestamp END) AS last_deposit_time
FROM evm_transactions t
WHERE t.chain_id = ? AND t.contract IN ? AND t.deleted_at IS NULL AND t.block_timestamp <= ?
GROUP BY t.address
HAVING share > 0`

	args := []interface{}{
		vault, chainID,
		[]string{vault, redeem},
		timestamp,
	}
	err = db.Raw(sql, args...).Scan(&records).Error
	if err != nil {
		return err
	}
	if len(records) == 0 {
		logx.Infof("no records symbol: %s", symbol)
		if err := db.Create(&rec).Error; err != nil {
			return err
		}
		return nil
	}
	totalShare := decimal.Zero
	for _, r := range records {
		totalShare = totalShare.Add(r.Share)
	}
	if totalShare.IsZero() {
		logx.Infof("totalshare zero symbol: %s", symbol)
		if err := db.Create(&rec).Error; err != nil {
			return err
		}
		return nil
	}

	for i := range records {
		records[i].Amount = records[i].Share.Div(totalShare).Mul(totalRewards).Floor()
		records[i].Symbol = symbol
		records[i].Timestamp = timestamp
	}
	//save to db
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := db.Create(&rec).Error; err != nil {
			return err
		}

		if err := db.CreateInBatches(records, 50).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logx.Errorf("distributeExtraReward symbol: %s,error: %v", symbol, err)
	}
	return
}
