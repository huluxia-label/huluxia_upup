package api

// CheckLoginJSON 检测登录返回的 `JSON`
type CheckLoginJSON struct {
	Code   int64       `json:"code"` // 103 表示未登录
	Msg    string      `json:"msg"`
	Status int64       `json:"status"` // 1 `key有效` | 0 `无效`
	Title  interface{} `json:"title"`
}

// LoginJSON 登录数据
type LoginJSON struct {
	Msg  string `json:"msg"`
	Key  string `json:"_key"`
	User struct {
		UserID          int         `json:"userID"`
		Role            int         `json:"role"`
		Nick            string      `json:"nick"`
		Avatar          string      `json:"avatar"`
		Birthday        int64       `json:"birthday"`
		Age             int         `json:"age"`
		Gender          int         `json:"gender"`
		Level           int         `json:"level"`
		Isgold          int         `json:"isgold"`
		IdentityTitle   interface{} `json:"identityTitle"`
		IdentityColor   int         `json:"identityColor"`
		NeedSetPassword int         `json:"needSetPassword"`
		NeedSetUserInfo int         `json:"needSetUserInfo"`
	} `json:"user"`
	SessionKey string `json:"session_key"`
	Status     int    `json:"status"`
}

// CategoryJSON 分类数据
type CategoryJSON struct {
	Msg      string `json:"msg"`
	PostInfo struct {
		PostID       int      `json:"postID"`
		Title        string   `json:"title"`
		Detail       string   `json:"detail"`
		Images       []string `json:"images"`
		Score        int      `json:"score"`
		ScoreTxt     string   `json:"scoreTxt"`
		Hit          int      `json:"hit"`
		CommentCount int      `json:"commentCount"`
		Notice       int      `json:"notice"`
		Weight       int      `json:"weight"`
		IsGood       int      `json:"isGood"`
		CreateTime   int64    `json:"createTime"`
		ActiveTime   int64    `json:"activeTime"`
		Line         int      `json:"line"`
		User         struct {
			UserID        int         `json:"userID"`
			Nick          string      `json:"nick"`
			Avatar        string      `json:"avatar"`
			Gender        int         `json:"gender"`
			Age           int         `json:"age"`
			Role          int         `json:"role"`
			Experience    int         `json:"experience"`
			Credits       int         `json:"credits"`
			IdentityTitle interface{} `json:"identityTitle"`
			IdentityColor int         `json:"identityColor"`
			Level         int         `json:"level"`
			LevelColor    int         `json:"levelColor"`
			Integral      int         `json:"integral"`
			UUID          int         `json:"uuid"`
			IntegralNick  interface{} `json:"integralNick"`
			Signature     interface{} `json:"signature"`
			MedalList     interface{} `json:"medalList"`
		} `json:"user"`
		Ext            interface{} `json:"ext"`
		Category       interface{} `json:"category"`
		Tagid          int         `json:"tagid"`
		Status         int         `json:"status"`
		Praise         int         `json:"praise"`
		Voice          string      `json:"voice"`
		IsAuthention   int         `json:"isAuthention"`
		IsRich         int         `json:"isRich"`
		AppOrientation int         `json:"appOrientation"`
		IsAppPost      int         `json:"isAppPost"`
		AppVersion     interface{} `json:"appVersion"`
		AppSize        float64     `json:"appSize"`
		AppSystem      interface{} `json:"appSystem"`
		AppLogo        interface{} `json:"appLogo"`
		Screenshots    interface{} `json:"screenshots"`
		AppIntroduce   interface{} `json:"appIntroduce"`
		AppURL         interface{} `json:"appUrl"`
		AppLanguage    interface{} `json:"appLanguage"`
		IsGif          int         `json:"isGif"`
	} `json:"postInfo"`
	Categories []struct {
		CategoryID        int64         `json:"categoryID"`
		Model             int           `json:"model"`
		Title             string        `json:"title"`
		Icon              string        `json:"icon"`
		Description       interface{}   `json:"description"`
		PostCount         int           `json:"postCount"`
		ViewCount         int           `json:"viewCount"`
		PostCountFormated interface{}   `json:"postCountFormated"`
		ViewCountFormated interface{}   `json:"viewCountFormated"`
		IsGood            int           `json:"isGood"`
		IsSubscribe       int           `json:"isSubscribe"`
		Hide              int           `json:"hide"`
		Seq               int           `json:"seq"`
		SubscribeType     int           `json:"subscribeType"`
		Forumid           int           `json:"forumid"`
		CateRule          interface{}   `json:"cateRule"`
		Type              string        `json:"type"`
		Wap               interface{}   `json:"wap"`
		AppID             int           `json:"appId"`
		Moderator         []interface{} `json:"moderator"`
		Tags              []interface{} `json:"tags"`
		IsSearch          string        `json:"isSearch"`
		HasChild          int           `json:"hasChild"`
		Pid               int           `json:"pid"`
	} `json:"categories"`
	More   int    `json:"more"`
	Start  string `json:"start"`
	Status int    `json:"status"`
}

// CatSigninJSON 签到接口数据
type CatSigninJSON struct {
	Msg                            string      `json:"msg"`
	ContinueDays                   int         `json:"continueDays"`
	ExperienceVal                  int         `json:"experienceVal"`
	RemainDaysToGetMoreExperiences int         `json:"remainDaysToGetMoreExperiences"`
	NextExperience                 int         `json:"nextExperience"`
	Signin                         int         `json:"signin"`
	IsFirstSignToday               int         `json:"isFirstSignToday"`
	Medal                          interface{} `json:"medal"`
	MedalURL                       interface{} `json:"medalUrl"`
	Ranking                        int         `json:"ranking"`
	Contrast                       int         `json:"contrast"`
	Status                         int         `json:"status"`
}
