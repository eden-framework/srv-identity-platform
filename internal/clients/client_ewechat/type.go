package client_ewechat

type CommonResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type CommonRequest struct {
	AccessToken string `name:"access_token" in:"query"`
}

type GetAccessTokenRequest struct {
	CorpID     string `name:"corpid" in:"query"`
	CorpSecret string `name:"corpsecret" in:"query"`
}

type GetAccessTokenResponse struct {
	CommonResponse
	AccessToken string `json:"access_token"`
	ExpireIn    int    `json:"expires_in,omitempty"`
}

type GetUserInfoByCodeRequest struct {
	CommonRequest
	Code string `name:"code" in:"query"`
}

type GetUserInfoByCodeResponse struct {
	CommonResponse
	UserID string `json:"UserId,omitempty"`
	OpenID string `json:"OpenId,omitempty"`
}

type GetUserInfoDetailRequest struct {
	CommonRequest
	UserID string `name:"userid" in:"query"`
}

type GetUserInfoDetailResponse struct {
	CommonResponse
	UserID               string                 `json:"userid"`
	Name                 string                 `json:"name,omitempty"`
	Mobile               string                 `json:"mobile,omitempty"`
	Department           []uint64               `json:"department,omitempty"`
	Order                []int                  `json:"order,omitempty"`
	Position             string                 `json:"position,omitempty"`
	Gender               string                 `json:"gender,omitempty"`
	Email                string                 `json:"email,omitempty"`
	IsLeaderInDepartment []int                  `json:"is_leader_in_dept,omitempty"`
	Avatar               string                 `json:"avatar,omitempty"`
	ThumbAvatar          string                 `json:"thumb_avatar,omitempty"`
	Telephone            string                 `json:"telephone,omitempty"`
	Alias                string                 `json:"alias,omitempty"`
	ExtAttr              map[string]interface{} `json:"extattr,omitempty"`
	Status               int                    `json:"status,omitempty"`
	QrCode               string                 `json:"qr_code,omitempty"`
	ExternalProfile      map[string]interface{} `json:"external_profile,omitempty"`
	ExternalPosition     string                 `json:"external_position,omitempty"`
	Address              string                 `json:"address,omitempty"`
	OpenUserID           string                 `json:"open_userid,omitempty"`
	MainDepartment       uint64                 `json:"main_department,omitempty"`
}
