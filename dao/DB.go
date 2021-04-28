package dao

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB         *gorm.DB
	TimeFormat = "2017-04-02 01:00:00"
	err        error
	Redisdb    *redis.Client
)

func init() {

	user := "mall"
	password := "mall"
	host := "37.99.211.213"
	dbName := "mall123"
	dbType := "mysql"
	tablePrefix := "fc_"

	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}

	DB.LogMode(true)
	gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	DB.SingularTable(true)

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	initRedis()
	fmt.Println("初始化完成")

}

//缓存
func initRedis() (err error) {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = Redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
