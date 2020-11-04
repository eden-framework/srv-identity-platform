package ewechat

import (
	"fmt"
	"github.com/eden-framework/courier/client"
	"github.com/eden-framework/srv-identity-platform/internal/clients/client_ewechat"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/common"
)

type EWechat struct {
	client *client_ewechat.ClientEWechat
}

func NewEWechatProvider(config global.EWechatConfig) *EWechat {
	p := &EWechat{
		client: &client_ewechat.ClientEWechat{
			Client: client.Client{
				Host: "qyapi.weixin.qq.com",
				Mode: "https",
			},

			CorpID:     config.CorpID,
			CortSecret: config.CorpSecret.String(),
		},
	}
	p.client.Init()
	return p
}

func (d EWechat) ProviderID() enums.BindType {
	return enums.BIND_TYPE__EWECHAT
}

func (d EWechat) ProviderName() string {
	return "企业微信"
}

func (d EWechat) ProviderIcon() string {
	return "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPHN2ZyB3aWR0aD0iMzJweCIgaGVpZ2h0PSIzMnB4IiB2aWV3Qm94PSIwIDAgMzIgMzIiIHZlcnNpb249IjEuMSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayI+CiAgICA8IS0tIEdlbmVyYXRvcjogU2tldGNoIDU1LjEgKDc4MTM2KSAtIGh0dHBzOi8vc2tldGNoYXBwLmNvbSAtLT4KICAgIDx0aXRsZT5lbnRlcnByaXNlIFdlQ2hhdDwvdGl0bGU+CiAgICA8ZGVzYz5DcmVhdGVkIHdpdGggU2tldGNoLjwvZGVzYz4KICAgIDxnIGlkPSJsYW5kaW5nLXBhZ2UiIHN0cm9rZT0ibm9uZSIgc3Ryb2tlLXdpZHRoPSIxIiBmaWxsPSJub25lIiBmaWxsLXJ1bGU9ImV2ZW5vZGQiPgogICAgICAgIDxnIGlkPSJsYW5kaW5nLXBhZ2UtZGVmYXVsdCIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTE5Mi4wMDAwMDAsIC00OTYuMDAwMDAwKSI+CiAgICAgICAgICAgIDxnIGlkPSLnvJbnu4QiIHRyYW5zZm9ybT0idHJhbnNsYXRlKDgwLjAwMDAwMCwgMjEyLjAwMDAwMCkiPgogICAgICAgICAgICAgICAgPGcgaWQ9ImVudGVycHJpc2UtV2VDaGF0IiB0cmFuc2Zvcm09InRyYW5zbGF0ZSgxMTIuMDAwMDAwLCAyODQuMDAwMDAwKSI+CiAgICAgICAgICAgICAgICAgICAgPGNpcmNsZSBpZD0i5qSt5ZyG5b2iIiBmaWxsPSIjMzE4Q0ZGIiBjeD0iMTYiIGN5PSIxNiIgcj0iMTYiPjwvY2lyY2xlPgogICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9Ik0yMC4xODkyMDQzLDIxLjMzNDI2NTcgQzIwLjI5NTc4MTIsMjEuMjI2MzIxOCAyMC40NjgzMDQyLDIxLjIyNjMyMTggMjAuNTc0ODgxNSwyMS4zMzQyNjU3IEMyMS4wODc3NzQxLDIxLjkwNjk5MjYgMjEuNzcwOTA3LDIyLjI5MzkzMiAyMi41MjAzMzMzLDIyLjQzNjIwNDggQzIyLjk4ODE3NTksMjIuNDg1ODkzNyAyMy4zNzc5MTcyLDIyLjgyMjIzNDcgMjMuNTAxNTkwMSwyMy4yODMwMTU1IEMyMy42NjcxMTU4LDIzLjg5OTczMjggMjMuMzA4ODc2NCwyNC41MzU5MTY0IDIyLjcwMTQzOTQsMjQuNzAzOTcxMSBDMjIuMDk0MDAyNCwyNC44NzIwMjU3IDIxLjQ2NzM5MjEsMjQuNTA4MzEzMiAyMS4zMDE4NjYyLDIzLjg5MTU5NiBDMjEuMTk2ODA0MiwyMy4wNzQ2ODk3IDIwLjgxNTU3NjYsMjIuMzIwMTg2IDIwLjIyMzMzNTIsMjEuNzU3MDE4OSBDMjAuMjEwNjA3LDIxLjc0NzgzNDQgMjAuMTk4Njc4MiwyMS43Mzc1NTg1IDIwLjE4NzY4MDEsMjEuNzI2MzAzOSBDMjAuMDgxNDcwOCwyMS42MTc2MTY1IDIwLjA4MjE1MzIsMjEuNDQyMDk0MSAyMC4xODkyMDQ0LDIxLjMzNDI2MjIgWiBNMjQuODE3MjY4LDE5LjUwMTE2NTggTDI0LjgyMDY4MzIsMTkuNTAxMTY5MyBDMjQuOTYxODE3NCwxOS4zNTcxOCAyNS4xMzc3OTA2LDE5LjI1MzQ0MDEgMjUuMzMwOTI0MiwxOS4yMDAzNzM5IEMyNS45Mzg2NDk2LDE5LjAzMzM5MiAyNi41NjQ2MzUzLDE5LjM5ODIxMTEgMjYuNzI5MTA0MywyMC4wMTUyMTk0IEMyNi44OTM1NzM0LDIwLjYzMjIyOTQgMjYuNTM0MjQ0MSwyMS4yNjc3Nzg4IDI1LjkyNjUyMDUsMjEuNDM0NzYwNSBDMjUuMTIxODIxOSwyMS41NDExNDg0IDI0LjM3ODU4MzYsMjEuOTI4MjQ2IDIzLjgyNDA2NDEsMjIuNTI5NzY5MiBDMjMuODEyOTIxMSwyMi41NDY4MjMgMjMuNzk5OTUyMywyMi41NjI1NzE3IDIzLjc4NTM5NDksMjIuNTc2NzI3MyBDMjMuNjc2NDg4MiwyMi42ODI2MjcgMjMuNTAzNjQ1NCwyMi42Nzg4NDA2IDIzLjM5OTMzOTMsMjIuNTY4MjcwNCBDMjMuMjk1MDMzMiwyMi40NTc2OTk5IDIzLjI5ODc2MjYsMjIuMjgyMjE2NCAyMy40MDc2Njg5LDIyLjE3NjMxNjggQzIzLjk3MTE2NDMsMjEuNjU0MTgwNiAyNC4zNTExMDU2LDIwLjk1OTM1MDQgMjQuNDg5NjEzMSwyMC4xOTc2NzQ0IEMyNC41MTY4NDI3LDE5LjkzNDIwMTEgMjQuNjMyNTM3MiwxOS42ODgyNjU0IDI0LjgxNzI2OCwxOS41MDExNjU4IFogTTIxLjYwNTYzMzYsMTYuMjM2OTIxMyBDMjEuNzQ0NzYxMywxNi4wOTU3MTM3IDIxLjkxNzUxMzIsMTUuOTkzNDY0IDIyLjEwNzA5MDQsMTUuOTQwMTE1NiBDMjIuNzEzODAyNCwxNS43NjkzODI2IDIzLjM0MTk2NDUsMTYuMTMwMzI2OSAyMy41MTAxMjgzLDE2Ljc0NjMwODEgQzIzLjYxNTE5MDMsMTcuNTYzMjE0NCAyMy45OTY0MTc5LDE4LjMxNzcxODEgMjQuNTg4NjU5MSwxOC44ODA4ODUgQzI0LjYwNTQ1NjMsMTguODkyMTk4MiAyNC42MjA5NjgsMTguOTA1MzY1MSAyNC42MzQ5MTA1LDE4LjkyMDE0NDkgQzI0LjczOTIxNjYsMTkuMDMwNzE1NSAyNC43MzU0ODcyLDE5LjIwNjE5OSAyNC42MjY1ODA5LDE5LjMxMjA5ODYgQzI0LjUxNzY3NDIsMTkuNDE3OTk4MiAyNC4zNDQ4MzEzLDE5LjQxNDIxMTggMjQuMjQwNTI1MywxOS4zMDM2NDE2IEMyMy43MjYyNDYyLDE4LjczMTUzNzMgMjMuMDQxODcxOCwxOC4zNDU3OTEzIDIyLjI5MTY1NzMsMTguMjA1MTY3OCBDMjIuMDMzMDY4NywxOC4xNzcxMzgxIDIxLjc5MTY5MjEsMTguMDYwMTgyIDIxLjYwNzU4MTEsMTcuODczNzEgQzIxLjE2MTg1ODUsMTcuNDIyMjcwNyAyMS4xNjA5ODM2LDE2LjY4OTQ1NjggMjEuNjA1NjI5NywxNi4yMzY5MjQ0IFogTTIxLjA0NTg4OTIsMTguMDQ1NzcwOCBDMjEuMTU4MDQ1NiwxNy45NDMzODMyIDIxLjMzMDcxOTYsMTcuOTUyNjkxOCAyMS40MzE1NjY0LDE4LjA2NjU2MjEgQzIxLjUzMjQxMzIsMTguMTgwNDMyIDIxLjUyMzI0NDgsMTguMzU1NzQ0MSAyMS40MTEwODgsMTguNDU4MTMxNiBDMjAuODQ2NTMyNywxOC45Nzk5Mjc2IDIwLjQ2NTQwMjEsMTkuNjc0NzQwOSAyMC4zMjU3MzA3LDIwLjQzNjc3NDQgQzIwLjI5ODEyMjgsMjAuNjk5MzEzNyAyMC4xODI5MjY2LDIwLjk0NDM3OCAxOS45OTkyNjA2LDIxLjEzMTMwMTggQzE5LjU1NDYxNDUsMjEuNTgzODM0IDE4LjgzMjgyNzksMjEuNTg0NzIyMyAxOC4zODcxMDUzLDIxLjEzMzI4MyBDMTguMjQxNDAwNywyMC45ODgwOTU4IDE4LjEzNzAzNjMsMjAuODA1NjQwNiAxOC4wODUwNDMzLDIwLjYwNTE5OTUgQzE3LjkyNTY2MzgsMTkuOTkwNzcxNCAxOC4yODcwNTk1LDE5LjM2MTUwNTEgMTguODkyMjQwMSwxOS4xOTk2OTE4IEMxOS43MTMwMzQxLDE5LjA5MDM5MzYgMjAuNDY4OTY5OCwxOC42ODg4NDY0IDIxLjAyNTQxMDYsMTguMDY2NTY1OSBMMjEuMDQ1ODg5MSwxOC4wNDU3NzQ2IFogTTE0LjE1ODMxNDUsOC4wNTU1NTM1NSBMMTQuMTU4Mjk0Miw4LjA1NTU3Mzc4IEMxNC44NTk0OTkxLDcuOTc5MjI2MiAxNS41NjY4NTg1LDcuOTgxNTUwNSAxNi4yNjc1NjYyLDguMDYyNTA0NTUgQzE4LjIzODM0OTIsOC4yNjU2MjkyMSAyMC4wNzg4OTM4LDkuMTU1ODc3IDIxLjQ3NTkxODgsMTAuNTgxNzI0MSBDMjEuOTg5MDQ2MywxMS4xMTc3NjM0IDIyLjQxNTI3NTUsMTEuNzMyOTU3OSAyMi43Mzg3NTM5LDEyLjQwNDQzMTIgQzIzLjIzMjk1MDgsMTMuNDIyODEyNSAyMy40NDc1NDIsMTQuNTU3NjU1OCAyMy4zNTk5MzE2LDE1LjY4OTQ0OSBMMjMuMzA4NzM1NSwxNS42MjcwNzUgQzIyLjkzNjQ3MDUsMTUuMjQ5MzgxOCAyMi4zNTc1MDE0LDE1LjE4MDU2NDcgMjEuOTA5Mzc1NiwxNS40NjA3NDQ3IEMyMS45MTI3ODg3LDE1LjM4NDUxIDIxLjkxOTYxNDksMTUuMzA0ODEgMjEuOTE5NjE0OSwxNS4yMjUxMSBDMjEuOTE5NjE0OSwxNC40NjYyMjc0IDIxLjc0ODk2MTIsMTMuNzI0NjcxIDIxLjQxMTA2NywxMy4wMjQ2OTcxIEMyMS4xNTg0OTk2LDEyLjQ5Nzk4NDEgMjAuODE3MTkyMywxMi4wMDI0NTggMjAuMzk3Mzg0MywxMS41NjIzNzU1IEMxOS4zMTIwMTcxLDEwLjQyNTc5MSAxNy43ODI5NjcxLDkuNjg3Njk5NzYgMTYuMDkzNDk2LDkuNDkwMTgyMzggQzE1LjUwNTk1NzUsOS40MjQzNzYxMyAxNC45MTMwNzEsOS40MjQzNzYzIDE0LjMyNTUzMTIsOS40OTAxODI5IEMxMi42MjU4Mjc1LDkuNjc3MzAxMjQgMTEuMDg5OTQ0NywxMC40MDQ5OTY4IDkuOTk3NzYxMzIsMTEuNTQ4NTE4NSBDOS41NzQ1NDAyNywxMS45ODg2MDExIDkuMjMzMjMyOTgsMTIuNDc3MTk2NyA4Ljk3MzgzOTQzLDEzLjAwMzkwOTcgQzguMTgxMTM2NTcsMTQuNjI4NzcxNyA4LjMxNjk2MTIxLDE2LjU2MjAwNzUgOS4zMjg4MDA3OCwxOC4wNTYxODk2IEM5LjcyNTIxNTU0LDE4LjY0ODA1NjUgMTAuMjE2NzA0NywxOS4xNjgxNDAzIDEwLjc4Mjc3MDQsMTkuNTk0NzQ2NSBDMTAuOTkyODEzOSwxOS43NTA3NjkyIDExLjA4MDgyOSwyMC4wMjU4MjkzIDExLjAwMTIwNzEsMjAuMjc3Mzk0MyBMMTAuODgxNzQ5NSwyMC43Mjc4NzI2IEwxMC43NzkzNTczLDIxLjEzMzMwMyBMMTAuNzE0NTA4OSwyMS4zNzkzMzM0IEwxMC42NTk4OTk4LDIxLjYwNDU3MjUgQzEwLjY1NzkzMTksMjEuNjU4ODAyNiAxMC42OTg3MjcxLDIxLjcwNDgyMzMgMTAuNzUyMDUyNywyMS43MDg1MjkgQzEwLjc2OTExODEsMjEuNzA4NTI5IDEwLjc4NjE4MzUsMjEuNzAxNTk4NiAxMC44MjAzMTQyLDIxLjY4MDgwNzMgTDEwLjg4MTc0OTUsMjEuNjQyNjg5OSBMMTIuMzIyMDY2MywyMC43ODMzMTYgTDEyLjM2OTg0OTMsMjAuNzU5MDU5NSBDMTIuNTU5Mjc4MiwyMC42NTE2MzIgMTIuNzgzODI5NCwyMC42Mjc4MzE4IDEyLjk5MTAyODYsMjAuNjkzMjIwNCBDMTMuMzcxMTM5NiwyMC44MDQ5Nzc1IDEzLjc1OTA3MzMsMjAuODg3MjI0NyAxNC4xNTE0NzA3LDIwLjkzOTI1MDkgTDE0LjMyNTUzOTEsMjAuOTYzNTA2OSBDMTUuNDEyMzcyOSwyMS4wODg2ODI4IDE2LjUxMjk5NTYsMjAuOTc3NjQ4NiAxNy41NTQzMDk0LDIwLjYzNzc3NjUgQzE3LjUwMTkzMTcsMjEuMTc1MzE4MyAxNy44MjM4NTM4LDIxLjY3NzgxNzEgMTguMzI5MDc3LDIxLjg0NzEzNzMgQzE3LjMyMjIzNzEsMjIuMjE2Mzk0NSAxNi4yNTk1NTgyLDIyLjQwNDAzMTkgMTUuMTg5MDQ5NiwyMi40MDE1NzIxIEMxNC4zMjMzNDksMjIuNDA3Njg3IDEzLjQ2MTI3MTQsMjIuMjg3NDg3MiAxMi42MjkyNDQ3LDIyLjA0NDY1NDYgTDEwLjg3ODM0NDksMjIuOTM4NjgwNyBMMTAuMjkxMjk2MywyMy4yMzY2ODkzIEwxMC4yNTM3NTI1LDIzLjI2NDQxMTEgTDEwLjI0MzUxMzMsMjMuMjY0NDExMSBDMTAuMTYzMzAyOCwyMy4zMTI1OTE1IDEwLjA3MzQ0MjgsMjMuMzQxODMwMyA5Ljk4MDU5MDA0LDIzLjM0OTk2MDggQzkuNjM4NzkxMTEsMjMuMzc5ODkwNiA5LjMzNzgxMDI0LDIzLjEyMjgzNjggOS4zMDgzMzEzMSwyMi43NzU4MTU5IEw5LjMwMTUwNTE3LDIyLjcxMzQ0MiBMOS4zMDgzMzEzMSwyMi42NTEwNjgxIEM5LjMxMTc0NDM5LDIyLjYyNjgxMTUgOS4zMTE3NDQzOSwyMi41OTkwODk4IDkuMzE4NTcwNTMsMjIuNTc0ODMzMyBMOS4zNDU4NzUxMiwyMi40OTE2NjgxIEw5LjU5MTYxNjM3LDIwLjUzMDM1MTggQzkuMDQ4OTM3NzcsMjAuMDM4Mjg0MiA4LjQ1NTA2MzA3LDE5LjM0NTI0MDggOC4wOTMyNzczNCwxOC43OTA4MDYgQzYuODEzMzk0OTcsMTYuODkyNzE5OCA2LjY0NDk1NDQ5LDE0LjQ0MDIxNTggNy42NTI5OTA3NCwxMi4zODAxNTQ3IEM3Ljk3MDQwNjUzLDExLjcyMTc2MDEgOC40MDA0NTM3MiwxMS4xMDQ5NTE0IDguOTI5NDgwMDMsMTAuNTU3NDQ3MSBDMTAuMjYwNTc4NSw5LjE3MTM1MzQ3IDEyLjExMzg4NzEsOC4yODQyNTc4OCAxNC4xNTgzMTQ1LDguMDU1NTUzNTUgWiIgaWQ9IuW9oueKtue7k+WQiCIgZmlsbD0iI0ZGRkZGRiI+PC9wYXRoPgogICAgICAgICAgICAgICAgPC9nPgogICAgICAgICAgICA8L2c+CiAgICAgICAgPC9nPgogICAgPC9nPgo8L3N2Zz4="
}

func (d EWechat) ProviderEntry() string {
	return fmt.Sprintf("https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=%s&agentid=%s&state=%s_{{RANDOM_NUMBER}}&redirect_uri={{REDIRECT_URI}}", global.ProviderConfig.EWechat.CorpID, global.ProviderConfig.EWechat.AgentID, d.ProviderID().String())
}

func (d *EWechat) GetUserID(token string) (string, error) {
	resp, err := d.client.GetUserInfoByCode(client_ewechat.GetUserInfoByCodeRequest{
		Code: token,
	})
	if err != nil {
		return "", err
	}

	return resp.UserID, nil
}

func (d *EWechat) GetUserInfo(userID string) (user common.UserInfo, err error) {
	userInfo, err := d.client.GetUserInfoDetail(client_ewechat.GetUserInfoDetailRequest{
		UserID: userID,
	})
	if err != nil {
		return
	}

	return common.UserInfo{
		UserID: userInfo.UserID,
		Name:   userInfo.Name,
		Mobile: userInfo.Mobile,
		Email:  userInfo.Email,
	}, nil
}
