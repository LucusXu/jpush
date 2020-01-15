package jpush

// getCid返回结果
type CidResult struct {
	CidList		[]string `json:"cidlist,omitempty"`
}

// 推送返回结果
type PushResult struct {
	SendNo  string `json:"sendno"`
	MsgId   string `json:"msg_id"`
	Error   Error  `json:"error"`
}

// 推送返回错误
type Error struct {
	Code 	int 	`json:"code"`
	Message string	`json:"message"`
}

// 推送返回结果
type StatsResult struct {
	AndroidPnsSent 	int 	`json:"android_pns_sent"`			// Android厂商用户推送到厂商服务器成功数
	IosApnsReceived int 	`json:"ios_apns_received"`			// iOS 通知送达到设备
	IosApnsSent 	int 	`json:"ios_apns_sent"`				// 通知推送到 APNs 成功
	IosMsgReceived 	int 	`json:"ios_msg_received"`			// iOS 自定义消息送达数
	JpushReceived 	int 	`json:"jpush_received"`				// 极光通道用户送达数
	MsgId 			string	`json:"msg_id"`						// 消息id
	WpMpnsSent 		int		`json:"wp_mpns_sent"`				// winphone通知送达
}

// 推送返回结果
type MessageResult struct {
	AndroidPns 	AndroidPns 		`json:"android_pns"`		// Android厂商通道统计数据, 走厂商通道下发统计数据
	Jpush 		JpushStats		`json:"jpush"`				// 极光通道统计数据，走极光通道下发的普通Android用户通知/自定义消息 以及 iOS用户自定义
	Ios			IosStats		`json:"ios"`				// iOS 统计数据
	MsgId 		string 			`json:"msg_id"`				// 消息id
	Winphone 	WinphoneStats 	`json:"winphone"`			// Winphone 统计数据
}

type AndroidPns struct {
	PnsSent		int				`json:"pns_sent"`
	PnsTarget	int				`json:"pns_target"`
	FcmDetail	AndroidStats	`json:"fcm_detail"`
	HwDetail	AndroidStats	`json:"hw_detail"`
	MzDetail	AndroidStats	`json:"mz_detail"`
	OppoDetail	AndroidStats	`json:"oppo_detail"`
	VivoDetail	AndroidStats	`json:"vivo_detail"`
	XmDetail	AndroidStats	`json:"xm_detail"`
}

type AndroidStats struct {
	Sent 	int `json:"sent"`			// 推送成功数
	Target 	int `json:"target"`			// 推送目标数
}

type IosStats struct {
	ApnsClick		int `json:"apns_click"`				// 通知点击数
	ApnsReceived	int	`json:"apns_received"`			// APNs 服务器下发到设备成功
	ApnsSent		int	`json:"apns_sent"`				// APNs 通知成功推送数
	ApnsTarget		int	`json:"apns_target"`			// APNs 通知推送目标数
	MsgReceived		int	`json:"msg_received"`			// 自定义消息送达数
	MsgTarget		int	`json:"msg_target"`				// 自定义消息目标数
}

type WinphoneStats struct {
	Click		int `json:"click"`					// MPNs 通知用户点击数
	MpnsSent	int	`json:"mpns_sent"`				// MPNS 通知成功推送数
	MpnsTarget	int	`json:"mpns_target"`			// MPNs 通知推送目标数
}

type JpushStats struct {
	Click		int `json:"click"`				// 用户点击数
	MsgClick	int	`json:"msg_click"`			// 自定义消息点击数
	OnlinePush	int	`json:"online_push"`		// 在线推送数
	Received	int	`json:"received"`			// 推送送达数
	Target		int	`json:"target"`				// 推送目标数
}

// 推送返回结果
type BatchPushResult struct {
	MessageId	int	    `json:"message_id"`
	Code		int     `json:"code"`
	Message	    string  `json:"message"`
	Error       Error   `json:"error"`
}
