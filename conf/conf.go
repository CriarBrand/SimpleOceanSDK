package conf

const (

	// API_HOST OpenAPI HOST
	API_HOST = "ad.oceanengine.com"

	// API_HTTP_SCHEME 协议
	API_HTTP_SCHEME = "https://"

	// API_OAUTH_CONNECT 生成授权链接
	API_OAUTH_CONNECT = "/openapi/qc/audit/oauth.html"

	// API_OAUTH_ACCESS_TOKEN 获取access_token
	API_OAUTH_ACCESS_TOKEN = "/open_api/oauth2/access_token/"

	// API_OAUTH_REFRESH_TOKEN 刷新access_token
	API_OAUTH_REFRESH_TOKEN = "/open_api/oauth2/refresh_token/"

	// API_GET_ADVERTISER 获取已授权账户
	API_GET_ADVERTISER = "/open_api/oauth2/advertiser/get/"

	// API_USER_INFO 获取授权时登录用户信息
	API_USER_INFO = "/open_api/2/user/info/"

	// API_TOOLS_ESTIMATE_AUDIENCE 查询受众预估结果
	API_TOOLS_ESTIMATE_AUDIENCE = "/open_api/2/tools/estimate_audience/"

	// API_AWEME_SIMILAR_AUTHOR_SEARCH 查询抖音类似帐号
	API_AWEME_SIMILAR_AUTHOR_SEARCH = "/open_api/2/tools/aweme_similar_author_search/"

	// API_AWEME_INFO_SEARCH  查询抖音帐号和类目信息
	API_AWEME_INFO_SEARCH = "/open_api/2/tools/aweme_info_search/"

	// API_AWEME_AUTHOR_INFO  查询抖音号id对应的达人信息
	API_AWEME_AUTHOR_INFO = "/open_api/2/tools/aweme_author_info/get/"

	// API_INTEREST_ACTION_ID2WORD  兴趣行为类目关键词id转词
	API_INTEREST_ACTION_ID2WORD = "/open_api/2/tools/interest_action/id2word/"

	// API_INTEREST_ACTION_KEYWORD_SUGGEST 获取行为兴趣推荐关键词
	API_INTEREST_ACTION_KEYWORD_SUGGEST = "/open_api/2/tools/interest_action/keyword/suggest/"
)
