package jpush

import (
	"fmt"
)

const (
	MaxTimeToLive 		= 86400 * 10
	DefaultTimeToLive 	= 86400
)

// Playload对象
func NewPayload(platform string) *Payload {
	return &Payload{
		Platform:   platform,
		Cid:    	"",
		Options:    Options{},
	}
}

type Payload struct {
	Platform		string 			`json:"platform,omitempty"`				// 推送平台，必选
	Audience		Audience		`json:"audience,omitempty"`				// 推送设备，必选
	Notification	Notification	`json:"notification,omitempty"`			// 可选，通知内容体，被推送到客户端的内容，与message二者必选其一，或者并存
	Message			Message			`json:"message,omitempty"`				// 可选，消息内容体，通notification
	SmsMessage		string			`json:"sms_message,omitempty"`			// 可选，短信渠道补充送达内容体
	Options			Options			`json:"options,omitempty"`				// 可选，推送参数
	Cid				string			`json:"cid,omitempty"`					// 可选，防止重复的标识
}

// 仅针对广播推送
type BroadcastPayload struct {
	Platform		string 			`json:"platform,omitempty"`				// 推送平台，必选
	Audience		string			`json:"audience,omitempty"`				// 推送设备，广播直接写all
	Notification	Notification	`json:"notification,omitempty"`			// 可选，通知内容体，被推送到客户端的内容，与message二者必选其一，或者并存
	Message			Message			`json:"message,omitempty"`				// 可选，消息内容体，通notification
	SmsMessage		string			`json:"sms_message,omitempty"`			// 可选，短信渠道补充送达内容体
	Options			Options			`json:"options,omitempty"`				// 可选，推送参数
	Cid				string			`json:"cid,omitempty"`					// 可选，防止重复的标识
}

type Options struct {
	Sendno 			int 			`json:"sendno,omitempty"`				// 可选，API 调用标识
	TimeToLive		int 			`json:"time_to_live,omitempty"`			// 可选，离线消息最大保留时间，默认86400，最大10天，0表示不保留
	OverrideMsgId	int64			`json:"override_msg_id,omitempty"`		// 可选，如果当前的推送要覆盖之前的一条推送，这里填写前一条推送的 msg_id 就会产生覆盖效果,仅对安卓有效
	ApnsProduction  bool			`json:"apns_production,omitempty"`		// 可选，针对iOS 的 Notification 有效，true表示推送生产环境，false表示开发环境，默认true
	ApnsCollapseId	string			`json:"apns_collapse_id,omitempty"`		// 可选，更新 iOS 通知的标识符
	BigPushDuration	int				`json:"big_push_duration,omitempty"`	// 可选，缓慢推送，把原本尽可能快的推送速度，降低下来，给定的 n 分钟内，均匀地向这次推送的目标用户推送。最大值为 1400。未设置则不是定速推送。
}

type Audience struct {
	Tag				[]string 		`json:"tag,omitempty"`					// 多个标签之间是 OR 的关系，即取并集
	TagAnd			[]string 		`json:"tag_and,omitempty"`				// 多个标签之间是 AND 关系，即取交集
	TagNot			[]string 		`json:"tag_not,omitempty"`				// 多个标签之间，先取多标签的并集，再对该结果取补集
	Alias			[]string 		`json:"alias,omitempty"`				// 多个别名之间是 OR 关系，即取并集
	Segment			[]string 		`json:"segment,omitempty"`				// 多个别名之间是 OR 关系，即取并集
	Abtest			[]string 		`json:"abtest,omitempty"`				// 在页面创建的 A/B 测试的 ID。定义为数组，但目前限制是一次只能推送一个
	RegistrationId	[]string 		`json:"registration_id,omitempty"`		// 设备标识。一次推送最多 1000 个。客户端集成 SDK 后可获取到该值。
}

type Notification struct {
	Alert			string			`json:"alert,omitempty"`				// 是一个快捷定义，各平台的 alert 信息如果都一样，则可不定义。如果各平台有定义，则覆盖这里的定义
	Android			Android			`json:"android,omitempty"`				// Android 平台上的通知，JPush SDK 按照一定的通知栏样式展示
	Ios				Ios				`json:"ios,omitempty"`					// AiOS 平台上 APNs 通知结构
	Winphone		Winphone		`json:"winphone,omitempty"`				// 该通知由 JPush 服务器代理向微软的 MPNs 服务器发送,Windows Phone 平台上的通知
}

