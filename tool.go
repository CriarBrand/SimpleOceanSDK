package SimpleOceanSDK

import (
	"github.com/CriarBrand/SimpleOceanSDK/conf"
	"github.com/CriarBrand/SimpleOceanSDK/utils"
	"github.com/guonaihong/gout"
)

// EstimateAudienceReq 查询受众预估结果请求
type EstimateAudienceReq struct {
	AccessToken string `json:"access_token"`
	EstimateAudienceReqBase
}

// EstimateAudienceReqBase 查询受众预估结果请求
type EstimateAudienceReqBase struct {
	AdvertiserId              int64                            `json:"advertiser_id"`                          // 广告主id
	AudiencePackageId         int64                            `json:"audience_package_id,omitempty"`          // 定向包ID，定向包ID由 【工具-定向包管理-获取定向包】 获取 如果同时传定向包ID和自定义用户定向参数时， 仅定向包中的定向生效
	District                  string                           `json:"district,omitempty"`                     // 地域类型 允许值: CITY 省市、 COUNTY 区县、 BUSINESS_DISTRICT 商圈、 REGION 行政区域、 OVERSEA 海外区域、 NONE 不限 使用省市示例 ：{"district": "CITY","city": [12]} 使用区县示例 ：{"district": "COUNTY","city": [130102]} 使用行政区域示例 ：{"district":"REGION": "city":[31], "region_versio":"1.0.0"} 使用海外区域示例 ：{"district":"OVERSEA": "city":[3041566], "region_versio":"1.0.0"}
	RegionVersion             string                           `json:"region_version,omitempty"`               // 行政区域版本号。通过 【获取行政信息】 接口获取 district = REGION / OVERSEA 时 必填
	City                      []int64                          `json:"city,omitempty"`                         // 地域定向省市或者区县列表(当传递省份ID时,旗下市县ID可省略不传)，当district= CITY / COUNTY / REGION / OVERSEA 时 必填 district = CITY / COUNTY 时，详见 【附件-city.json】 district = REGION / OVERSEA 时，通过 【获取行政信息】 接口获取
	Geolocation               []EstimateAudienceGeolocationReq `json:"geolocation,omitempty"`                  // 从地图添加(地图位置)，district为 "BUSINESS_DISTRICT" 时填写
	LocationType              string                           `json:"location_type,omitempty"`                // 位置类型 允许值： CURRENT 正在该地区的用户， HOME 居住在该地区的用户， TRAVEL 到该地区旅行的用户， ALL 该地区内的所有用户 当city和district有值时 必填
	Gender                    string                           `json:"gender,omitempty"`                       // 性别, 详见 【附录-受众性别】 允许值: GENDER_FEMALE , GENDER_MALE , NONE
	Age                       []string                         `json:"age,omitempty"`                          // 年龄, 详见 【附录-受众年龄区间】 允许值: AGE_BETWEEN_18_23 , AGE_BETWEEN_24_30 , AGE_BETWEEN_31_40 , AGE_BETWEEN_41_49 , AGE_ABOVE_50
	RetargetingTagsInclude    []int64                          `json:"retargeting_tags_include,omitempty"`     // 定向人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude IOS14.5 专属广告计划在广告位为穿山甲时 不支持
	RetargetingTagsExclude    []int64                          `json:"retargeting_tags_exclude,omitempty"`     // 排除人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude IOS14.5 专属广告计划在广告位为穿山甲时 不支持
	InterestActionMode        string                           `json:"interest_action_mode,omitempty"`         // 行为兴趣 允许值： UNLIMITED 不限, CUSTOM 自定义, RECOMMEND 系统推荐。 若与自定义人群同时使用，系统推荐( RECOMMEND )不生效 IOS14.5 专属广告计划在广告位为穿山甲时 只支持 UNLIMITED
	ActionScene               []string                         `json:"action_scene,omitempty"`                 // 行为场景，详见 【附录-行为场景】 ，当interest_action_mode传 CUSTOM 时有效 允许值： E-COMMERCE 电商互动行为, NEWS 资讯互动行为, APP APP推广互动行为 IOS14.5 专属广告计划在广告位为穿山甲时 不支持 行为场景
	ActionDays                int64                            `json:"action_days,omitempty"`                  // 用户发生行为天数，当interest_action_mode传 CUSTOM 时有效 允许值： 7 , 15 , 30 , 60 , 90 , 180 , 365 IOS14.5 专属广告计划在广告位为穿山甲时 不支持 用户发生行为天数
	ActionCategories          []int64                          `json:"action_categories,omitempty"`            // 行为类目词，当interest_action_mode传 CUSTOM 时有效 行为类目可以通过 【工具-行为兴趣词管理-行为类目查询】 获取 IOS14.5 专属广告计划在广告位为穿山甲时 不支持 行为类目词
	ActionWords               []int64                          `json:"action_words,omitempty"`                 // 行为关键词，当interest_action_mode传 CUSTOM 时有效 行为类目可以通过 【工具-行为兴趣词管理-行为关键词查询】 获取 IOS14.5 专属广告计划在广告位为穿山甲时 不支持 行为关键词
	InterestCategories        []int64                          `json:"interest_categories,omitempty"`          // 兴趣类目词，当interest_action_mode传 CUSTOM 时有效 IOS14.5 专属广告计划在广告位为穿山甲时 不支持 兴趣类目词
	InterestWords             []int64                          `json:"interest_words,omitempty"`               // 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。当interest_action_mode传 CUSTOM 时有效 IOS14.5 专属广告计划在广告位为穿山甲时 不支持 兴趣关键词
	AwemeFanBehaviors         []string                         `json:"aweme_fan_behaviors,omitempty"`          // 抖音达人互动用户行为类型, 详见 【附录-抖音用户行为类型】 允许值: FOLLOWED_USER , COMMENTED_USER , LIKED_USER , SHARED_USER （抖音达人定向）
	AwemeFanTimeScope         string                           `json:"aweme_fan_time_scope,omitempty"`         // 抖音达人互动行为时间范围，允许值： FIFTEEN_DAYS 15天、 THIRTY_DAYS 30天、 SIXTY_DAYS 60天 ，默认时间范围为15天
	AwemeFanCategories        []int64                          `json:"aweme_fan_categories,omitempty"`         // 抖音达人分类ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向），可通过 【工具-抖音达人-查询抖音类目列表】 接口获取
	AwemeFanAccounts          []int64                          `json:"aweme_fan_accounts,omitempty"`           // 抖音达人ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向），可通过 【工具-抖音达人-查询抖音类目下的推荐达人】 接口获取。
	FilterAwemeAbnormalActive int64                            `json:"filter_aweme_abnormal_active,omitempty"` // （抖音号、直播间推广特有）过滤高活跃用户 允许值：0表示不过滤，1表示过滤
	FilterAwemeFansCount      int64                            `json:"filter_aweme_fans_count,omitempty"`      // （抖音号、直播间推广特有）过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户
	FilterOwnAwemeFans        int64                            `json:"filter_own_aweme_fans,omitempty"`        // （抖音号、直播间推广特有）过滤自己的粉丝 允许值： 0 表示不过滤， 1 表示过滤
	SuperiorPopularityType    string                           `json:"superior_popularity_type,omitempty"`     // 媒体定向，详见 【附录-媒体定向】 。 对于选择自定义媒体定向流量包，该字段不传，传 flow_package 或 exclude_flow_package 字段即可 媒体定向仅支持穿山甲、穿山甲-精选游戏广告位
	FlowPackage               []int64                          `json:"flow_package,omitempty"`                 // 定向逻辑，可通过 【工具-穿山甲流量包-获取穿山甲流量包】 定向逻辑和排除定向逻辑只能选其一
	ExcludeFlowPackage        []int64                          `json:"exclude_flow_package,omitempty"`         // 排除定向逻辑，可通过 【工具-穿山甲流量包-获取穿山甲流量包】 定向逻辑和排除定向逻辑只能选其一
	Platform                  []string                         `json:"platform,omitempty"`                     // 平台，当下载方式包含下载链接时，平台类型需与选择的下载链接类型对应，当下载方式不包含下载链接的时候，平台可多选。 为保证投放效果,平台类型定向PC与移动端互斥。详见 【附录-受众平台类型】 允许值: ANDROID , IOS
	AndroidOsv                string                           `json:"android_osv,omitempty"`                  // 最低安卓版本，当app_type为 APP_ANDROID 选填 ,其余情况不填, 详见 【附录-受众android版本】 允许值: 0.0 , 2.0 , 2.1 , 2.2 , 2.3 , 3.0 , 3.1 , 3.2 , 4.0 , 4.1 , 4.2 , 4.3 , 4.4 , 4.5 , 5.0 , NONE ……
	IosOsv                    string                           `json:"ios_osv,omitempty"`                      // 最低IOS版本，当app_type为 APP_IOS 选填 ,其余情况不填, 详见【附录-受众ios版本】允许值: 0.0 , 4.0 , 4.1 , 4.2 , 4.3 , 5.0 , 5.1 , 6.0 , 7.0 , 7.1 , 8.0 , 8.1 , 8.2 , 9.0 , NONE ……
	DeviceType                []string                         `json:"device_type,omitempty"`                  // 设备类型。 允许值是： MOBILE , PAD 。缺省表示不限设备类型。 穿山甲已经全量，投放范围为默认时需要有白名单权限才可以
	Ac                        []string                         `json:"ac,omitempty"`                           // 网络类型, 详见 【附录-受众网络类型】 允许值: WIFI , 2G , 3G , 4G , 5G
	Carrier                   []string                         `json:"carrier,omitempty"`                      // 运营商, 详见 【附录-受众运营商类型】 允许值: MOBILE , UNICOM , TELCOM
	ActivateType              []string                         `json:"activate_type,omitempty"`                // 新用户(新用户使用头条的时间，只有选择快应用+穿山甲时，才可选择转化的时间范围)，详见 【附录-用户首次激活时间】 ， 允许值: WITH_IN_A_MONTH 一个月以内, ONE_MONTH_2_THREE_MONTH 一个月到三个月, THREE_MONTH_EAILIER 三个月或更早
	ArticleCategory           []string                         `json:"article_category,omitempty"`             // 文章分类, 详见 【附录-受众文章分类】 文章分类：只针对投放在详情页位置的广告生效, 不支持人群预估
	DeviceBrand               []string                         `json:"device_brand,omitempty"`                 // 手机品牌, 详见 【附录-手机品牌】 当下载链接为Android应用时，不能选择 APPLE 品牌，当下载链接为IOS应用时，不能选择非 APPLE 品牌，否则会报错
	LaunchPrice               []int64                          `json:"launch_price,omitempty"`                 // 手机价格,传入价格区间，最高传入11000（表示1w以上） 传值示例 "launch_price": [2000, 11000]，表示2000元以上; 注意：手机价格须为500的整数，launch_price数组长度固定为2
	AutoExtendEnabled         int64                            `json:"auto_extend_enabled,omitempty"`          // 是否启用智能放量。 允许值是： 0 、 1 。缺省为 0 。 注意：智能放量不支持受众预估
	AutoExtendTargets         []string                         `json:"auto_extend_targets,omitempty"`          // 可放开定向。当auto_extend_enabled=1 时选填，缺省为全不选。 允许值： AGE 年龄 REGION 地域 GENDER 性别 INTEREST_ACTION 行为兴趣 CUSTOM_AUDIENCE 自定人群-定向
}

