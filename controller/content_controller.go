package controller

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// ScriptInfo 脚本文件注释信息
type ScriptInfo struct {
	HostLocation string   `json:"HostLocation"`
	HostIP       []string `json:"HostIP"`
	FileInfo     struct {
		Path  string `json:"Path"`
		Owner string `json:"Owner"`
		Group string `json:"Group"`
		Perm  string `json:"Perm"`
	} `json:"FileInfo"`
	CrontabEnable bool `json:"CrontabEnable"`
	CrontabData   struct {
		Time    string `json:"Time"`
		Command string `json:"command"`
		Arg     string `json:"arg"`
	} `json:"CrontabData"`
	Language    string `json:"Language"`
	Author      string `json:"Author"`
	Description string `json:"Description"`
}

// ExtractContent 从原始内容中提取脚本信息以及脚本内容
func ExtractContent(originalContent *string) (*ScriptInfo, *string, error) {
	// 使用正则表达式匹配注释块
	re := regexp.MustCompile(`(?s): <<COMMENT_BLOCK([\s\S]*?)COMMENT_BLOCK`)
	scriptInfoContent := re.FindStringSubmatch(*originalContent)
	scriptContent := re.ReplaceAllString(*originalContent, "")
	if len(scriptInfoContent) != 2 {
		return nil, nil, fmt.Errorf("no comment block found")
	}

	// 解析JSON到Config结构体
	var scriptInfo ScriptInfo
	err := json.Unmarshal([]byte(scriptInfoContent[1]), &scriptInfo)
	if err != nil {
		return nil, nil, err
	}
	return &scriptInfo, &scriptContent, nil
}
