package main

import (
	"fmt"
	"github.com/konglingyinxia/win-start-client/server"
	"github.com/konglingyinxia/win-start-client/server/api"
	"github.com/konglingyinxia/win-start-client/server/dto/req"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/model"
	"github.com/konglingyinxia/win-start-client/server/utils/env"
	"github.com/konglingyinxia/win-start-client/server/utils/id"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

func TestEnv(t *testing.T) {
	server.Start()
	env := model.Env{
		EnvName: "dd",
		EnvDir:  "dd",
		Remark:  "{}",
		EnvVars: "{}",
		BaseModel: model.BaseModel{
			ID: id.UuidSimple(),
		},
	}
	fmt.Println(env)
	result := global.DB.Create(&env)
	fmt.Println(env.ID)
	fmt.Println(result.Error)
}
func TestDeleteEnv(t *testing.T) {
	server.Start()
	err := global.DB.Delete(&model.Env{BaseModel: model.BaseModel{ID: "783b59ba4ce443aea2fae742edc3008f"}}).Error
	fmt.Println(err)
}
func TestCheckEnv(t *testing.T) {
	server.Start()
	result := api.NewEnvManager().CheckEnvVars(req.EnvCheckReq{
		ID:      "edb245dae83b4532a7ec120e5d0dd967",
		Command: "java",
	})
	fmt.Println(result.Data)

}

func TestEnvList(t *testing.T) {
	server.Start()
	//环境变量转化为map
	envMap := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2) //分割环境变量名和值
		envMap[pair[0]] = pair[1]
	}
	envMap["JAVA_HOME"] = "/path/to/java"
	envMap["MY_PATH"] = "$ENV_HOME:$JAVA_HOME/bin"
	envMap["PATH_WITH_MY_PATH"] = "%MY_PATH%:%JAVA_HOME%/lib"
	// 原始包含环境变量引用的字符串
	input := "Using JAVA_HOME: $JAVA_HOME; MY_PATH: %MY_PATH%; PATH_WITH_MY_PATH: %PATH_WITH_MY_PATH%"
	// 准备一个 map 记录已解析的环境变量
	parsed := make(map[string]string)
	// 解析引用
	parsedResult, err := env.ParseEnvVars(input, parsed, envMap)
	if err != nil {
		fmt.Println("解析出错:", err)
		return
	}
	// 打印解析结果
	fmt.Println(parsedResult)

}

// 解析并替换环境变量引用
func ParseEnvVars(input string) (string, error) {
	// 正则表达式匹配 $VAR 或 ${VAR}
	re := regexp.MustCompile(`\$(\w+)|%(\w+)%`)
	// 替换函数
	result := re.ReplaceAllStringFunc(input, func(match string) string {
		var varName string
		if match[0] == '$' {
			// 处理 $VAR 和 ${VAR}
			varName = match[1:] // 删除前面的 $ 和后面的 }（如果存在)
		} else if match[0] == '%' {
			varName = match[1 : len(match)-1]
		} else {
			varName = match[1:] // 删除 $
		}
		// 获取环境变量值
		if value, exists := os.LookupEnv(varName); exists {
			return value
		}
		return match // 返回原字符串（没有找到环境变量）
	})
	return result, nil
}

func TestProgramStart(t *testing.T) {
	// 设置 JAVA_HOME 环境变量
	javaHome := "/home/kongling/work/work/project/go/win-start-client/opt/env/corretto-1.8.0_412" // 替换为你的 Java 安装路径
	os.Setenv("JAVA_HOME", javaHome)
	// 设置 PATH 环境变量，确保 Java 的 bin 目录在 PATH 中
	//path := os.Getenv("PATH")
	os.Setenv("PATH", fmt.Sprintf("%s/bin", javaHome))
	// 创建命令对象
	cmd := exec.Command("java", "-version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	} // // 执行命令，并获取输出
	fmt.Println(string(output))
}
