package main

import (
	log2 "github.com/labstack/gommon/log"
	"testing"
)

func TestLog(t *testing.T) {
	log := log2.New("")
	log.SetHeader("${time_rfc3339}  ${level} --- ${short_file} ${line} ${message}")
	log.EnableColor()
	log.Info("system")

}
