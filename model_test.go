package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"testing"
	"time"
)

func TestModel(t *testing.T) {
	tTime := time.Now()
	r := []byte(fmt.Sprintf("%v", tTime.Format("2006-01-02 15:04:05")))
	fmt.Println(string(r))
	tTime, err := time.Parse("2006-01-02 15:04:05", string(r))
	if err != nil {
		log.Error("解析时间错误", err)
		return
	}
	log.Info("解析时间:", tTime)
}
