package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

/**
总结：
1. json字符串解析时，需要一个“接收体”接受解析后的数据，且  Unmarshal时接收体必须传递指针  。否则解析虽不报错，但数据无法赋值到接受体中。

2. 解析时，接收体可自行定义。json串中的key自动在接收体中寻找匹配的项进行赋值。匹配规则是：
	2.1 先查找与key一样的json标签，找到则赋值给该标签对应的变量(如Name)。
	2.2 没有json标签的，就从上往下依次查找变量名与key一样的变量，如Age。或者变量名忽略大小写后与key一样的变量。如HIgh，Class。
  第一个匹配的就赋值，后面就算有匹配的也忽略。(前提是该变量必需是可导出的，即首字母大写)。

3. 不可导出的变量无法被解析（如sex变量，虽然json串中有key为sex的k-v，解析后其值仍为nil,即空值）

4. 当接收体中存在json串中匹配不了的项时，解析会自动忽略该项，该项仍保留原值。如变量Test，保留空值nil。

5. 没有指定变量的具体类型，json自动将value为  复合结构   的数据解析为  map[string]interface{} 类型的项。

*/

type Result2 struct {
	Code    int         `json:"code"`
	message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {

	// 示例一
	str := "{\"id\":2076190,\"topic\":\"t.u.20002868550\",\"offset\":42,\"timestamp\":1545901211198,\"type\":769,\"subtype\":123,\"body\":\"{\\\"appId\\\":\\\"MyMessage\\\",\\\"targetAppId\\\":null,\\\"title\\\":\\\"测试消息体是否能正确解析-title\\\",\\\"subTitle\\\":\\\"测试消息体是否能正确解析-content\\\",\\\"coverImageUrl\\\":\\\"COVER_IMAGE_URL\\\",\\\"videoUrl\\\":\\\"COVER_VIDEO_URL\\\",\\\"targetUrl\\\":\\\"\\\",\\\"targetAppParameters\\\":\\\"\\\",\\\"expiredAt\\\":null}\",\"sender_id\":\"mybiz-message-admin\"}"

	title, body := extractBody(str)
	fmt.Println(title)
	fmt.Println(body)

	// 示例二
	str2 := "{\"code\":0,\"data\":{\"auth_token\":\"2e865fe0-af8b-663c-8d0c-a698d58f1168\",\"create_time\":1551151794582},\"message\":\"Success\"}"

	result2 := &Result2{} // 注意这里是指针类型，要不然unmarshal方法会报错
	if err := json.Unmarshal([]byte(str2), result2); err != nil {
		log.Errorf("unmarshal error:", err)
		return
	}

	/*
		注意:1. result2这个对象里面以小写字母开头的属性message，没有成功的设置值，除非将该属性首字母大写变为“可导出”的;
		     2. interface{}类型的属性Data,unmarshal方法反序列化后默认是用map[string]interface{}接收的;
	*/
	data := result2.Data.(map[string]interface{}) // 强制转化为map[string]interface{}

	fmt.Println("auth_token is:", data["auth_token"])

	type Result struct {
		result int    `json:"result"`
		desc   string `json:"desc"`
		taskId string `json:"taskId"`
	}

	str3 := "{\"result\":0,\"desc\":\"请求成功\",\"taskId\":\"569226442568265728\"}"

	var res Result
	if err := json.Unmarshal([]byte(str3), &res); err != nil {
		fmt.Println(err)
	}
	fmt.Println("taskId:", res.taskId)

	var jsonBlob = []byte(` [ 
        { "Name" : "Platypus" , "Order" : "Monotremata" } , 
        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" } 
    ] `)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)

}

func extractBody(msg string) (title, body string) {
	var result map[string]interface{}
	var result2 map[string]string
	err := json.Unmarshal([]byte(msg), &result)
	if err == nil {
		msgBody := result["body"].(string)
		err = json.Unmarshal([]byte(msgBody), &result2)

		if err == nil {
			title = result2["title"]
			body = result2["subTitle"]
		} else {
			body = msgBody
			log.Error(err)
		}
	} else {
		body = msg
		log.Error(err)
	}
	return
}
