package SimqianchuanSDK

import (
	"github.com/CriarBrand/SimpleOceanSDK/conf"
	"github.com/guonaihong/gout"
)

// -----------------------------------------------------获取授权时登录用户信息----------------------------------------------

type UserInfoRes struct {
	ID          int64  `json:"id"`           //用户id
	Email       string `json:"email"`        //邮箱（已经脱敏处理）
	DisplayName string `json:"display_name"` // 用户名
}

// GetUserInfo 获取授权时登录用户信息
func (client *Client) GetUserInfo(accessToken string, response *UserInfoRes) error {
	df := gout.GET(client.url(conf.API_USER_INFO)).
		SetHeader(gout.H{
			"Access-Token": accessToken,
		})
	return client.DoRequest(df, response)
}
