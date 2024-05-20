package controller

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// ScriptInfo 脚本文件注释信息
type ScriptInfo struct {
	HostLocation  string   `json:"HostLocation"`
	HostIP        []string `json:"HostIP"`
	HostPath      string   `json:"HostPath"`
	HostUser      string   `json:"HostUser"`
	CrontabEnable bool     `json:"CrontabEnable"`
	CrontabData   struct {
		Time    string `json:"Time"`
		Command string `json:"command"`
		Arg     string `json:"arg"`
	} `json:"CrontabData"`
	Language    string `json:"Language"`
	Author      string `json:"Authorer"`
	Description string `json:"Description"`
}

// ExtractCommentBlock 解析并返回注释块中的结构体
func ExtractCommentBlock(content *string) (*ScriptInfo, error) {
	// 使用正则表达式匹配注释块
	re := regexp.MustCompile(`(?s): <<COMMENT_BLOCK([\s\S]*?)COMMENT_BLOCK`)
	match := re.FindStringSubmatch(*content)
	if len(match) != 2 {
		return nil, fmt.Errorf("no comment block found")
	}

	// 解析JSON到Config结构体
	var info ScriptInfo
	err := json.Unmarshal([]byte(match[1]), &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
