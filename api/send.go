package api

import (
	"errors"
	"fmt"

	"github.com/huluxia-label/huluxia_upup/config"
	"github.com/huluxia-label/huluxia_upup/utils"
	"github.com/imroc/req"
)

// createBaseHeader 创建一个默认的 `headers`
func createBaseHeader() req.QueryParam {
	defaultQS := req.QueryParam{
		"platform":    "2",
		"gkey":        "000000",
		"app_version": "4.0.0.6.2", // 版本号可能会旧
		"versioncode": "20141433",
		"market_id":   "floor_huluxia",
		"device_code": "%5Bw%5D02%3A00%3A00%3A00%3A00%3A00-%5Bi%5D008796755300310", // !
	}
	return defaultQS
}

// Login 登录
func Login(username string, password string) (LoginJSON, error) {
	// req.Debug = true
	var md5 = utils.CreateMd5(password)
	var Fbody = fmt.Sprintf(`account=%v&login_type=%v&password=%v`, username, 2, md5)
	t := req.Header{
		"Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
		"User-Agent":   "okhttp/3.8.1",
	}
	r, err := req.Post(config.LoginAPI, createBaseHeader(), t, Fbody)
	var x LoginJSON
	if err != nil {
		return x, errors.New("登录失败, 可能是网络问题")
	}
	err = r.ToJSON(&x)
	if err != nil {
		return x, errors.New("登录失败, 解析数据错误")
	}
	var status = x.Status == 1
	if !status {
		return x, errors.New("登录失败, 账号密码错误")
	}
	return x, nil
}

// CheckLogin 检测 `key` 是否过期
func CheckLogin(token string) error {
	var tokenHd = req.QueryParam{
		"_key": token,
	}
	r, err := req.Get(config.CheckLoginAPI, createBaseHeader(), tokenHd)
	if err != nil {
		return errors.New("网络请求失败")
	}
	var x CheckLoginJSON
	err = r.ToJSON(&x)
	if err != nil {
		return errors.New("解析数据失败")
	}
	if x.Status == 1 {
		return nil
	} else {
		return errors.New(x.Msg)
	}
}

// GetCategoryList 获取分类
func GetCategoryList() (CategoryJSON, error) {
	var x CategoryJSON
	r, err := req.Get(config.CategoryAPI, createBaseHeader())
	if err != nil {
		return x, errors.New("网络请求失败")
	}
	err = r.ToJSON(&x)
	if err != nil {
		return x, errors.New("解析数据失败")
	}
	return x, nil
}

// HandleCatSignin 板块签到
func HandleCatSignin(token string, id int64) (CatSigninJSON, error) {
	var Xctx = req.QueryParam{
		"_key":   token,
		"cat_id": id,
	}
	var x CatSigninJSON
	r, err := req.Get(config.SigninAPI, createBaseHeader(), Xctx)
	if err != nil {
		return x, errors.New("网络请求失败")
	}
	err = r.ToJSON(&x)
	if err != nil {
		return x, errors.New("解析数据失败")
	}
	return x, nil
}
