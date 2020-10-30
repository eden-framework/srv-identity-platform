package common

import "github.com/eden-framework/srv-identity-platform/internal/constants/enums"

type Provider interface {
	ProviderID() enums.BindType
	GetUserID(token string) (string, error)
	GetUserInfo(userID string) (user UserInfo, err error)
}

type UserInfo struct {
	UserID string `json:"userID"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}
