package database

import (
	"fmt"
	"oa/internal/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewData(db *gorm.DB, redisClient *redis.Client) *Data {
	return &Data{
		DB:          db,
		RedisClient: redisClient,
	}
}

func NewRedis() *redis.Client {
	// 初始化 Redis 客户端
	host := config.AppConfig.Redis.Host
	port := config.AppConfig.Redis.Port
	// 如果需要用户名，可以从配置中获取
	// user := config.AppConfig.Redis.User
	password := config.AppConfig.Redis.Password
	RedisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		// Username: ,
		Password: password,
	})
	return RedisClient
}

func NewDatabase() (*gorm.DB, error) {
	user := config.AppConfig.Database.User
	password := config.AppConfig.Database.Password
	host := config.AppConfig.Database.Host
	port := config.AppConfig.Database.Port
	dbname := config.AppConfig.Database.DBName
	charset := config.AppConfig.Database.Charset
	parseTime := config.AppConfig.Database.ParseTime
	loc := config.AppConfig.Database.Loc
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s", user, password, host, port, dbname, charset, parseTime, loc)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, nil
}