type EstimateAudienceGeolocationReq struct {
	Radius int64   `json:"radius"` // 半径，单位为m，允许值为：3000-15000
	Name   string  `json:"name"`   // 地点名称
	Long   float64 `json:"long"`   // 经度
	Lat    float64 `json:"lat"`    // 纬度
}

// EstimateAudienceResData 查询受众预估结果返回
type EstimateAudienceResData struct {
	Aweme struct {
		Num int64 `json:"num"`
	} `json:"aweme"`
	TomatoNovelApp struct {
		Num int64 `json:"num"`
	} `json:"tomato_novel_app"`
	VideoApp struct {
		Num int64 `json:"num"`
	} `json:"video_app"`
	Toutiao struct {
		Num int64 `json:"num"`
	} `json:"toutiao"`
}

// GetEstimateAudience 查询受众预估结果
func (client *Client) GetEstimateAudience(request *EstimateAudienceReq, response *EstimateAudienceResData) error {
	df := gout.GET(client.url(conf.API_TOOLS_ESTIMATE_AUDIENCE)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.EstimateAudienceReqBase))
	return client.DoRequest(df, response)
}

// -----------------------------------------------------查询抖音类似帐号---------------------------------------------------

// AwemeSimilarAuthorReq 查询抖音类似帐号
type AwemeSimilarAuthorReq struct {
	AccessToken string `json:"access_token"`
	AwemeSimilarAuthorReqBase
}

