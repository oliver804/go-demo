package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func getHtml(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if nil != err {
		log.Fatal("error:", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", data)
}

type Proxy struct {
	//IP	PORT	匿名度	类型	位置	响应速度	最后验证时间	付费方式
	ip                 string
	port               int
	anonymity          string
	httpType           string
	addr               string
	respSpeed          string
	latestCheckTime    time.Time
	latestCheckTimeStr string
	chargeType         string
}

func parseHtml(html string) (list.List, error) {
	doc, err := htmlquery.Parse(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}
	trs := htmlquery.Find(doc, "//*[@id=\"list\"]/div[1]/table/tbody/tr")
	list := list.New()
	for _, tr := range trs {
		fmt.Println(tr)
		tds := htmlquery.Find(tr, "//td")
		var proxy Proxy
		proxy.ip = htmlText(tds[0])
		proxy.port, err = strconv.Atoi(htmlText(tds[1]))
		proxy.anonymity = htmlText(tds[2])
		proxy.httpType = htmlText(tds[3])
		proxy.addr = htmlText(tds[4])
		proxy.respSpeed = htmlText(tds[5])
		proxy.latestCheckTime, _ = time.Parse(time.DateTime, htmlText(tds[6]))
		proxy.latestCheckTimeStr = proxy.latestCheckTime.Format(time.DateTime)
		proxy.chargeType = htmlText(tds[7])
		list.PushFront(proxy)
	}
	return *list, nil
}

func htmlText(node *html.Node) string {
	return htmlquery.OutputHTML(node, false)
}

func main() {
	url := "https://www.kuaidaili.com/free/inha/"
	html := getHtml(url)
	list, err := parseHtml(html)
	if err != nil {
		log.Fatal(err)
	}
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
