package main

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

type AuthResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// message struct -- start

type Message struct {
	Data    string        `json:"data"`
	Android AndroidConfig `json:"android"` // android推送栏消息必填
	Token   []string      `json:"token"`
}

type AndroidConfig struct {
	/*
			0：对每个应用发送到该用户设备的离线消息只会缓存最新的一条；
			-1：对所有离线消息都缓存
			1~100：离线消息缓存分组标识，对离线消息进行分组缓存，每个应用每一组最多缓存一条离线消息；
		           如果开发者发送了10条消息，其中前5条的collapse_key为1，后5条的collapse_key为2，
		           那么待用户上线后collapse_key为1和2的分别下发最新的一条消息给最终用户。
	*/
	CollapseKey  int                 `json:"collapse_key"`
	Notification AndroidNotification `json:"notification"`
	Category     string              `json:"category"`
}

type AndroidNotification struct {
	ClickAction ClickAction `json:"click_action"`
	Title       string      `json:"title"`
	Body        string      `json:"body"`
}

type ClickAction struct {
	/*
		消息点击行为类型，取值如下：
		1：用户自定义点击行为
		2：点击后打开特定url
		3：点击后打开应用App
		4：点击后打开富媒体信息
	*/
	Type   int    `json:"type"`
	Intent string `json:"intent"`
	Url    string `json:"url"`
}

// message struct -- end

func main() {
	clientID := "100009975"
	clientSecret := "71233be71e72be8924b1fe89a4d75290"

	// reuse
	httpClient := &http.Client{}

	authToken, err := auth("client_credentials", clientID, clientSecret, httpClient)

	pushTokenArr := []string{"IQAAAACy0NDSAAAU3cDOvEkDixas7iEzYNEzoZOenBJZ2U3Tbrc-k1Tvcg4WZ_6oO1RYy2-rkUVX_bzQGDQ1k90dMrZIzJRVDxGxaT9Xnif7gZ2svQ"}
	if err != nil {
		log.Errorf("request AT error:", err)
		return
	}

	message := Message{
		Data:  "{'a':'a'}",
		Token: pushTokenArr,
		Android: AndroidConfig{
			CollapseKey: -1,
			Notification: AndroidNotification{
				Title: "hello huawei test!",
				Body:  "hello huawei test!",
				ClickAction: ClickAction{
					Type:   1,
					Intent: "mkcskin://com.mk/notify_detail?action=1",
					//Url:  "https://www.baidu.com",
				},
			},
			Category: "WORK",
		},
	}

	params := map[string]interface{}{}
	params["message"] = message
	params["validate_only"] = true //控制当前是否为测试消息，测试消息只做格式合法性校验，不会推送给用户设备, true：测试消息 false：正式消息

	paramStr, err := json.Marshal(params)
	if err != nil {
		log.Errorf("marshal error:", err)
		return
	}

	req, err := http.NewRequest("POST",
		strings.Replace("https://push-api.cloud.huawei.com/v1/[appid]/messages:send", "[appid]", clientID, 1),
		strings.NewReader(string(paramStr)))

	if err != nil {
		log.Errorf("send message,crated a request error:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	resp, err := httpClient.Do(req)
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
		log.Errorf("send message,read response Body error:", err)
		return
	}

	log.Info("send message,response Body is:", string(body))
}

//eg: "grant_type=client_credentials&client_id=10770940&client_secret=dc6282cad191bd743569b1b09bda9fbf"
func auth(grantType, clientId, clientSecret string, client *http.Client) (string, error) {
	var bufferStr bytes.Buffer
	bufferStr.WriteString("grant_type=")
	bufferStr.WriteString(grantType)
	bufferStr.WriteString("&client_id=")
	bufferStr.WriteString(clientId)
	bufferStr.WriteString("&client_secret=")
	bufferStr.WriteString(clientSecret)

	req, err := http.NewRequest("POST", "https://oauth-login.cloud.huawei.com/oauth2/v2/token", strings.NewReader(bufferStr.String()))
	if err != nil {
		log.Errorf("request AT,crated a request error:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("request AT,response error:", err)
		return "", err
	}

	// Body必须关闭，否则可能影响到后续的操作
	defer resp.Body.Close()

	if 200 != resp.StatusCode {
		log.Errorf("request AT,request error,Body:", resp)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("request AT,read response Body error:", err)
		return "", err
	}

	res := &AuthResult{}
	if err := json.Unmarshal(body, &res); err != nil {
		log.Errorf("request AT,unmarshal error:", err)
		return "", err
	}

	log.Info("request AT,response Body data is ", res.AccessToken)

	atBuffer := bytes.Buffer{}
	atBuffer.WriteString("Bearer ")
	atBuffer.WriteString(res.AccessToken)

	return atBuffer.String(), nil
}
