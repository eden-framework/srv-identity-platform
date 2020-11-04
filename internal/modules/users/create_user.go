package users

import "github.com/eden-framework/srv-identity-platform/internal/databases"

type CreateUserOpt func(model *databases.Users)

func (c *Controller) CreateUser(opts ...CreateUserOpt) error {
	user := &databases.Users{}
	for _, opt := range opts {
		opt(user)
	}

	err := user.Create(c.db)
	return err
}

func WithUserID(userID uint64) CreateUserOpt {
	return func(model *databases.Users) {
		model.UserID = userID
	}
}

func WithUserName(userName string) CreateUserOpt {
	return func(model *databases.Users) {
		model.UserName = userName
	}
}

func WithPassword(password, salt string) CreateUserOpt {
	return func(model *databases.Users) {
		model.Password = password
		model.Salt = salt
	}
}

func WithMobile(mobile string) CreateUserOpt {
	return func(model *databases.Users) {
		model.Mobile = mobile
	}
}

func WithEmail(email string) CreateUserOpt {
	return func(model *databases.Users) {
		model.Email = email
	}
}
