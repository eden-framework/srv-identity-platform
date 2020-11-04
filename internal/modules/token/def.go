package token

import (
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
