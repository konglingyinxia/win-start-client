package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"testing"
)

func Test_trim(t *testing.T) {
	cmds := "./xjar.exe java -server -jar  project-platform-item-1.0-SNAPSHOT.jar"
	//空格分割Command
	commands := strutil.SplitAndTrim(cmds, " ")
	for i, command := range commands {
		fmt.Println("index:", i, " item:", command)
	}
}
