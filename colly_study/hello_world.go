package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

/*
Collector对象接受多种回调方法，有不同的作用，调用顺序：

OnRequest。请求前
OnError。请求过程中发生错误
OnResponse。收到响应后
OnHTML。如果收到的响应内容是HTML调用它。
OnXML。如果收到的响应内容是XML 调用它。写爬虫基本用不到，所以上面我没有使用它。
OnScraped。在OnXML/OnHTML回调完成后调用。不过官网写的是Called after OnXML callbacks，实际上对于OnHTML也有效，大家可以注意一下。

*/
func main() {
	// collect:收集 collector:收藏家
	collector := colly.NewCollector()

	// 事件监听
	collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
		link := element.Attr("href")
		fmt.Printf("Link found:%q -> %s\n", element.Text, link)
		collector.Visit(link)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	// 启动网页访问
	collector.Visit("http://go-colly.org/")
}
