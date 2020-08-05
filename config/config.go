package config

import (
	"fmt"
)

// BaseAPI 基础
var BaseAPI = `http://floor.huluxia.com`

// CheckLoginAPI 检测登录状态
var CheckLoginAPI = fmt.Sprintf(`%v/user/position/ANDROID/2.1`, BaseAPI)

// LoginAPI 登录
var LoginAPI = fmt.Sprintf(`%v/account/login/ANDROID/4.0`, BaseAPI)
