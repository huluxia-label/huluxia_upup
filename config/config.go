// create by d1y<chenhonzhou@gmail.com>

package config

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/huluxia-label/huluxia_upup/fs"
)

// BaseAPI 基础
var BaseAPI = `http://floor.huluxia.com`

// CheckLoginAPI 检测登录状态
var CheckLoginAPI = fmt.Sprintf(`%v/user/status/ANDROID/2.1`, BaseAPI)

// LoginAPI 登录
var LoginAPI = fmt.Sprintf(`%v/account/login/ANDROID/4.0`, BaseAPI)

// CategoryAPI 分类
var CategoryAPI = fmt.Sprintf(`%v/category/list/ANDROID/2.0`, BaseAPI)

// SigninAPI 签到
var SigninAPI = fmt.Sprintf(`%v/user/signin/ANDROID/4.0`, BaseAPI)

// IgnoreCategoryID 分类忽略的 `id`
var IgnoreCategoryID = []int64{
	0, // 系统推荐之类的, 这个 `id` 不可以签到
}

// Username 用户名
var Username = ""

// Password 密码
var Password = ""

// Token 登录秘钥
var Token = ""

func init() {
	ctx, err := fs.GetHuluxiaConfig()
	if err != nil {
		color.Red(err.Error())
		os.Exit(2)
	}
	var username = ctx.Username
	var password = ctx.Password
	var token = ctx.Token
	if len(username) > 1 && len(password) > 1 {
		Username = username
		Password = password
	} else {
		color.Red("配置文件加载错误")
		os.Exit(2)
	}
	if len(token) > 10 {
		Token = token
	}
}
