package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/eden-framework/client"
	"github.com/eden-framework/client/client_srv_id"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/profzone/envconfig"
	"time"
)

type tokenManager struct {
	issuer          string
	db              sqlx.DBExecutor
	clientID        client_srv_id.ClientSrvIDInterface
	privateKey      *RSA
	defaultExpireIn time.Duration
}

func newTokenManager(issuer string, db sqlx.DBExecutor, clientID client_srv_id.ClientSrvIDInterface, privateKey *RSA, defaultExpireIn time.Duration) *tokenManager {
	return &tokenManager{
		issuer:          issuer,
		db:              db,
		clientID:        clientID,
		privateKey:      privateKey,
		defaultExpireIn: defaultExpireIn,
	}
}

func (m *tokenManager) ExchangeAccessToken(subject enums.TokenSubject, audience string) (token AccessToken, err error) {
	accessToken, err := m.NewSignedToken(subject, audience, m.defaultExpireIn)
	if err != nil {
		return
	}

	refreshToken, err := m.NewSignedToken(enums.TOKEN_SUBJECT__REFRESHER, audience, m.defaultExpireIn*2)
	if err != nil {
		return
	}

	token.AccessToken = accessToken.Signed
	token.RefreshToken = refreshToken.Signed
	token.Subject = accessToken.Subject
	token.Audience = audience
	token.ExpireIn = envconfig.Duration(m.defaultExpireIn)
	token.Issuer = m.issuer

	return
}

func (m *tokenManager) NewSignedToken(subject enums.TokenSubject, audience string, expireIn time.Duration) (token SignedToken, err error) {
	id, err := client.GetUniqueID(m.clientID)
	if err != nil {
		return
	}

	tok := &databases.Token{
		TokenID: id,
		TokenScope: databases.TokenScope{
			Issuer:   m.issuer,
			Subject:  subject,
			Audience: audience,
		},
	}
	tok.Issue(expireIn)

	signed, err := jwt.NewWithClaims(jwt.SigningMethodRS256, tok.ToStandardClaims()).SignedString(m.privateKey.PrivateKey())
	if err != nil {
		return
	}

	token.Token = tok
	token.Signed = signed

	return
}

func (m tokenManager) ValidateToken(token string) (*databases.Token, error) {
	tok, err := m.ParseSignedToken(token)
	if err != nil {
		return nil, errors.InternalError.StatusError().WithDesc(err.Error())
	}

	if tok.Issuer != m.issuer {
		return nil, errors.InvalidToken
	}

	if time.Until(time.Time(tok.ExpireAt)) < 0 {
		return nil, errors.TokenExpired
	}

	return tok, nil
}

func (m *tokenManager) ParseSignedToken(token string) (*databases.Token, error) {
	t, err := jwt.ParseWithClaims(token, NewClaimsWithClockDrift(5), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &m.privateKey.PrivateKey().PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return databases.ParseTokenFromStandardClaims(t.Claims.(*ClaimsWithClockDrift).StandardClaims)
}

func (m *tokenManager) PublicKey() []byte {
	key := &RSAPrivateKey{
		PrivateKey: m.privateKey.PrivateKey(),
	}
	return key.PublicPEM()
}

func NewClaimsWithClockDrift(second int64) *ClaimsWithClockDrift {
	return &ClaimsWithClockDrift{drift: second, StandardClaims: &jwt.StandardClaims{}}
}

type ClaimsWithClockDrift struct {
	drift int64
	*jwt.StandardClaims
}

func (c *ClaimsWithClockDrift) Valid() error {
	c.StandardClaims.IssuedAt -= c.drift
	valid := c.StandardClaims.Valid()
	c.StandardClaims.IssuedAt += c.drift
	return valid
}
