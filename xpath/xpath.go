package xpath

import (
	"errors"
	"os"
	"path"

	"github.com/huluxia-label/huluxia_upup/utils"
)

// GetConfigFile 获取配置文件
func GetConfigFile() (string, error) {
	curr, err := os.Getwd()
	if err != nil {
		return "", errors.New("获取当前目录失败")
	}
	var i18nConfigFile = path.Join(curr, "conf.json")
	var cnConfigFile = path.Join(curr, "配置.json")
	var readFile string
	if !utils.Exists(i18nConfigFile) {
		if utils.Exists(cnConfigFile) {
			readFile = cnConfigFile
		} else {
			return "", errors.New("未发现配置文件")
		}
	} else {
		readFile = i18nConfigFile
	}
	return readFile, nil
}
