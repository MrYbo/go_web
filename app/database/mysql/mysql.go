package mysql

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"login/app/config"
	"os"
	"time"
)

var DB *gorm.DB

func Init() {
	conf := config.Conf

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.LogLevel(conf.Mysql.LogLevel),
			Colorful: true,
		},
	)

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		conf.Mysql.User,
		conf.Mysql.Password,
		conf.Mysql.Path,
		conf.Mysql.Database,
		conf.Mysql.Config)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         conf.Mysql.DefaultStringSize,
		DisableDatetimePrecision:  conf.Mysql.DisableDatetimePrecision,
		DontSupportRenameIndex:    conf.Mysql.DontSupportRenameIndex,
		DontSupportRenameColumn:   conf.Mysql.DontSupportRenameColumn,
		SkipInitializeWithVersion: conf.Mysql.SkipInitializeWithVersion,
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		logrus.Error("mysql connection error: %s\n", err)
	}

	logrus.Info("mysql connection and Init success.")

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(conf.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
}
