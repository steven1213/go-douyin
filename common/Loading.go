package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"steven.com/go-douyin/router"
	"time"
)

var Db *gorm.DB
var err error

func Loading() {
	// 读取系统配置
	sysConfig := NewConfig().ReadConfig()
	// 加载Banner
	readBanner(sysConfig.Server.Banner.Name)
	// 加载mysql数据库连接
	InitDb(sysConfig.DB)
	// 加载redis数据库连接
	InitRedisDB(sysConfig.Redis)
	// 加载路由
	err := router.InitRouter(sysConfig.Server.Model).Run(sysConfig.Server.Host + ":" + sysConfig.Server.Port)
	if err != nil {
		log.Fatalf("server run error: %s\n", err)
		return
	}
}

type RDBManager struct {
	OpenTx bool
	DsName string
	Db     *gorm.DB
	Tx     *gorm.Tx
	Errors []error
}

func InitDb(dbConfig *DbConfig) {
	//启用打印日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level: Silent、Error、Warn、Info
			Colorful:      false,       // 禁用彩色打印
		},
	)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password,
		dbConfig.Host, dbConfig.Port, dbConfig.Database)

	dialect := mysql.New(mysql.Config{
		DSN:                       dbURI,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	})

	conn, err := gorm.Open(dialect, &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("connection database[%s:%s/%s] error: %s\n",
			dbConfig.Host, dbConfig.Port, dbConfig.Database, err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatalf("connection database[%s:%s/%s] error: %s\n",
			dbConfig.Host, dbConfig.Port, dbConfig.Database, err)
	}
	sqlDB.SetMaxOpenConns(dbConfig.DbPool.MaxOpenConns)
	sqlDB.SetMaxIdleConns(dbConfig.DbPool.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(dbConfig.DbPool.ConnMaxLifetime)

	log.Printf("connection database[%s:%s/%s] success\n",
		dbConfig.Host, dbConfig.Port, dbConfig.Database)
}

func InitRedisDB(redis *RedisConfig) {

}

// readBanner 读取banner
func readBanner(name string) {
	file, err := os.ReadFile("config/" + name)
	if err != nil {
		log.Panicln("read banner error", err)
	}
	fmt.Println(string(file))
}
