package api

// CheckLoginJSON 检测登录返回的 `JSON`
type CheckLoginJSON struct {
	Code   int64       `json:"code"` // 103 表示未登录
	Msg    string      `json:"msg"`
	Status int64       `json:"status"` // 1 `key有效` | 0 `无效`
	Title  interface{} `json:"title"`
}