type AwemeSimilarAuthorReqBase struct {
	AdvertiserId int64    `json:"advertiser_id"`       // 广告主ID
	AwemeId      string   `json:"aweme_id"`            // 抖音id
	Behaviors    []string `json:"behaviors,omitempty"` // 抖音用户行为类型, 详见【附录-抖音用户行为类型】 允许值:"FOLLOWED_USER","COMMENTED_USER","LIKED_USER","SHARED_USER"
}

type AwemeSimilarAuthorRes struct {
	List []AwemeSimilarAuthorResBase `json:"list"`
}

type AwemeSimilarAuthorResBase struct {
	AuthorName      string `json:"author_name"`        // 抖音作者名
	Avatar          string `json:"avatar"`             // 抖音头像
	AwemeId         string `json:"aweme_id"`           // 抖音号id
	CategoryName    string `json:"category_name"`      // 抖音分类
	CoverNumStr     string `json:"cover_num_str"`      // 覆盖人群数
	LabelId         int64  `json:"label_id"`           // 抖音号id
	TotalFansNumStr string `json:"total_fans_num_str"` // 粉丝数
}

// GetAwemeSimilarAuthor 查询抖音类似帐号
func (client *Client) GetAwemeSimilarAuthor(request *AwemeSimilarAuthorReq, response *AwemeSimilarAuthorRes) error {
	df := gout.GET(client.url(conf.API_AWEME_SIMILAR_AUTHOR_SEARCH)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.AwemeSimilarAuthorReqBase))
	return client.DoRequest(df, response)
}

