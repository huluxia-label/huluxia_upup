package cli

import (
	"github.com/fatih/color"
	"github.com/huluxia-label/huluxia_upup/api"
)

// Help `cli` 提示
func Help() string {
	return `
葫芦侠签到工具
	huluxia_upup login # 登录
	huluxia_upup check # 检测key是否过期
	huluxia_upup up    # 一键签到
`
}

// SayHiLoginUser 登录用户提示
func SayHiLoginUser(x api.LoginJSON) {
	color.Cyan("登录成功, 账号昵称: %s", x.User.Nick)
}
