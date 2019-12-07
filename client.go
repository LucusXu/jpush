package jpush

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"fmt"
	"encoding/json"
	"bytes"
)

type JiPush struct {
	appKey 	string
	appSecret   	string
	pushHost        string
	reportHost      string
}

func NewClient(appSecret string, appKey string) *JiPush {
	return &JiPush{
		appKey: 		appKey,
		appSecret:   	appSecret,
		pushHost:       PushProductionHost,
		reportHost:     ReportProductionHost,
	}
}

// 推送
func (ji *JiPush) Push(ctx context.Context, payload *Payload) (*PushResult, error) {
	pushUrl := ji.pushHost + RegURL
	resBytes, err := ji.doPost(ctx, pushUrl, payload)

	var result PushResult
	errJson := json.Unmarshal(resBytes, &result)
	if errJson != nil {
		return nil, errJson
	}
	if err != nil {
		return &result, err
	}
	return &result, nil
}

/**
 *   批量推送多设备
 */
func (jc *JiPush) PushRigsList(ctx context.Context, platform, cid string, rigs []string, notice map[string]string) (*PushResult, error) {
	if len(rigs) == 0 {
		fmt.Println("rigs empty")
		return nil, nil
	}

	var payload *Payload = NewPayload(platform)
	if cid != "" {
		payload.SetCid(cid)
	} else {
		var cids []string
		cids, err := jc.GetCids(ctx, 1)
		if err != nil {
			fmt.Println("get cid error", err)
		}

		if len(cids) > 0 {
			payload.SetCid(cids[0])
		}
	}

	// tokens
	payload.SetRegistrationId(rigs)
	var ttl, _ = strconv.Atoi(notice["ttl"])
	if ttl > 0 {
		// 最大存活时间s
		payload.SetTimeToLive(ttl)
	}

	var title = notice["title"]
	var alert = notice["content"]
	var custom = notice["custom"]
	var extra = make(map[string]string, 0)
	if custom != "" {
		err := json.Unmarshal([]byte(custom), &extra)
		if err != nil {
			fmt.Println("json unmarshal error ", err)
		}
	}
	if notice["image"] != "" {
		extra["image"] = notice["image"]
	}

	var android *Android = NewAndroidNotification()
	payload.Notification.Android = *android
	payload.SetAndroidAlert(alert)
	payload.SetAndroidTitle(title)
	payload.SetAndroidExtras(extra)

	var ios *Ios = NewIosNotification()
	payload.Notification.Ios = *ios
	payload.SetIosAlert(alert)

	if platform == "ios" {
		if notice["sound"] != "" {
			payload.SetIosSound(notice["sound"])
		} else {
			payload.SetIosSound("default")
		}

		if title != "" {
			extra["title"] = title
		}

		var threadId = notice["thread_id"]
		if threadId == "" {
			payload.SetIosThreadId("default")
		} else {
			payload.SetIosThreadId(threadId)
		}
		payload.SetIosExtras(extra)
	}

	var msgContent = notice["msgContent"]
	var msgTitle = notice["msgTitle"]
	var msgType = notice["msgType"]
	var msgExtras = notice["msgExtras"]
	if msgContent != "" {
		var msgExtrasArray = make(map[string]string, 0)
		if msgExtras != "" {
			err := json.Unmarshal([]byte(msgExtras), &msgExtrasArray)
			if err != nil {
				fmt.Println("json unmarshal extras error ", err)
			}
		}
		payload.SetMessage(msgContent, msgTitle, msgType, msgExtrasArray)
	}

	result, err := jc.Push(ctx, payload)
	return result, err
}

// stats
func (ji *JiPush) Stats(ctx context.Context, msgIds string) ([]string, error) {
	var statusUrl = ji.reportHost + MessagesStatusURL + "?msg_ids=" + msgIds
	sBytes, err := ji.doGet(ctx, statusUrl, "")
	if err != nil {
		return nil, err
	}
	var result CidResult
	err = json.Unmarshal(sBytes, &result)
	if err != nil {
		return nil, err
	}
	return result.CidList, nil
}

