package database

import (
	"fmt"
	"go-echo-template/common"
	"go-echo-template/config"
	"log"
	"os"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	dbCfg := config.AppConfig.DB

	cfg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Database,
	)
	common.Logger.LogInfo().Msg(cfg)
	db, err := gorm.Open(mysql.Open(cfg), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{SlowThreshold: time.Millisecond,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  true}),
	})
	if err != nil {
		common.Logger.LogError().Err(err).Msg("Database intialization")

		log.Fatal(err)
	}
	fmt.Println("Database connected")
	return db
}
