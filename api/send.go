package api

import "errors"

// Login 登录
func Login() (string, error) {
	return "", errors.New("登录失败, 可能是网络问题")
}

// CheckLogin 检测 `key` 是否过期
func CheckLogin() bool {
	return false
}
