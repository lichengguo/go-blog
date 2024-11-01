package model

import (
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	"go-blog-server/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db  *gorm.DB
	err error
)

// InitDb 初始化数据库连接
func InitDb() {
	// 拼接数据库连接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)

	// 连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 外键约束
		SkipDefaultTransaction:                   true, // 禁用默认事务（提高运行速度）
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名是 `user`
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败,请检查参数")
		panic(err)
	}

	sqlDB, _ := db.DB()
	// 设置连接池中的最大闲置连接数
	sqlDB.SetMaxIdleConns(16)

	// 设置数据库的最大连接数量
	sqlDB.SetMaxOpenConns(128)

	// 设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	// 迁移表
	_ = db.AutoMigrate(&User{}, &Article{}, &Category{})

	// 程序启动的时候，如果user表没有数据，则在数据库添加一个默认的管理用户
	DefaultUserAdd()

	// 此处不能关闭mysql连接，不然该函数结束以后就会断开和数据库的连接
	// 应该考虑在其他地方释放掉mysql连接，比如程序退出的时候
	//defer db.Close()
}
