package Database

import (
	"github.com/garyburd/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// Mysql连接

type Mysql struct {
	mysqlpath string
	db *gorm.DB
}
func connMySQL(mysqlpath string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(mysqlpath), &gorm.Config{})
	if err != nil{
		log.Fatalln("Mysql连接错误：", err)
	}
	return db
}

func NewMysql(mysqlpath string) *Mysql {
	return &Mysql{
		mysqlpath: mysqlpath,
		db: connMySQL(mysqlpath),
	}
}

type Redis struct {
	Host string // 127.0.0.1:6379
	MaxIdle int //最大空闲链接数量
	MaxActive int //表示和数据库最大链接数，0表示，并发不限制数量
	IdleTimeout time.Duration //最大空闲时间，用完链接后100秒后就回收到链接池
	DialDatabase int // redis数据库
	RedisDb *redis.Pool
}

func connRedis(Host string, MaxIdle int, MaxActive int, IdleTimeout time.Duration, DialDatabase int)  *redis.Pool{
	return &redis.Pool{
		MaxIdle: MaxIdle, //最大空闲链接数量
		MaxActive: MaxActive, //表示和数据库最大链接数，0表示，并发不限制数量
		IdleTimeout: IdleTimeout, //最大空闲时间，用完链接后100秒后就回收到链接池
		Dial: func()(redis.Conn,error){
			return redis.Dial("tcp",Host, redis.DialDatabase(DialDatabase))
		},
	}
}

func NewRedis(Host string, MaxIdle int, MaxActive int, IdleTimeout time.Duration, DialDatabase int)  *Redis{
	return &Redis{
		Host: Host,
		MaxIdle: MaxIdle,
		MaxActive: MaxActive,
		IdleTimeout: IdleTimeout,
		DialDatabase: DialDatabase,
		RedisDb: connRedis(Host, MaxIdle, MaxActive, IdleTimeout, DialDatabase),
	}

}