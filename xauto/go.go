package xauto

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/huluxia-label/huluxia_upup/api"
	"github.com/huluxia-label/huluxia_upup/cli"
	"github.com/huluxia-label/huluxia_upup/config"
	"github.com/huluxia-label/huluxia_upup/fs"
	"github.com/huluxia-label/huluxia_upup/utils"
	"github.com/huluxia-label/huluxia_upup/utils/middleware"
)

// HuluxiaAuto 自动化流程
func HuluxiaAuto() {
	var token = ""
	log.Println(fmt.Sprintf("账号: %v", config.Username))
	if len(config.Token) > 8 {
		token = config.Token
		log.Println("`token`已存在, 已登录")
	} else {
		token = waifuLogin(config.Username, config.Password)
	}
	log.Println("正在验证 `token` 是否过期")
	var isLoginActive = api.CheckLogin(token)
	if isLoginActive != nil {
		log.Println(isLoginActive)
		var confirmLogin = utils.AskForConfirmation("是否重新登录?")
		if confirmLogin {
			token = waifuLogin(config.Username, config.Password)
		} else {
			log.Fatalln("已取消登录")
		}
	}
	d, err := api.GetCategoryList()
	if err != nil {
		log.Fatalln("获取板块列表失败")
	}
	waifuCatEach(d, token)
}

// 登录
func waifuLogin(u string, p string) string {
	x, err := api.Login(u, p)
	if err != nil {
		color.Red(err.Error())
		os.Exit(2)
	}
	var t = x.Key
	fs.ChangeConfigToken(t)
	cli.SayHiLoginUser(x)
	return t
}

// 板块签到
func waifuCatEach(u api.CategoryJSON, token string) {
	for _, item := range u.Categories {
		var title = item.Title
		var id = item.CategoryID
		var ig bool = false
		for _, igID := range config.IgnoreCategoryID {
			if igID == id {
				ig = true
			}
		}
		if !ig {
			ctx, err := api.HandleCatSignin(token, id)
			var isSuccess bool = false
			if err != nil {
				isSuccess = false
			}
			isSuccess = middleware.GetCatSigninStatus(ctx)
			var statusText = "失败"
			if isSuccess {
				statusText = "成功"
			}
			log.Println("开始签到板块: ", title, "签到状态: ", statusText)
		}
	}
	log.Println("签到完成")
}
