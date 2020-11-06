package token

import (
	"bytes"
	"net/http"
)

func ParseAuthorizationToken(s string) (auths AuthorizationToken) {
	auths = AuthorizationToken{}
	if len(s) == 0 {
		return
	}
	tokens := bytes.Split([]byte(s), []byte(";"))
	for _, token := range tokens {
		kv := bytes.Split(bytes.TrimSpace(token), []byte(" "))
		v := ""
		if len(kv) == 2 {
			v = string(bytes.TrimSpace(kv[1]))
		}
		auths[http.CanonicalHeaderKey(string(bytes.TrimSpace(kv[0])))] = v
	}
	return
}

type AuthorizationToken map[string]string

func (a AuthorizationToken) Add(k, v string) {
	a[http.CanonicalHeaderKey(k)] = v
}

func (a AuthorizationToken) Get(k string) string {
	if v, ok := a[http.CanonicalHeaderKey(k)]; ok {
		return v
	}
	return ""
}

func (a AuthorizationToken) String() string {
	buf := bytes.Buffer{}

	count := 0
	for tpe, token := range a {
		if count > 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(http.CanonicalHeaderKey(tpe))
		buf.WriteString(" ")
		buf.WriteString(token)
		count++
	}
	return buf.String()
}
