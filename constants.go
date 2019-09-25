package jpush

const (
	PushProductionHost 			= "https://api.jpush.cn"					// 推送域名
	ReportProductionHost 		= "https://report.jpush.cn"					// 统计域名
)

const (
	RegURL                      = "/v3/push"                		  	// 推送url
	CIDURL 						= "/v3/push/cid"						// 获取推送的cid
)

const (
	MessagesStatusURL  			= "/v3/received/detail"   				// 送达统计详情, 最多传100个msg_id
	VipMessagesStatusURL  		= "/v3/messages/detail"   				// 送达状态查询
)

var (
	PostRetryTimes = 3
)