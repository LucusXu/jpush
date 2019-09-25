# jpush
极光推送服务 Golang SDK

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

### Stats APIs

- [x] Stats(msgIds string)