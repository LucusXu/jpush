# jpush
Jiguang Push Service Golang SDK

Production ready, full golang implementation of jiguang Push API (https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report/)

```Go
var client = jpush.NewClient("yourappSecret", "appkey")

func main() {
    var payload *Payload = NewPayload("android")
    payload.SetRegistrationId(regIds)
    result, err := client.Push(context.TODO(), payload)
}

```

### Sender APIs

- [x] Push(ctx context.Context, payload *Payload)
- [x] PushMultiTags(ctx context.Context, platform, cid string, tags, andTags, notTags []string, notice map[string]string)
- [x] BatchPush(ctx context.Context, platform string, rigs []map[string]string, notice map[string]string)

### Stats APIs

- [x] Stats(msgIds string)
- [x] VipStats(ctx context.Context, msgIds string)

### Other

- [x] GetCids(ctx context.Context, count int)
