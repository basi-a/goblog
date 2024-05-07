package global

import (

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func OpenSqlite() error {
	var err error
	DB, err = gorm.Open(sqlite.Open(Config.DBfile), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置此选项为true，以禁用表名自动加's'
		},
	})
	if err != nil {
		return err
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	//设置连接池参数
	sqlDB.SetConnMaxIdleTime(20)
	sqlDB.SetMaxOpenConns(1024)

	return nil
}
