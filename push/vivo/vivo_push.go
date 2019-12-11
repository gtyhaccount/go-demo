package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	timestamp := time.Now().UnixNano() / 1e6
	//appId := 11788
	//appKey := "28fc0f98-e029-4c9b-80b1-456544a4ca87"
	//appSecret := "09d3cb69-b41b-40b2-91b2-936d1a1c2095"

	// 1. vivo auth
	//at, _ := vivoAuth(timestamp, appId, appKey, appSecret)
	//fmt.Println(at)
	// 2. device push
	at := "e3039c1c-8147-4136-87b1-9fb48ff68c04"
	regId := "15554962178631178811065"
	if err := vivoPush(at, regId, "Hello world:"+strconv.FormatInt(timestamp, 10), "Hello xiao xiang:"+strconv.FormatInt(timestamp, 10)); err != nil {
		fmt.Errorf("push error:%s", err)
	}

}

func vivoAuth(timestamp int64, appId int, appKey, appSecret string) (accessToken string, err error) {
	sign := md5.Sum([]byte(strconv.Itoa(appId) + appKey + strconv.FormatInt(timestamp, 10) + appSecret))

	signStr := fmt.Sprintf("%x", sign)

	param := make(map[string]interface{})
	param["appId"] = appId
	param["appKey"] = appKey
	param["timestamp"] = timestamp
	param["sign"] = signStr

	paramJson, _ := json.Marshal(param)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api-push.vivo.com.cn/message/auth", strings.NewReader(string(paramJson)))
	if err != nil {
		return "", errors.New(fmt.Sprintf("request AT,crated a request error:%s", err))
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(fmt.Sprintf("request AT,response error:%s", err))
	}

	defer resp.Body.Close()

	if 200 != resp.StatusCode {
		return "", errors.New(fmt.Sprintf("request AT,request error,body:%s", err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("request AT,read response body error:%s", err))
	}

	type Result struct {
		Result    int    `json:"result"`
		Desc      string `json:"desc"`
		AuthToken string `json:"authToken"`
	}

	res := &Result{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", errors.New(fmt.Sprintf("request AT,unmarshal error:%s", err))
	}

	return res.AuthToken, nil
}

func vivoPush(accessToken, regId, title, content string) error {
	param := make(map[string]interface{})
	param["regId"] = regId
	// 通知类型 1:无，2:响铃，3:振动，4:响铃和振动
	param["notifyType"] = 1
	param["title"] = title
	param["content"] = content
	// 点击跳转类型 1：打开 APP 首页 2：打开链接 3：自定义 4:打开 app 内指定页面
	param["skipType"] = 2
	param["skipContent"] = "mkskin:///notify_detail?action=paramxxx"
	// 用户请求唯一标识 最大 64 字符
	param["requestId"], _ = uuid.NewV4()
	paramJson, _ := json.Marshal(param)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api-push.vivo.com.cn/message/send", strings.NewReader(string(paramJson)))
	if err != nil {
		return errors.New(fmt.Sprintf("send message,crated a request error:%s", err))
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("authToken", accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("send message,response error:%s", err))
	}

	defer resp.Body.Close()

	if 200 != resp.StatusCode {
		return errors.New(fmt.Sprintf("send message,request error,body:%s", err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("send message,read response body error:%s", err))
	}

	type Result struct {
		Result int    `json:"result"`
		Desc   string `json:"desc"`
		taskId string `json:"taskId"`
	}

	res := &Result{}
	if err := json.Unmarshal(body, &res); err != nil {
		return errors.New(fmt.Sprintf("send message,unmarshal error:%s", err))
	}

	log.WithFields(log.Fields{
		"regId": regId,
	}).Debug("response body:%s", string(body))
	fmt.Println("response body:%s", string(body))

	return nil
}
