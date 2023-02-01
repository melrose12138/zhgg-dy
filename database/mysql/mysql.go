package mysql

import (
	"fmt"
	"github.com/melrose12138/zhgg-dy/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// getDSN gets DSN from config file
func getDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%v&loc=%s",
		config.Cfg.DB.Username, config.Cfg.DB.Password, config.Cfg.DB.Host, config.Cfg.DB.Port,
		config.Cfg.DB.Database, config.Cfg.DB.Charset, config.Cfg.DB.ParseTime, config.Cfg.DB.Loc)
	fmt.Println(dsn)
	return dsn
}

func InitDB() {
	var err error
	// 加载配置文件已经完成
	dsn := getDSN()
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Open Database name: %s", DB.Name())
}
