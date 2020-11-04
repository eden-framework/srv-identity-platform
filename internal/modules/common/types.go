package common

import "github.com/eden-framework/srv-identity-platform/internal/constants/enums"

type Provider interface {
	ProviderID() enums.BindType
	ProviderName() string
	ProviderIcon() string
	ProviderEntry() string
	GetUserID(token string) (string, error)
	GetUserInfo(userID string) (user UserInfo, err error)
}

type ProviderInfo struct {
	ID    enums.BindType `json:"id"`
	Name  string         `json:"name"`
	Icon  string         `json:"icon"`
	Entry string         `json:"entry"`
}

type UserInfo struct {
	UserID string `json:"userID"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}
