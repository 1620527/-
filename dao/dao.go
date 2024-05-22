package dao

import (
	"fmt"
	"log"

	"ginStudy/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type User struct {
// 	ID   uint
// 	Name string
// 	Age  uint
// }

var Db *gorm.DB
var dbName string

// func InitDb() {
// 	sync.OnceFunc(func() {
// 		dsn := config.Conf.MysqlDB.Dsn
// 		var err error
// 		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 		if err != nil {
// 			log.Fatalf("无法连接到数据库: %v", err)
// 		}

// 		// 自动迁移模式 - 使用 `db.AutoMigrate` 函数自动创建或更新数据库表结构，确保与定义的数据模型一致。
// 		//Db.AutoMigrate(&User{})

// 		// 设置连接的最大生命周期为 1 小时
// 		// sqlDB, err2 := Db.DB()
// 		// if err2 != nil {
// 		// 	log.Fatalf("获取数据库连接失败: %v", err)
// 		// }
// 		// sqlDB.SetConnMaxLifetime(time.Hour)
// 		// sqlDB.SetMaxOpenConns(100)
// 		// sqlDB.SetMaxIdleConns(10)
// 		// fmt.Println("数据库连接成功")

// 		_ = Db.Raw("SELECT DATABASE();").Scan(&dbName).Error
// 		fmt.Println("当前数据库名称：", dbName)
// 	})
// }

func init() {

	fmt.Println("dao init")
	dsn := config.Conf.MysqlDB.Dsn
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	// 自动迁移模式 - 使用 `db.AutoMigrate` 函数自动创建或更新数据库表结构，确保与定义的数据模型一致。
	//Db.AutoMigrate(&User{})

	// 设置连接的最大生命周期为 1 小时
	// sqlDB, err2 := Db.DB()
	// if err2 != nil {
	// 	log.Fatalf("获取数据库连接失败: %v", err)
	// }
	// sqlDB.SetConnMaxLifetime(time.Hour)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetMaxIdleConns(10)
	// fmt.Println("数据库连接成功")

	_ = Db.Raw("SELECT DATABASE();").Scan(&dbName).Error
	fmt.Println("当前数据库名称：", dbName)
}
