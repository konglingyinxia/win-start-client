package env

import (
	"fmt"
	"github.com/konglingyinxia/win-start-client/server/global"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Marshal(envMap map[string]string) (string, error) {
	lines := MarshalArray(envMap)
	return strings.Join(lines, "\n"), nil
}
func MarshalArray(envMap map[string]string) []string {
	lines := make([]string, 0, len(envMap))
	for k, v := range envMap {
		if d, err := strconv.Atoi(v); err == nil {
			lines = append(lines, fmt.Sprintf(`%s=%d`, k, d))
		} else {
			lines = append(lines, fmt.Sprintf(`%s=%s`, k, v))
		}
	}
	sort.Strings(lines)
	return lines
}

// ParseEnvVars 在确定环境变量列表中解析并替换环境变量引用
// 环境变量引用格式：$VAR、${VAR}和%VAR%
// 其中 $VAR 和 ${VAR} 表示变量名，%VAR% 表示变量名，不区分大小写
// 变量名必须符合字母、数字和下划线的组合，且不以数字开头
// param input 输入字符串
// param parsed 已解析的变量列表
// param envMap 环境变量列表
// return 解析并替换后的字符串，以及可能的错误
func ParseEnvVars(input string, parsed map[string]string, envMap map[string]string) (string, error) {
	// 正则表达式匹配 $VAR、${VAR}和%VAR%
	re := regexp.MustCompile(`(\$([A-Z_][A-Z0-9_]*|\{[A-Z_][A-Z0-9_]*\}))|(%([A-Z_][A-Z0-9_]*)%)`)
	var replaceFunc func(string) string
	replaceFunc = func(match string) string {
		var varName string
		if match[0] == '$' {
			// 处理 $VAR 和 ${VAR}
			if match[1] == '{' {
				varName = match[2 : len(match)-1]
			} else {
				varName = match[1:]
			}
		} else if match[0] == '%' {
			// 处理 %VAR%
			varName = match[1 : len(match)-1]
		}
		// 忽略已解析的变量，避免递归
		if value, exists := parsed[varName]; exists {
			return value
		}
		// 获取环境变量值，且将其添加到已解析列表中
		if value, exists := envMap[varName]; exists {
			// 如果值中还包含其他环境变量引用，递归解析
			resolvedValue, _ := ParseEnvVars(value, parsed, envMap)
			parsed[varName] = resolvedValue
			return resolvedValue
		}
		return match // 返回原字符串（没有找到环境变量）
	}
	// 替换所有匹配的变量引用
	result := re.ReplaceAllStringFunc(input, replaceFunc)
	return result, nil
}

// ParseAllEnv  处理引用解析环境变量
func ParseAllEnv(envMaps map[string]string) map[string]string {
	//复自定义环境变量
	tempMap := make(map[string]string)
	for key, value := range envMaps {
		tempMap[strings.ToUpper(key)] = value
	}
	//系统环境变量
	systemEnv := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2) //分割环境变量名和值
		systemEnv[strings.ToUpper(pair[0])] = pair[1]
	}
	//处理环境变量：自定义环境变量存在同名优先解析处理
	for key, value := range systemEnv {
		if existValue, exist := envMaps[key]; exist {
			parsed := make(map[string]string)
			vars, _ := ParseEnvVars(existValue, parsed, systemEnv)
			global.LOG.Info(fmt.Sprintf("解析同名环境境变量：%s=%s,解析结果：%s", key, existValue, vars))
			systemEnv[key] = vars
			tempMap[key] = vars
		} else {
			tempMap[key] = value
		}
	}
	var resultMap = make(map[string]string)
	parsed := make(map[string]string)
	for key, value := range envMaps {
		vars, err := ParseEnvVars(value, parsed, tempMap)
		global.LOG.Info(fmt.Sprintf("解析环境变量：%s=%s,解析结果：%s,错误信息：%v", key, value, vars, err))
		if err == nil {
			resultMap[strings.ToUpper(key)] = vars
		} else {
			resultMap[strings.ToUpper(key)] = value
		}
	}
	return resultMap
}