// -----------------------------------------------------查询抖音帐号和类目信息---------------------------------------------------

// AwemeInfoSearchReq 查询抖音帐号和类目信息
type AwemeInfoSearchReq struct {
	AccessToken string `json:"access_token"`
	AwemeInfoSearchReqBase
}

type AwemeInfoSearchReqBase struct {
	AdvertiserId int64    `json:"advertiser_id"`       // 广告主ID
	QueryWord    string   `json:"query_word"`          // 搜索词
	Behaviors    []string `json:"behaviors,omitempty"` // 抖音用户行为类型, 详见【附录-抖音用户行为类型】 允许值:"FOLLOWED_USER","COMMENTED_USER","LIKED_USER","SHARED_USER"
}

type AwemeInfoSearchRes struct {
	List []AwemeInfoSearchResBase `json:"list"`
}

type AwemeInfoSearchResBase struct {
	Authors    []AwemeInfoSearchAuthors    `json:"authors"`
	Categories []AwemeInfoSearchCategories `json:"categories"`
}
type AwemeInfoSearchAuthors struct {
	AuthorName      string `json:"author_name"`
	Avatar          string `json:"avatar"`
	AwemeId         string `json:"aweme_id"`
	CategoryName    string `json:"category_name"`
	CoverNumStr     string `json:"cover_num_str"`
	LabelId         int64  `json:"label_id"`
	TotalFansNumStr string `json:"total_fans_num_str"`
}

type AwemeInfoSearchCategories struct {
	Id          int64  `json:"id"`
	CoverNumStr string `json:"cover_num_str"`
	FansNumStr  string `json:"fans_num_str"`
	Value       string `json:"value"`
	Children    struct {
		Id          int64  `json:"id"`
		CoverNumStr string `json:"cover_num_str"`
		FansNumStr  string `json:"fans_num_str"`
		Value       string `json:"value"`
		Children    struct {
			Id          int64  `json:"id"`
			CoverNumStr string `json:"cover_num_str"`
			FansNumStr  string `json:"fans_num_str"`
			Value       string `json:"value"`
		} `json:"children"`
	} `json:"children"`
}

// GetAwemeInfoSearch 查询抖音帐号和类目信息
func (client *Client) GetAwemeInfoSearch(request *AwemeInfoSearchReq, response *AwemeInfoSearchRes) error {
	df := gout.GET(client.url(conf.API_AWEME_INFO_SEARCH)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.AwemeInfoSearchReqBase))
	return client.DoRequest(df, response)
}

// ---------------------------------------------------查询抖音号id对应的达人信息-----------------------------------------------

// AwemeAuthorInfoReq 查询抖音号id对应的达人信息
type AwemeAuthorInfoReq struct {
	AccessToken string `json:"access_token"`
	AwemeAuthorInfoReqBase
}

type AwemeAuthorInfoReqBase struct {
	AdvertiserId int64    `json:"advertiser_id"`       // 广告主ID
	LabelIds     []int64  `json:"label_ids"`           // 抖音号id列表，取值大小：1～50；label_id即计划中设置的抖音达人账号的id
	Behaviors    []string `json:"behaviors,omitempty"` // 抖音用户行为类型, 详见【附录-抖音用户行为类型】 允许值:"FOLLOWED_USER","COMMENTED_USER","LIKED_USER","SHARED_USER"
}

type AwemeAuthorInfoRes struct {
	Authors []AwemeAuthorInfoAuthors `json:"authors"`
}

type AwemeAuthorInfoAuthors struct {
	AuthorName      string `json:"author_name"`        // 抖音作者名
	Avatar          string `json:"avatar"`             // 抖音头像
	AwemeId         string `json:"aweme_id"`           // 抖音号id
	CoverNumStr     string `json:"cover_num_str"`      // 覆盖人群数
	LabelId         int64  `json:"label_id"`           // 抖音号id
	TotalFansNumStr string `json:"total_fans_num_str"` // 粉丝数
}

