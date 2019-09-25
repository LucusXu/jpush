package jpush

import (
	"testing"
	"golang.org/x/net/context"
	"fmt"
	"time"
	"strconv"
	"encoding/json"
)

var appKey string = "testappkey"
var secret string = "testsecret"

var client = NewClient(secret, appKey)
var regIds = []string{"testregid"}

var title = "标题"
var content = "推送内容"
var msgId ="testmsgid"

// 获取cid
func TestJiPush_GetCids(t *testing.T) {
	result, err := client.GetCids(context.TODO(), 10)
	if err != nil {
		t.Errorf("TestJiPush_GetCids failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
	str, err := json.Marshal(result)
	fmt.Println("getCids:", string(str))
}

// 多设备推送
func TestJiPush_Push(t *testing.T) {
	var notice = make(map[string]string)
	notice["title"] = title
	notice["content"] = content
	notice["ttl"] = "60000"

	threadId := time.Now().Unix()
	notice["thread_id"] = strconv.FormatInt(threadId, 10)
	fmt.Println("notice", notice)

	result, err := client.PushRigsList(context.TODO(), "ios", "", regIds, notice)

	if err != nil {
		t.Errorf("TestJiPush_Push failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
	str, err := json.Marshal(result)
	fmt.Println("push_end:msgid:", string(str), result.MsgId)
}

// 多设备推送
func TestJiPush_VipStats(t *testing.T) {
	result, err := client.VipStats(context.TODO(), msgId)

	fmt.Println("res", result)
	if err != nil {
		t.Errorf("TestJiPush_Push failed :%v\n", err)
	}
	fmt.Println("res0", result[0])
}