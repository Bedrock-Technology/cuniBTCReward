package crontab

import (
	"cuniBTCReward/pkg/gormz"
	"testing"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_Reward(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3317)/cunibtcreward?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: gormz.NewGormLogger(true),
	})
	if err != nil {
		return
	}
	totalReward := decimal.New(1, 8)
	err = distributeExtraReward(db, "b1uniBTC", 1790070192, 11155111, totalReward, "0xd20568150Fd34eD3D457530a2f14E6AEfa791FF3", "0x48bAdE1602D122146Cc3FB0D89ad7E6004321B81")
	if err != nil {
		t.Log("err:", err)
	}
}
