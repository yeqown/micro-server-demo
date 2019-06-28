package global_test

import (
	"testing"

	"github.com/yeqown/micro-server-demo/global"
)

func Test_Config(t *testing.T) {
	c := new(global.Config)
	c.Mysql = nil
	global.SetConfig(c)
}
