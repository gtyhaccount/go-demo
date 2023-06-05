package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Message struct {
	ID           string
	DeviceToken  string
	IsSandbox    bool // only for iOS sandbox
	Notification Notification
	Data         string
	IsBackground bool
	CategoryID   string
}
type Notification struct {
	Title   string
	Body    string
	Android Android
	Badge   int
}
type Android struct {
	ClickAction string
}

func (m Message) Payload() string {
	b, _ := json.Marshal(&Payload{
		Id:   m.ID,
		Body: m.Data,
	})

	return string(b)
}

type Payload struct {
	Id        string `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Offset    int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Type      int32  `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Subtype   int32  `protobuf:"varint,6,opt,name=subtype,proto3" json:"subtype,omitempty"`
	Body      string `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	Extra     string `protobuf:"bytes,8,opt,name=extra,proto3" json:"extra,omitempty"`
	SenderId  string `protobuf:"bytes,9,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	ExpiresAt int64  `protobuf:"varint,10,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

// https://dev.mi.com/console/doc/detail?pId=1163
func PushMessage(n *Message, secret string) (err error) {

	params := url.Values{}
	params.Add("registration_id", n.DeviceToken)
	params.Add("title", n.Notification.Title)
	params.Add("description", n.Notification.Body)
	/*
		可选项，预定义通知栏消息的点击行为。通过设置extra.notify_effect的值以得到不同的预定义点击行为。
		“1″：通知栏点击后打开app的Launcher Activity。
		“2″：通知栏点击后打开app的任一Activity（开发者还需要传入extra.intent_uri）。
		“3″：通知栏点击后打开网页（开发者还需要传入extra.web_uri）。
	*/
	params.Add("extra.notify_effect", "2")

	params.Add("notify_id", "2")
	params.Add("extra.channel_id", "101556")
	params.Add("extra.badge", fmt.Sprintf("%d", n.Notification.Badge))
	params.Add("extra.intent_uri", n.Notification.Android.ClickAction+url.QueryEscape(n.Payload()))
	req, err := http.NewRequest("POST", "https://api.xmpush.xiaomi.com/v3/message/regid", strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Set("Authorization", "key="+secret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("xiaomi channel return code: %d, body: %s", resp.StatusCode, string(b))
	}

	var result PushResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	if result.Code != 0 {
		return fmt.Errorf("xiaomi channel return code=%d, message=%s", result.Code, result.Reason)
	}

	return
}

type PushResult struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"` // error reason
}

func main() {
	msg := &Message{
		ID:          "2023011722222",
		DeviceToken: "WmYoQN9YcaZOzNBWVE/ubwGEAOK2n0j35Qs0U9FgKVJAkGHVgv0+nq7v4CBaYuhx",
		Notification: Notification{
			Title:   "xiaomi test",
			Body:    "xiaomi test",
			Android: Android{ClickAction: ""},
		},
		Data:         "xiaomi test",
		IsBackground: false,
		IsSandbox:    false,
	}
	err := PushMessage(msg, "xlTHuV8KoU6rXFD68eIe/A==")
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

}
