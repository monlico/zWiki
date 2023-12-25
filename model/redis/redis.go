package redis

import (
	"context"
	"log"
	"zWiki/pkg/setting"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func init() {
	var (
		err            error
		password, host string
		dbNum          int
	)

	sec, err := setting.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbNum = sec.Key("DB").MustInt()
	password = sec.Key("PASSWORD").MustString("")
	host = sec.Key("HOST").String()

	Redis = redis.NewClient(&redis.Options{
		Addr:     host,     // Redis服务器地址
		Password: password, // 密码
		DB:       dbNum,    // 使用默认的数据库
	})

	if _, err := Redis.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	// 记录连接成功的日志
	log.Printf("Connected to Redis at %s", host)
}
