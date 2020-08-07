package xface

// HuluxiaConfigJSON 配置
type HuluxiaConfigJSON struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Token    string `json:"token"`    // 登录之后的 `key`
}
