package sutil

import (
	"testing"
	"time"

	"github.com/astaxie/beego"
)

const (
	FormatTime     = "15:04:05"
	FormatDate     = "2006-01-02"
	FormatDateTime = "2006-01-02 15:04:05"
)

func TestAppId(t *testing.T) {
	beego.Info(To_md5(time.Now().Format(FormatTime)))
	beego.Info(To_md5(time.Now().Format(FormatDate)))
	beego.Info(To_md5(time.Now().Format(FormatDateTime)))
	beego.Info(To_md5("测试" + time.Now().Format(FormatDateTime)))
	beego.Info(To_md5(time.Now().Format(FormatDateTime)))
	beego.Info(To_md5(time.Now().Format(FormatDateTime))[0:16])
}
