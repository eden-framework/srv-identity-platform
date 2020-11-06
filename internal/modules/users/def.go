package users

import "github.com/eden-framework/srv-identity-platform/internal/constants/enums"

type UserIDWithUserType struct {
	UserID   uint64         `json:"userID"`
	UserType enums.UserType `json:"userType"`
}
