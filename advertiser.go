package SimpleOceanSDK

import (
	"github.com/CriarBrand/SimpleOceanSDK/conf"
	"github.com/guonaihong/gout"
)

// AdvertiserListResData 获取已授权的账户（店铺/代理商）-返回
type AdvertiserListResData struct {
	AccountRole    string `json:"account_role"`    // 新版账号角色，见【枚举值-账户角色】
	AdvertiserId   int64  `json:"advertiser_id"`   // 账号id
	AdvertiserName string `json:"advertiser_name"` // 账号名称
	AdvertiserRole int    `json:"advertiser_role"` // 旧版账号角色，1-普通广告主，2-纵横组织账户，3-一级代理商，4-二级代理商，6-星图账号
	CompanyList    []struct {
		CustomerCompanyId   int64  `json:"customer_company_id"`   // 客户公司id
		CustomerCompanyName string `json:"customer_company_name"` // 客户公司名
	} `json:"company_list"` // 代理商账户下勾选账户，但授权时选择的是代理商类型账户时，该字段才有意义 company_list为空时，代表当前代理商账户下所有adv均可访问；不为空时，代表仅能访问该部分客户id下的adv
	IsValid bool `json:"is_valid"` // 授权有效性，允许值：true/false；false表示对应的user在客户中心/一站式平台代理商平台变更了对此账号的权限,需要到对应平台进行调整过来；
}

type AdvertiserListResDataCom struct {
	List []AdvertiserListResData `json:"list"`
}

// GetAdvertiserList 获取已授权的账户（店铺/代理商）
func (client *Client) GetAdvertiserList(accessToken string, response *AdvertiserListResDataCom) error {
	df := gout.GET(client.url(conf.API_GET_ADVERTISER)).
		SetQuery(gout.H{
			"access_token": accessToken,
			"app_id":       client.appId,
			"secret":       client.secret,
		})
	return client.DoRequest(df, response)
}
