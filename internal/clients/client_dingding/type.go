package client_dingding

type AuthRequest struct {
	AccessKey string `name:"accessKey" in:"query" default:""`
	Timestamp uint64 `name:"timestamp,string" in:"query" default:""`
	Signature string `name:"signature" in:"query" default:""`
}

type CommonReturn struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type GetUserInfoByCodeBody struct {
	AuthCode string `json:"tmp_auth_code"`
}

type GetUserInfoByCodeRequest struct {
	AuthRequest
	Body GetUserInfoByCodeBody `in:"body"`
}

type GetUserInfoByCodeResponse struct {
	CommonReturn
	UserInfo UserInfo `json:"user_info"`
}

type UserInfo struct {
	Nick    string `json:"nick"`
	OpenID  string `json:"openid"`
	UnionID string `json:"unionid"`
}

type GetAccessTokenRequest struct {
	AppKey    string `name:"appkey" in:"query"`
	AppSecret string `name:"appsecret" in:"query"`
}

type GetAccessTokenResponse struct {
	CommonReturn
	AccessToken string `json:"access_token"`
	ExpireIn    int    `json:"expires_in,omitempty"`
}

type AccessRequest struct {
	AccessToken string `name:"access_token" in:"query" default:""`
}

type GetUserIDByUnionIDRequest struct {
	AccessRequest
	UnionID string `name:"unionid" in:"query"`
}

type GetUserIDByUnionIDResponse struct {
	CommonReturn
	ContactType int    `json:"contactType,omitempty"`
	UserID      string `json:"userid"`
}

type GetUserInfoDetailRequest struct {
	AccessRequest
	UserID string `name:"userid" in:"query"`
	Lang   string `name:"lang" in:"query" default:"zh_CN"`
}

type GetUserInfoDetailResponse struct {
	CommonReturn
	UserID          string                 `json:"userid"`
	UnionID         string                 `json:"unionid"`
	Name            string                 `json:"name"`
	Tel             string                 `json:"tel,omitempty"`
	WorkPlace       string                 `json:"workPlace,omitempty"`
	Remark          string                 `json:"remark,omitempty"`
	Mobile          string                 `json:"mobile,omitempty"`
	Email           string                 `json:"email,omitempty"`
	OrgEmail        string                 `json:"orgEmail,omitempty"`
	Active          bool                   `json:"active,omitempty"`
	OrderInDepts    string                 `json:"orderInDepts,omitempty"`
	IsAdmin         bool                   `json:"isAdmin,omitempty"`
	IsBoss          bool                   `json:"isBoss,omitempty"`
	IsLeaderInDepts string                 `json:"isLeaderInDepts,omitempty"`
	IsHide          bool                   `json:"isHide,omitempty"`
	Department      []uint64               `json:"department,omitempty"`
	Position        string                 `json:"position,omitempty"`
	Avatar          string                 `json:"avatar,omitempty"`
	HiredDate       uint64                 `json:"hiredDate,omitempty"`
	JobNumber       string                 `json:"jobNumber,omitempty"`
	ExtAttr         map[string]interface{} `json:"extattr,omitempty"`
	IsSenior        bool                   `json:"isSenior,omitempty"`
	StateCode       string                 `json:"stateCode,omitempty"`
	RealAuthed      bool                   `json:"realAuthed,omitempty"`
	Roles           []UserRole             `json:"roles,omitempty"`
}

type UserRole struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
}