// stats
func (ji *JiPush) VipStats(ctx context.Context, msgIds string) ([]MessageResult, error) {
	var statusUrl = ji.reportHost + VipMessagesStatusURL + "?msg_ids=" + msgIds
	sBytes, err := ji.doGet(ctx, statusUrl, "")
	if err != nil {
		return nil, err
	}
	fmt.Println("bytes", string(sBytes))
	var result []MessageResult
	err = json.Unmarshal(sBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/**
 * batch get cids
 */
func (ji *JiPush) GetCids(ctx context.Context, count int) ([]string, error) {
	var cidUrl = ji.pushHost + CIDURL + "?count=" + strconv.Itoa(count)
	strBytes, err := ji.doGet(ctx, cidUrl, "")
	if err != nil {
		fmt.Println("get cid error", err)
		return nil, err
	}
	var result CidResult
	err = json.Unmarshal(strBytes, &result)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return nil, err
	}
	return result.CidList, nil
}

func (ji *JiPush) doGet(ctx context.Context, url string, params string) ([]byte, error) {
	var result []byte
	var req *http.Request
	var res *http.Response
	var err error
	req, err = http.NewRequest("GET", url+params, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(ji.appKey, ji.appSecret)
	client := &http.Client{}
	res, err = ctxhttp.Do(ctx, client, req)
	if res.Body == nil {
		panic("jpush response is nil")
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("network error," + strconv.Itoa(res.StatusCode))
	}
	result, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ji *JiPush) doPost(ctx context.Context, url string, payload *Payload) ([]byte, error) {
	var result []byte
	var req *http.Request
	var res *http.Response
	var err error
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(payload)

	req, err = http.NewRequest("POST", url, buf)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(ji.appKey, ji.appSecret)
	client := &http.Client{}
	tryTime := 0
tryAgain:
	res, err = ctxhttp.Do(ctx, client, req)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, err
		default:
		}
		tryTime += 1
		if tryTime < PostRetryTimes {
			goto tryAgain
		}
		return nil, err
	}
	if res.Body == nil {
		panic("jpush response is nil")
	}
	defer res.Body.Close()

	result, err = ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("jpush push post ioutil.ReadAll err:", err)
		return result, err
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("jpush push post fail:", err, string(result))
		return result, errors.New("network error," + strconv.Itoa(res.StatusCode) + " result:" + string(result))
	}

	return result, nil
}

/**
 *  多标签推送
 */
func (jc *JiPush) PushMultiTags(ctx context.Context, platform, cid string, tags, andTags, notTags []string, notice map[string]string) (*PushResult, error) {
	if len(tags) == 0 && len(andTags) == 0 && len(notTags) == 0 {
		return nil, errors.New("tag empty error")
	}

	if len(tags) > 20 || len(andTags) > 20 || len(notTags) > 20 {
		return nil, errors.New("tag count more than 20 error")
	}

	var payload *Payload = NewPayload(platform)
	if cid != "" {
		payload.SetCid(cid)
	} else {
		var cids []string
		cids, err := jc.GetCids(ctx, 1)
		if err != nil {
			fmt.Println("get cid error", err)
		}

		if len(cids) > 0 {
			payload.SetCid(cids[0])
		}
	}

	// 多标签都支持
	if len(tags) > 0 {
		payload.SetTag(tags)
	}

	if len(andTags) > 0 {
		payload.SetTagAnd(andTags)
	}

	if len(notTags) > 0 {
		payload.SetTagNot(notTags)
	}

	var ttl, _ = strconv.Atoi(notice["ttl"])
	if ttl > 0 {
		// 最大存活时间s
		payload.SetTimeToLive(ttl)
	}

	var title = notice["title"]
	var alert = notice["content"]
	var custom = notice["custom"]
	var extra = make(map[string]string, 0)
	if custom != "" {
		err := json.Unmarshal([]byte(custom), &extra)
		if err != nil {
			fmt.Println("json unmarshal error ", err)
		}
	}
	if notice["image"] != "" {
		extra["image"] = notice["image"]
	}

	var android *Android = NewAndroidNotification()
	payload.Notification.Android = *android
	payload.SetAndroidAlert(alert)
	payload.SetAndroidTitle(title)
	payload.SetAndroidExtras(extra)

	var ios *Ios = NewIosNotification()
	payload.Notification.Ios = *ios
	payload.SetIosAlert(alert)

	if platform == "ios" {
		if notice["sound"] != "" {
			payload.SetIosSound(notice["sound"])
		} else {
			payload.SetIosSound("default")
		}

		if title != "" {
			extra["title"] = title
		}

		var threadId= notice["thread_id"]
		if threadId == "" {
			payload.SetIosThreadId("default")
		} else {
			payload.SetIosThreadId(threadId)
		}
		payload.SetIosExtras(extra)
	}

	var msgContent = notice["msgContent"]
	var msgTitle = notice["msgTitle"]
	var msgType = notice["msgType"]
	var msgExtras = notice["msgExtras"]
	if msgContent != "" {
		var msgExtrasArray = make(map[string]string, 0)
		if msgExtras != "" {
			err := json.Unmarshal([]byte(msgExtras), &msgExtrasArray)
			if err != nil {
				fmt.Println("json unmarshal extras error ", err)
			}
		}
		payload.SetMessage(msgContent, msgTitle, msgType, msgExtrasArray)
	}

	result, err := jc.Push(ctx, payload)
	return result, err
}
