package fs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/huluxia-label/huluxia_upup/xface"
	"github.com/huluxia-label/huluxia_upup/xpath"
)

// GetHuluxiaConfig 获取配置
func GetHuluxiaConfig() (xface.HuluxiaConfigJSON, error) {
	readFile, err := xpath.GetConfigFile()
	var r xface.HuluxiaConfigJSON
	if err != nil {
		return r, err
	}
	file, err := os.Open(readFile)
	if err != nil {
		return r, errors.New("读取配置文件失败")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&r)
	if err != nil {
		return r, errors.New("配置文件格式不对, 请参照示例格式来")
	}
	return r, nil
}

// ChangeConfigToken 修改配置文件的 `token`
func ChangeConfigToken(token string) error {
	data, err := GetHuluxiaConfig()
	if err != nil {
		return err
	}
	output, err := xpath.GetConfigFile()
	if err != nil {
		return err
	}
	var rx = []byte(fmt.Sprintf("{\n  \"username\": \"%v\",\n  \"password\": \"%v\",\n  \"token\": \"%v\"\n}", data.Username, data.Password, token))
	err = ioutil.WriteFile(output, rx, 0777)
	if err != nil {
		return errors.New("写入文件失败")
	}
	return nil
}

// GetConfigToken 获取配置文件的 `token`
// 已废弃
func GetConfigToken() string {
	return ""
}
