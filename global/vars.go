package global

import (
	"github.com/yeqown/infrastructure/types"
)

// Config .
type Config struct {
	Mysql   *types.MysqlC        `json:"mysql"`
	Sqlite3 *types.SQLite3Config `json:"sqlite3"`
}

var (
	_cfg    = new(Config)
	_curEnv types.Envrion
)

// SetConfig .
func SetConfig(cfg *Config) {
	_cfg = cfg
}

// GetConfig .
func GetConfig() *Config {
	return _cfg
}

// SetEnv .
func SetEnv(env types.Envrion) {
	_curEnv = env
}

// GetEnv .
func GetEnv() types.Envrion {
	return _curEnv
}
