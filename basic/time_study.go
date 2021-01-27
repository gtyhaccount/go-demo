package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取本地时区的当前时间；2019-07-12 16:26:46.5359889 +0800 CST m=+0.006979101
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC()) // 获取UTC时间；2019-07-12 08:27:33.6798146 +0000 UTC

	currentTime := time.Now()
	fmt.Println(currentTime.Second())
	fmt.Println(currentTime.String()) // 2019-07-12 16:40:05.7581141 +0800 CST m=+0.030900801
	/*
		2019-07-12 16:40:05.7581141 +0800 CST m=+0.030900801
		2019-07-12 16:40:05 年月日时分秒
		.7581141 精确到纳秒
		+0800 CST 时区为+8区
		m=+0.030900801 表示当前系统启动后，流逝的时间。称为“monotonic clock reading”(单调的时钟读数)
	*/
	fmt.Println(currentTime.UnixNano()) //1562920805758114100
	fmt.Println(currentTime.Unix())     //    1562920805
	fmt.Println(currentTime.Date())
	fmt.Println(currentTime.Day())
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Month())
	fmt.Println(currentTime.Weekday())
	fmt.Println(currentTime.Year())
	fmt.Println(currentTime.YearDay())

	tm := time.Unix(currentTime.Unix(), 0)
	fmt.Println(tm)

	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano() / 1e6)

	var t2 int64
	t2 = 1593916771778
	fmt.Println(time.Unix(t2/1e3, 0))

}
