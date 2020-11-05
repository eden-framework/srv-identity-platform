package token

import (
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/profzone/envconfig"
)

type SecureCode struct {
	Code     string             `json:"code"`
	ExpireIn envconfig.Duration `json:"expireIn"`
}

type State string

func (s *State) UnmarshalBinary(data []byte) error {
	*s = State(data)
	return nil
}

func (s State) MarshalBinary() (data []byte, err error) {
	data = []byte(s)
	return
}

type AccessToken struct {
	databases.TokenScope
	AccessToken  string             `json:"access_token"`
	RefreshToken string             `json:"refresh_token"`
	ExpireIn     envconfig.Duration `json:"expire_in"`
}

type SignedToken struct {
	*databases.Token
	Signed string `json:"signed"`
}
