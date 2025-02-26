package id

import (
	"github.com/google/uuid"
	"strings"
)

func UuidSimple() string {
	u := uuid.New().String()
	//去除-
	return strings.ReplaceAll(u, "-", "")
}

// SnowFlakeId 雪花算法生成ID
func SnowFlakeId() int64 {
	//todo

	return 0
}
