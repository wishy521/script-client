package controller

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"scripts-client/common"
	"strconv"
)

// GetUserIDAndGroupID 获取uid和gid
func GetUserIDAndGroupID(username, groupname string) (int, int, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to lookup user: %v", err)
	}

	g, err := user.LookupGroup(groupname)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to lookup group: %v", err)
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid UID: %v", err)
	}
	gid, err := strconv.Atoi(g.Gid)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid GID: %v", err)
	}
	return uid, gid, nil
}

// WriteContentToFile 整合内容并内容到文件
func WriteContentToFile(scriptInfo *ScriptInfo, scriptContent *string) error {
	var info = *scriptInfo
	var results = false
	// 校验脚本路径是否合法
	for _, script := range common.Conf.Script {
		if script == info.FileInfo.Path {
			results = true
		}
	}
	if !results {
		common.Log.Error("The provided path is inconsistent with the client path")
		return fmt.Errorf("The provided path %s is inconsistent with the client path ", info.FileInfo.Path)
	}

	// 生成文件内容
	commentContent := fmt.Sprintf("#!%s\n# Author: %s\n# Description: %s\n", info.Language, info.Author, info.Description)
	var crontabContent string
	if info.CrontabEnable == true {
		crontabContent = fmt.Sprintf("# crontab: %s %s %s", info.CrontabData.Time, info.CrontabData.Command, info.CrontabData.Arg)
	}
	fileContent := commentContent + crontabContent + *scriptContent

	// string转换为8进制
	perm, err := strconv.ParseUint(info.FileInfo.Perm, 8, 32)
	if err != nil {
		common.Log.Error("parse file permissions failed %s", err)
		return fmt.Errorf("parse file permissions failed %s", err)
	}

	// 写入新文件
	err = ioutil.WriteFile(info.FileInfo.Path, []byte(fileContent), fs.FileMode(perm))
	if err != nil {
		common.Log.Error("write new file failed", err)
		return fmt.Errorf("write new file failed %s", err)
	}
	// 文件属主
	uid, gid, err := GetUserIDAndGroupID(info.FileInfo.Owner, info.FileInfo.Group)
	if err != nil {
		common.Log.Error("get UserID and GroupID failed", err)
		return fmt.Errorf("get UserID and GroupID failed %s", err)
	}
	err = os.Chown(info.FileInfo.Path, uid, gid)
	if err != nil {
		common.Log.Error("set file owner failed", err)
		return fmt.Errorf("set file owner failed: %s", err)
	}
	common.Log.Info(fmt.Sprintf("script file %s created successfully", info.FileInfo.Path))
	return nil
}
