package model

import (
	"cuniBTCReward/pkg/gormz"
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMigrate(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3317)/cunibtcreward?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: gormz.NewGormLogger(),
	})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").
		AutoMigrate(&EvmTransaction{}, &Cursor{}, &Strategy{}, &AirDropRecord{}, &AirDropEpoch{}, &Epoch{}, &DelayRedeemRecord{})
	fmt.Printf("err: %v", err)
}
