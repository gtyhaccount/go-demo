package main

import (
	"encoding/json"
	"fmt"
	"github.com/FrontMage/xinge"
	xingeAuth "github.com/FrontMage/xinge/auth"
	"github.com/FrontMage/xinge/req"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	auther := xingeAuth.Auther{"6f306fc27dc5a", "51505f48397bc9dfe21b9187b8f63e69"}
	pushReq, _ := req.NewPushReq(
		&xinge.Request{},
		req.Platform(xinge.PlatformAndroid),
		req.AudienceType(xinge.AdTokenList),
		req.MessageType(xinge.MsgTypeMessage),
		//req.MessageType(xinge.MsgTypeNotify),
		req.TokenList([]string{"f353daf3a6a219830bf5ed41385e0aea1cc37c70"}), //ca20cb1e26d2a47948a09258edf1204021c9a31f
		req.Message(xinge.Message{
			Title:   "Community Hello 祥哥" + time.Now().String(),
			Content: "Community 你收到了评论回复，快去看看吧~",
			Android: &xinge.AndroidParams{
				//Action: map[string]interface{}{
				//	"intent": "mkcintouch:///notify_detail?action={\"id\":666," +
				//		"\"topic\":\"t.u.1001000048\",\"offset\":31,\"timestamp\":1543989812303,\"type\":769," +
				//		"\"body\":\"{\"A\":[\"42894C89-B7CA-48F2-945B-0FA6ADF7131E\",\"873a6491-5b2d-4fee-b14c-ae7e1f360f18\",\"\",\"UGC\",\"我的评论被回复1\",\"你收到了\\u201c1378\\u201d的评论回复，快去看看吧~\"],\"H\":\"Community\",\"M\":\"ArticleComment\",\"N\":\"你收到了\\u201c1378\\u201d的评论回复，快去看看吧~\"}\",\"sender_id\":\"community.devicemgr\"}",
				//	"action_type": 1,
				//},
				CustomContent: map[string]string{
					"intent": "mkcintouch:///notify_detail?action={\"id\":666," +
						"\"topic\":\"t.u.1001000048\",\"offset\":31,\"timestamp\":1543989812303,\"type\":769," +
						"\"body\":\"{\"A\":[\"42894C89-B7CA-48F2-945B-0FA6ADF7131E\",\"873a6491-5b2d-4fee-b14c-ae7e1f360f18\",\"\",\"UGC\",\"我的评论被回复1\",\"你收到了\\u201c1378\\u201d的评论回复，快去看看吧~\"],\"H\":\"Community\",\"M\":\"ArticleComment\",\"N\":\"你收到了\\u201c1378\\u201d的评论回复，快去看看吧~\"}\",\"sender_id\":\"community.devicemgr\"}",
				},
			},
		}),
	)

	auther.Auth(pushReq)

	c := &http.Client{}
	rsp, _ := c.Do(pushReq)

	defer rsp.Body.Close()
	body, _ := ioutil.ReadAll(rsp.Body)

	r := &xinge.CommonRsp{}
	json.Unmarshal(body, r)

	if r.PushID == "" {
		fmt.Println("xinge push message error,response body:" + string(body))
		return
	}

	fmt.Println("xinge push message end,response body:" + string(body))
}
