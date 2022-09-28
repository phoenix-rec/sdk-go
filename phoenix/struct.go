package phoenix

const (
	defaultRecUrl   = "http://phoenix-api.icocofun.com/recommend/openapi/rec/"
	defaultWriteUrl = "http://phoenix-api.icocofun.com/data/openapi/write/"
)

type DataResp struct {
	Ret     int    `json:"ret"`
	Errcode int    `json:"errcode"`
	Msg     string `json:"msg"`
}

type RecommendReq struct {
	CustomID  int64    `json:"custom_id"`
	ProjectID int64    `json:"project_id"`
	SceneID   int64    `json:"scene_id"`
	User      UserInfo `json:"user"`
}

type UserInfo struct {
	UserID  string `json:"user_id"`
	Did     string `json:"did"`
	Ip      string `json:"ip"`
	Channel string `json:"channel"`
	Network int32  `json:"network"` // 枚举：1-WIFI，2-2G，3-3G，4-4G，5-5G，0-未知
	Os      string `json:"os"`      //枚举：ios | android
	Model   string `json:"model"`
}

type RecResp struct {
	Ret     int           `json:"ret"`
	Errcode int           `json:"errcode"`
	Msg     string        `json:"msg"`
	Data    RecommendResp `json:"data"`
}

type RecommendResp struct {
	List []RecommendItemSt `json:"list"`

	ItemInfos []ItemInfoSt `json:"items"`
	UserInfo  UserInfoSt   `json:"user_info"`

	ExpInfo ExpInfoSt `json:"exp_info"`
}

type RecommendItemSt struct {
	ItemID string `json:"item_id"`
}

type ItemInfoSt struct {
	ItemID    string `json:"item_id"`
	AuthorID  string `json:"author_id"`
	Gid       int64  `json:"gid"`
	MediaType int64  `json:"media_type"`
}

type UserInfoSt struct {
	HistoryLength int64 `json:"history_length"`
}

type ExpInfoSt struct {
	ExpTag string `json:"exp_tag"`
}
