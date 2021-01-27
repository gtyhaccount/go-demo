package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/**
oppo 推送服务测试
注意：任何值字符串都是需要 urlencode 编码
*/
type Result struct {
	Code    int
	message string
	Data    interface{}
}
type Message struct {
	TargetType   int    `json:"target_type"`
	TargetValue  string `json:"target_value"`
	Notification struct {
		// App开发者自定义消息Id，OPPO推送平台根据此ID做去重处理，对于广播推送相同appMessageId只会保存一次，对于单推相同appMessageId只会推送一次
		AppMessageId string `json:"app_message_id"`
		Title        string `json:"title"`
		SubTitle     string `json:"sub_title"`
		Content      string `json:"content"`

		// 非必填
		ClickActionType     int    `json:"click_action_type"`
		ClickActionActivity string `json:"click_action_activity"`
		ClickActionUrl      string `json:"click_action_url"`
		ActionParameters    string `json:"action_parameters"`
	}
}

func main() {
	timestamp := time.Now().UnixNano() / 1e6
	uuidstr := uuid.NewV4().String()

	// 2. 单推
	authToken, _ := auth()
	message := &Message{
		TargetType:  2,
		TargetValue: "CN_8bee29f41ac599362b0c2c7246d5cf4d",
	}
	message.Notification.AppMessageId = uuidstr
	// title max length is 32
	message.Notification.Title = "Hello 祥哥"
	//message.Notification.Title = "Hello Oppo" + strconv.FormatInt(timestamp, 10)
	// sub_title max length is 10
	//message.Notification.SubTitle = "push test"
	// content max length is 200
	message.Notification.Content = "mkcintouch:///notify_detail?action=" + "6人处于DIQ考核期" + strconv.FormatInt(timestamp, 10)
	//message.Notification.Content = "ClickActionType is 5.Hello World," + uuidstr

	message.Notification.ClickActionType = 5
	//message.Notification.ClickActionActivity = ""
	message.Notification.ClickActionUrl = "mkcskin://notify_detail?action="
	//message.Notification.ActionParameters = "{\"key1\":\"value1\",\"key2\":\"value2\"}"

	messageJsonStr, err := json.Marshal(message)
	if err != nil {
		log.Errorf("marshal error:", err)
		return
	}

	params := url.Values{}
	params.Add("message", string(messageJsonStr))

	log.Info("messageStr:", string(messageJsonStr))
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.push.oppomobile.com/server/v1/message/notification/unicast", strings.NewReader(params.Encode()))

	if err != nil {
		log.Errorf("send message,crated a request error:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Set("auth_token", authToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("send message,response error:", err)
		return
	}

	defer resp.Body.Close()

	if 200 != resp.StatusCode {
		log.Errorf("send message,response code not 200:", resp)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("send message,read response body error:", err)
		return
	}

	log.Info("send message,response body is:", string(body))

	res := &Result{}
	if err := json.Unmarshal(body, &res); err != nil {
		log.Errorf("send message,unmarshal error:", err)
		return
	}

	if 0 != res.Code {
		log.Errorf("send message,error:", res.message)
		return
	}

	resData := res.Data.(map[string]interface{})
	fmt.Println("messageId is:", resData["messageId"].(string))

}

func auth() (string, error) {
	// 1. 获取access_token(之后每次调用都需要access_token)
	timestamp := time.Now().UnixNano() / 1e6
	appKey := "9zDsNzYZk8gso8Kg0G08o4Gcw"
	masterSecret := "8Bb48daf56a00FD91dEAC93ddbb2053b"
	sign := sha256.Sum256([]byte(appKey + strconv.FormatInt(timestamp, 10) + masterSecret))

	signStr := fmt.Sprintf("%x", sign)

	values := url.Values{}
	values.Add("app_key", appKey)
	values.Add("sign", signStr)
	values.Add("timestamp", strconv.FormatInt(timestamp, 10))

	client := &http.Client{}

	//jsonStr, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", "https://api.push.oppomobile.com/server/v1/auth", strings.NewReader(values.Encode()))
	if err != nil {
		log.Errorf("request AT,crated a request error:", err)
		return "", nil
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("request AT,response error:", err)
		return "", nil
	}

	// Body必须关闭，否则可能影响到后续的操作
	defer resp.Body.Close()

	if 200 != resp.StatusCode {
		log.Errorf("request AT,request error,body:", resp)
		return "", nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("request AT,read response body error:", err)
		return "", nil
	}

	log.Info("request AT,response body is:", string(body))

	res := &Result{}
	if err := json.Unmarshal(body, &res); err != nil {
		log.Errorf("request AT,unmarshal error:", err)
		return "", nil
	}

	log.Info("request AT,response body data is ", res.Data)

	resData := res.Data.(map[string]interface{})

	fmt.Printf("request AT,auth_token is %s\n", resData["auth_token"])

	return resData["auth_token"].(string), nil
}