type Android struct {
	Alert 			string				`json:"alert,omitempty"`			// 必填，通知内容
	Title 			string 				`json:"title,omitempty"`			// 可选，通知标题
	BuilderId 		int 				`json:"builder_id,omitempty"`		// 可选，通知栏样式 ID
	ChannelId 		string 				`json:"channel_id,omitempty"`		// 可选，android通知channel_id
	Priority 		int 				`json:"priority,omitempty"`			// 可选，通知栏展示优先级
	Category 		string 				`json:"category,omitempty"`			// 可选，通知栏条目过滤或排序
	Style 			int 				`json:"style,omitempty"`			// 可选，通知栏样式类型
	AlertType 		int 				`json:"alert_type,omitempty"`		// 可选，通知提醒方式
	BigText 		string 				`json:"big_text,omitempty"`			// 可选，大文本通知栏样式
	Inbox 			map[string]string 	`json:"inbox,omitempty"`			// 可选，文本条目通知栏样式
	BigPicPath 		string 				`json:"big_pic_path,omitempty"`		// 可选，文本条目通知栏样式
	Extras 			map[string]string	`json:"extras,omitempty"`			// 扩展字段
	LargeIcon 		string 				`json:"large_icon,omitempty"`		// 可选，通知栏大图标
	Intent 			map[string]string 	`json:"intent,omitempty"`			// 可选，指定跳转页面
}

type Ios struct {
	Alert 				string				`json:"alert,omitempty"`				// 必填，通知内容
	Sound				string				`json:"sound,omitempty"`				// 可选，通知提示声音或警告通知
	Badge				string				`json:"badge,omitempty"`				// 可选，应用角标
	ContentAvailable	bool				`json:"content-available,omitempty"`	// 可选，推送唤醒
	MutableContent		bool				`json:"mutable-content,omitempty"`		// 可选，通知扩展
	Category			string				`json:"category,omitempty"`				// IOS 8 才支持。设置 APNs payload 中的 "category" 字段值
	Extras 				map[string]string	`json:"extras,omitempty"`				// 扩展字段
	ThreadId			string 				`json:"thread-id,omittmpey"`			// 可选，通知分组
}

type Winphone struct {
	Alert 			string				`json:"alert,omitempty"`			// 必填，通知内容
	Title 			string 				`json:"title,omitempty"`			// 可选，通知标题
	OpenPage		string 				`json:"_open_page,omittmpey"`		// 可选，点击打开的页面名称
	Extras 			map[string]string	`json:"extras,omitempty"`			// 扩展字段
}

type Message struct {
	MsgContent		string				`json:"msg_content,omitempty"`		// 必填，自定义消息，透传消息
	Title 			string 				`json:"title,omitempty"`			// 可选，消息标题
	ContentType 	string  			`json:"content_type,omitempty"`		// 可选，消息内容类型
	Extras   		map[string]string	`json:"extras,omitempty"`			// 可选，JSON 格式的可选参数
}

func (p *Payload) SetPlatform(platform string) *Payload {
	p.Platform = platform
	return p
}

func (p *Payload) SetCid(cid string) *Payload {
	p.Cid = cid
	return p
}

//*********************************options***********************************//
func (p *Payload) SetTimeToLive(ttl int) *Payload {
	if ttl < 0 {
		return p
	}
	// 10 days
	if ttl > MaxTimeToLive {
		ttl = MaxTimeToLive
	}
	// default not need set
	if ttl == DefaultTimeToLive {
		return p
	}
	p.Options.TimeToLive = ttl
	return p
}

func (p *Payload) SetSendno(sendno int) *Payload {
	p.Options.Sendno = sendno
	return p
}

func (p *Payload) SetOverrideMsgId(overrideMsgId int64) *Payload {
	p.Options.OverrideMsgId = overrideMsgId
	return p
}

