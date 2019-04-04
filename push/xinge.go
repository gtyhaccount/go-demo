package push

import (
	"encoding/json"
	"fmt"
	"github.com/FrontMage/xinge"
	"github.com/FrontMage/xinge/auth"
	"github.com/FrontMage/xinge/req"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	auther := auth.Auther{"40c5e85d75390", "69de747f3557a1dfd75f627e3b320ca7"}
	pushReq, _ := req.NewPushReq(
		&xinge.Request{},
		req.Platform(xinge.PlatformAndroid),
		req.AudienceType(xinge.AdTokenList),
		req.MessageType(xinge.MsgTypeNotify),
		req.TokenList([]string{"06a74630e4f8bc8c42f095c2349ef47cc41a7b67"}), //ca20cb1e26d2a47948a09258edf1204021c9a31f
		req.Message(xinge.Message{
			Title:   "890火哥的花氧" + time.Now().String(),
			Content: "你收到了\\u201c1378\\u201d的评论回复，快去看看吧~",
			Android: &xinge.AndroidParams{
				Action: map[string]interface{}{
					"intent": "mkccommunity:///notify_detail?action={\"id\":4748," +
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
	fmt.Printf("%+v", r)
	if r.RetCode != 0 {
		fmt.Errorf("Failed rsp=%+v", r)
	}
}
