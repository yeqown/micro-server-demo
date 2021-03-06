package global

import (
	"github.com/yeqown/micro-server-demo/internal/repository"
	"github.com/yeqown/micro-server-demo/pkg/types"

	"github.com/yeqown/infrastructure/framework/gormic"
	logger "github.com/yeqown/infrastructure/framework/logrus-logger"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

var (
	_cfg *types.Config

	_mysqlDB     *gorm.DB
	_redisClient *redis.Client
	_mgoSession  *mgo.Session

	// Repos of current services
	Repos wholeRepo
)

func GetConfig() *types.Config {
	return _cfg
}

// GetMysqlDB .
func GetMysqlDB() *gorm.DB {
	return _mysqlDB
}

// GetRedisClient .
func GetRedisClient() *redis.Client {
	return _redisClient
}

// GetMgoSession .
func GetMgoSession() *mgo.Session {
	return _mgoSession
}

type wholeRepo struct {
	FooRepo repository.FooRepo
}

// InitRepos .
func InitRepos(cfg *types.Config) error {
	_mysqlDB, err := gormic.ConnectSqlite3(cfg.Sqlite3)
	if err != nil {
		logger.Log.Errorf("global.InitRepos failed to connect sqlDB, err=%v", err)
		return err
	}

	Repos = wholeRepo{
		FooRepo: repository.NewFooRepo(_mysqlDB),
	}

	return nil
}
