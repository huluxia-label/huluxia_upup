package middleware

import "github.com/huluxia-label/huluxia_upup/api"

// GetCatSigninStatus 获取签到状态
func GetCatSigninStatus(x api.CatSigninJSON) bool {
	var f = x.Status
	if f == 1 {
		return true
	} else if f == 0 {
		return false
	}
	return false
}
