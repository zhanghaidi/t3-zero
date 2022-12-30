package model

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var DB *gorm.DB

func Init(dataSource string) *gorm.DB {

	//启动Gorm支持
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t3_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	//如果出错就GameOver了
	if err != nil {
		log.Panicf("Database connection failed:%v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panicf("Database connection failed:%v", err)
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(20)

	//定义全局DB，供项目整体调用
	DB = db

	return db
}

var RDB *redis.Client

func InitRedis(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
		// MaxRetries: 1,
	})
	_, err := client.Ping().Result()

	if err != nil {
		log.Panic("Redis连接失败", err)
	}

	RDB = client

	return client
}
