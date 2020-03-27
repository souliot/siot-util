package sutil

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/astaxie/beego"
)

const (
	FormatTime     = "15:04:05"
	FormatDate     = "2006-01-02"
	FormatDateTime = "2006-01-02 15:04:05"
)

type User struct {
	Name string
}

func TestAppId(t *testing.T) {
	beego.Info(To_md5(time.Now().Format(FormatTime)))
	beego.Info(To_md5(time.Now().Format(FormatDate)))
	beego.Info(To_md5(time.Now().Format(FormatDateTime)))
	beego.Info(To_md5("测试" + time.Now().Format(FormatDateTime)))
	beego.Info(To_md5(time.Now().Format(FormatDateTime)))
	beego.Info(To_md5(time.Now().Format(FormatDateTime))[0:16])
}

func TestOther(t *testing.T) {
	u := &User{}
	err := json.Unmarshal([]byte{}, u)
	beego.Info(err)
}