func (p *Payload) SetApnsProduction(apnsProduction bool) *Payload {
	p.Options.ApnsProduction = apnsProduction
	return p
}

func (p *Payload) SetApnsCollapseId(apnsCollapseId string) *Payload {
	p.Options.ApnsCollapseId = apnsCollapseId
	return p
}

func (p *Payload) SetBigPushDuration(bigPushDuration int) *Payload {
	p.Options.BigPushDuration = bigPushDuration
	return p
}

//*************************************Audience*********************************//
func (p *Payload) SetRegistrationId(regIds []string) *Payload {
	if len(regIds) == 0 || len(regIds) > 1000 {
		panic("wrong number regIDList")
	}
	fmt.Println("regids", regIds)
	p.Audience.RegistrationId = regIds
	return p
}

func (p *Payload) SetTag(tags []string) *Payload {
	p.Audience.Tag = tags
	return p
}


func (p *Payload) SetTagAnd(tags []string) *Payload {
	p.Audience.TagAnd = tags
	return p
}

func (p *Payload) SetTagNot(tags []string) *Payload {
	p.Audience.TagNot = tags
	return p
}

func (p *Payload) SetAlias(aliases []string) *Payload {
	p.Audience.Alias = aliases
	return p
}

//*****************************IOS******************************//
func (p *Payload) SetIos(alert, sound, threadId, badge string, contentAvailable, mutableContent bool, extras map[string]string) *Payload {
	p.Notification.Ios.Alert = alert
	p.Notification.Ios.Sound = sound
	p.Notification.Ios.ThreadId = threadId
	p.Notification.Ios.Badge = badge
	p.Notification.Ios.ContentAvailable = contentAvailable
	p.Notification.Ios.MutableContent = mutableContent
	p.Notification.Ios.Extras = extras
	return p
}

func NewIosNotification() *Ios {
	var extras = make(map[string]string, 0)
	return &Ios{
		Alert: "",
		Sound: "",
		ThreadId: "",
		Badge: "+1",
		ContentAvailable: false,			// 尽量别用true, 苹果会从后台自动打开推送
		MutableContent: true,
		Extras: extras,
	}
}

func (p *Payload) SetIosAlert(alert string) *Payload {
	p.Notification.Ios.Alert = alert
	return p
}

func (p *Payload) SetIosSound(sound string) *Payload {
	p.Notification.Ios.Sound = sound
	return p
}

func (p *Payload) SetIosThreadId(threadId string) *Payload {
	p.Notification.Ios.ThreadId = threadId
	return p
}

func (p *Payload) SetIosBadge(badge string) *Payload {
	p.Notification.Ios.Badge = badge
	return p
}


func (p *Payload) SetIosContentAvailable(contentAvailable bool) *Payload {
	p.Notification.Ios.ContentAvailable = contentAvailable
	return p
}

func (p *Payload) SetIosMutableContent(mutableContent bool) *Payload {
	p.Notification.Ios.MutableContent = mutableContent
	return p
}

func (p *Payload) SetIosExtras(extras map[string]string) *Payload {
	p.Notification.Ios.Extras = extras
	return p
}

//*************************Android****************************************//
func NewAndroidNotification() *Android {
	var extras = make(map[string]string, 0)
	var inbox = make(map[string]string, 0)
	var intent = make(map[string]string, 0)

	return &Android{
		Alert: "",
		Title: "",
		BuilderId: 1,
		Priority: 0,
		Style: 0,
		AlertType: -1,
		Inbox: inbox,
		Intent: intent,
		Extras: extras,
	}
}

func (p *Payload) SetAndroidAlert(alert string) *Payload {
	p.Notification.Android.Alert = alert
	return p
}

func (p *Payload) SetAndroidTitle(title string) *Payload {
	p.Notification.Android.Title = title
	return p
}

func (p *Payload) SetAndroidExtras(extras map[string]string) *Payload {
	p.Notification.Android.Extras = extras
	return p
}

//*************************message***************************************//
func (p *Payload) SetMessage(msgContent, title, contentType string, extras map[string]string) *Payload {
	p.Message.MsgContent = msgContent
	p.Message.Title = title
	p.Message.ContentType = contentType
	p.Message.Extras = extras
	return p
}