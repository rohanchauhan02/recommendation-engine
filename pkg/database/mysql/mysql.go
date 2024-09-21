package database

import (
	"fmt"
	"time"

	mysqlLib "github.com/go-sql-driver/mysql"
	"github.com/rohanchauhan02/recommendation-engine/pkg/config"

	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	// register mysql with dd-trace
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
)

type MysqlSess interface {
	Init() (*gorm.DB, error)
}

type mysql struct {
	conf config.ImmutableConfig
}

func (db *mysql) Init() (*gorm.DB, error) {
	fmt.Println("Start open mysql connection...")
	mysqlConf := db.conf.GetDatabase().Mysql

	sqltrace.Register("mysql", &mysqlLib.MySQLDriver{}, sqltrace.WithServiceName("RECOMMENDATION_ENGINE"))
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.Name,
	)
	gormConfig := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	}

	gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Info)

	gormDB, err := gormtrace.Open(gormMysql.Open(connectionString), gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(8)
	sqlDB.SetMaxOpenConns(16)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	return gormDB, nil
}

func NewMysql(conf config.ImmutableConfig) MysqlSess {
	return &mysql{
		conf: conf,
	}
}