// GetAwemeAuthorInfo 查询抖音号id对应的达人信息
func (client *Client) GetAwemeAuthorInfo(request *AwemeAuthorInfoReq, response *AwemeAuthorInfoRes) error {
	df := gout.GET(client.url(conf.API_AWEME_AUTHOR_INFO)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.AwemeAuthorInfoReqBase))
	return client.DoRequest(df, response)
}

// ---------------------------------------------------兴趣行为类目关键词id转词-----------------------------------------------

// InterestActionID2WordReq 兴趣行为类目关键词id转词
type InterestActionID2WordReq struct {
	AccessToken string `json:"access_token"`
	InterestActionID2WordReqBase
}

type InterestActionID2WordReqBase struct {
	AdvertiserId  int64    `json:"advertiser_id"`          // 广告主ID
	Ids           []int64  `json:"ids"`                    // 类目或关键词id列表
	TagType       string   `json:"tag_type"`               // 查询类型：类目还是关键词 允许值：CATEGORY（类目）、KEYWORD（关键词）
	TargetingType string   `json:"targeting_type"`         // 查询目标：兴趣还是行为 允许值：ACTION（行为）、INTEREST（兴趣）
	ActionScene   []string `json:"action_scene,omitempty"` // 行为场景，查询目标为行为时必填，兴趣不生效 允许值：E-COMMERCE、NEWS、APP
	ActionDays    int64    `json:"action_days,omitempty"`  // 行为天数，查询目标为行为时必填，兴趣不生效 允许值：7, 15, 30, 60, 90, 180, 365
}

type InterestActionID2WordRes struct {
	Categories []InterestActionID2WordResBase `json:"categories"` // 类目列表
	Keywords   []InterestActionID2WordResBase `json:"keywords"`   // 关键词列表
}

type InterestActionID2WordResBase struct {
	Id   int64  `json:"id"`   // 类目ID
	Name string `json:"name"` // 类目名称
	Num  int64  `json:"num"`  // 覆盖人数
}

// GetInterestActionID2Word 兴趣行为类目关键词id转词
func (client *Client) GetInterestActionID2Word(request *InterestActionID2WordReq, response *InterestActionID2WordRes) error {
	df := gout.GET(client.url(conf.API_INTEREST_ACTION_ID2WORD)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.InterestActionID2WordReqBase))
	return client.DoRequest(df, response)
}

// ---------------------------------------------------获取行为兴趣推荐关键词-----------------------------------------------

// InterestActionKeywordSuggestReq 获取行为兴趣推荐关键词
type InterestActionKeywordSuggestReq struct {
	AccessToken string `json:"access_token"`
	InterestActionKeywordSuggestReqBase
}

type InterestActionKeywordSuggestReqBase struct {
	AdvertiserId  int64    `json:"advertiser_id"`          // 广告主ID
	Id            int64    `json:"id"`                     // 类目或关键词id
	TagType       string   `json:"tag_type"`               // 查询类型：类目还是关键词 允许值：CATEGORY（类目）、KEYWORD（关键词）
	TargetingType string   `json:"targeting_type"`         // 查询目标：兴趣还是行为 允许值：ACTION（行为）、INTEREST（兴趣）
	ActionScene   []string `json:"action_scene,omitempty"` // 行为场景，查询目标为行为时必填，兴趣不生效 允许值：E-COMMERCE、NEWS、APP
	ActionDays    int64    `json:"action_days,omitempty"`  // 行为天数，查询目标为行为时必填，兴趣不生效 允许值：7, 15, 30, 60, 90, 180, 365
}

type InterestActionKeywordSuggestRes struct {
	Keywords []InterestActionKeywordSuggestResBase `json:"keywords"` // 关键词列表
}

type InterestActionKeywordSuggestResBase struct {
	Id   int64  `json:"id"`   // 类目ID
	Name string `json:"name"` // 类目名称
	Num  int64  `json:"num"`  // 覆盖人数
}

// GetInterestActionKeywordSuggest 获取行为兴趣推荐关键词
func (client *Client) GetInterestActionKeywordSuggest(request *InterestActionKeywordSuggestReq, response *InterestActionKeywordSuggestRes) error {
	df := gout.GET(client.url(conf.API_INTEREST_ACTION_KEYWORD_SUGGEST)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.InterestActionKeywordSuggestReqBase))
	return client.DoRequest(df, response)
}
