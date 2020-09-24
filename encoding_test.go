package sutil

import (
	"testing"
)

func TestTMd5(t *testing.T) {
	t.Log(GetMd5String("1"))
}
