package types

import (
	types2 "github.com/yeqown/infrastructure/types"
)

// Config .
type Config struct {
	Mysql   *types2.MysqlC        `json:"mysql"`
	Sqlite3 *types2.SQLite3Config `json:"sqlite3"`
}

var (
	_cfg    = new(Config)
	_curEnv types2.Envrion
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
func SetEnv(env types2.Envrion) {
	_curEnv = env
}

// GetEnv .
func GetEnv() types2.Envrion {
	return _curEnv
}
