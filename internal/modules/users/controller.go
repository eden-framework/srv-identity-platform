package users

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	db sqlx.DBExecutor
}

func NewController(db sqlx.DBExecutor) *Controller {
	return &Controller{
		db: db,
	}
}

func (c *Controller) GetUserByUserID(userID uint64) (user *databases.Users, err error) {
	user = &databases.Users{
		UserID: userID,
	}
	err = user.FetchByUserID(c.db)
	if err != nil {
		if !sqlx.DBErr(err).IsNotFound() {
			logrus.Errorf("[user.Controller.GetUserByUserID] err: %v, userID: %d", err, userID)
		}
	}
	return
}

func (c *Controller) GetUserByMobile(mobile string) (user *databases.Users, err error) {
	user = &databases.Users{
		Mobile: mobile,
	}
	err = user.FetchByMobile(c.db)
	if err != nil {
		if !sqlx.DBErr(err).IsNotFound() {
			logrus.Errorf("[user.Controller.GetUserByUserID] err: %v, mobile: %s", err, mobile)
		}
	}
	return
}

func (c *Controller) GetUserByBindID(typ enums.BindType, bindID string) (user *databases.Users, err error) {
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
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.UserNotFound.StatusError().WithMsg("根据绑定账号没有找到用户")
			return
		}
		return
	}

	return
}

func (c *Controller) CreateBind(userID uint64, bingID string, bindType enums.BindType) (*databases.UserBinds, error) {
	bind := &databases.UserBinds{
		UserID: userID,
		BindID: bingID,
		Type:   bindType,
	}
	err := bind.Create(c.db)
	return bind, err
}

func (c *Controller) CreateUserAndBind(userID uint64, bingID string, bindType enums.BindType, opts ...CreateUserOpt) (*databases.Users, *databases.UserBinds, error) {
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
