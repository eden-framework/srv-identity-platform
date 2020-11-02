package users

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
)

type UserController struct {
	db sqlx.DBExecutor
}

func NewUserController(db sqlx.DBExecutor) *UserController {
	return &UserController{
		db: db,
	}
}

func (c *UserController) GetUserByMobile(mobile string) (user *databases.Users, err error) {
	user = &databases.Users{
		Mobile: mobile,
	}
	err = user.FetchByMobile(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.UserNotFound.StatusError().WithMsg("根据手机号没有找到用户，请核对手机号码")
			return
		}
		return nil, err
	}

	return
}

func (c *UserController) GetUserByBindID(typ enums.BindType, bindID string) (user *databases.Users, err error) {
	bind := databases.UserBinds{
		Type:   typ,
		BindID: bindID,
	}
	err = bind.FetchByTypeAndBindID(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.UserBindNotFound
			return
		}
		return
	}

	user = &databases.Users{
		UserID: bind.UserID,
	}
	err = user.FetchByUserID(c.db)
	if err != nil {
		return
	}

	return
}

func (c *UserController) CreateBind(userID uint64, bingID string, bindType enums.BindType) (*databases.UserBinds, error) {
	bind := &databases.UserBinds{
		UserID: userID,
		BindID: bingID,
		Type:   bindType,
	}
	err := bind.Create(c.db)
	return bind, err
}

func (c *UserController) CreateUserAndBind(userID uint64, bingID string, bindType enums.BindType, opts ...CreateUserOpt) (*databases.Users, *databases.UserBinds, error) {
	tx := sqlx.NewTasks(c.db)
	user := &databases.Users{}
	for _, opt := range opts {
		opt(user)
	}
	WithUserID(userID)(user)
	tx = tx.With(user.Create)

	bind := &databases.UserBinds{
		UserID: userID,
		BindID: bingID,
		Type:   bindType,
	}
	tx = tx.With(bind.Create)

	err := tx.Do()
	return user, bind, err
}
