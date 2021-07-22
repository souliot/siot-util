package sutil

import (
	"encoding/json"
	"testing"
	"time"
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
	t.Log(To_md5(time.Now().Format(FormatTime)))
	t.Log(To_md5(time.Now().Format(FormatDate)))
	t.Log(To_md5(time.Now().Format(FormatDateTime)))
	t.Log(To_md5("测试" + time.Now().Format(FormatDateTime)))
	t.Log(To_md5(time.Now().Format(FormatDateTime)))
	t.Log(To_md5(time.Now().Format(FormatDateTime))[0:16])
}

func TestOther(t *testing.T) {
	u := &User{}
	err := json.Unmarshal([]byte{}, u)
	t.Log(err)
}
