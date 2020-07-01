package main

import (
	"encoding/json"
	"fmt"
	"github.com/ByronLeeLee/go/study/push/ios/message"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"log"
	"time"
)

func main() {

	// cert, err := certificate.FromP12File("./t/community_cn_prod.p12", "111111")
	cert, err := certificate.FromP12File("./push/ios/sa_cn_prod.p12", "111111")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "94684113a30f18e507305e39ca3b0c18cd254644073fedb2d5db135337447673"
	notification.Topic = "com.marykay.cn.xiaofu"
	//notification.Topic = "com.marykay.elearning.tw"
	//notification.Topic = "com.marykay.china.mobile"
	//notification.Topic = "com.marykay.cn.productzone"

	p := payload.NewPayload()
	//p.ContentAvailable()

	p.AlertTitle("Hello")
	p.AlertBody("Hello CF!")
	p.SoundName("{\"critical\":1,\"name\":name,\"volume\":1.0}")

	payload := &message.Payload{
		Type:      769,
		Subtype:   123,
		Timestamp: time.Now().UnixNano() / 1e6,
		Topic:     "t.u.20002867158",
		SenderId:  "20004007627",
		Body:      "6 no aps message body",
		Extra:     "extra",
		ExpiresAt: 1,
	}

	data, _ := json.Marshal(payload)
	p.Custom("data", string(data))

	notification.Payload = p // See Payload section below
	client := apns2.NewClient(cert).Production()
	// dev
	//res, err := client.Development().Push(notification)
	// prod
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
