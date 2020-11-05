package databases

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/timex"
	"strconv"
	"time"
)

//go:generate eden generate model Token --database Config.DB --with-comments
//go:generate eden generate tag Token --defaults=true
// @def primary ID
// @def unique_index U_token_id TokenID
type Token struct {
	datatypes.PrimaryID
	// 业务ID
	TokenID uint64 `json:"tokenID,string" db:"f_token_id"`
	TokenScope
	Comment string `json:"comment" db:"f_comment,default=''"`
	// 签发时间
	IssuedAt datatypes.MySQLTimestamp `json:"issuedAt" db:"f_issued_at,default='0'"`
	// 过期时间
	ExpireAt datatypes.MySQLTimestamp `json:"expireAt" db:"f_expire_at,default='0'"`
}

func (m *Token) Issue(expireIn time.Duration) {
	now := timex.Now()
	m.IssuedAt = datatypes.MySQLTimestamp(now)
	m.ExpireAt = datatypes.MySQLTimestamp(now.Add(expireIn))
}

func (m *Token) ToStandardClaims() *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Audience:  m.Audience,
		ExpiresAt: time.Time(m.ExpireAt).Unix(),
		Id:        fmt.Sprintf("%d", m.TokenID),
		IssuedAt:  time.Time(m.IssuedAt).Unix(),
		Issuer:    m.Issuer,
		Subject:   m.Subject.String(),
	}
}

func ParseTokenFromStandardClaims(c *jwt.StandardClaims) (*Token, error) {
	sub, err := enums.ParseTokenSubjectFromString(c.Subject)
	if err != nil {
		return nil, err
	}

	token := &Token{
		TokenScope: TokenScope{
			Issuer:   c.Issuer,
			Subject:  sub,
			Audience: c.Audience,
		},
	}

	id, err := strconv.ParseUint(c.Id, 10, 64)
	if err != nil {
		return nil, err
	}

	token.TokenID = id
	token.IssuedAt = datatypes.MySQLTimestamp(time.Unix(c.IssuedAt, 0))
	token.ExpireAt = datatypes.MySQLTimestamp(time.Unix(c.ExpiresAt, 0))

	return token, nil
}

type TokenScope struct {
	// 签发方
	Issuer string `db:"f_issuer" json:"issuer"`
	// 用途
	Subject enums.TokenSubject `db:"f_subject" json:"subject"`
	// 使用方
	Audience string `db:"f_audience" json:"audience"`
}
