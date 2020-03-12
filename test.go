package main

import (
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	t, err := time.Parse("1/2/2006 3:4:5 PM", "01/03/2020 09:37:15 AM")
	if err != nil {
		log.Errorf("t parse error:%s", err)
	}

	log.Info(t)
	log.Info(time.Now().Unix())
	log.Info(time.Now().UnixNano())
	a
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("efeb6ff8d09f4a4ad763c0c4:f5e223e25c8e3da931745919")))

}
