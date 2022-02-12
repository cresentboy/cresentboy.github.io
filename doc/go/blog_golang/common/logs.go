package common

import (
	"github.com/beego/beego/v2/core/logs"
)

func init() {
	logs.SetLogger("console")
}
