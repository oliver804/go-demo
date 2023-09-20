package main

import (
	"container/list"
	"fmt"
	"go-demo/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func getPapers() list.List {
	var indexUrl = "https://top.baidu.com/board?tab=realtime"
	req, _ := http.NewRequest(http.MethodGet, indexUrl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	html := fmt.Sprintf("%s", data)
	return convert(html)
}

func convert(html string) list.List {
	doc, _ := htmlquery.Parse(strings.NewReader(html))
	divs := htmlquery.Find(doc, "//*[@id=\"sanRoot\"]/main/div[2]/div/div[2]/div")
	list := list.New()
	for _, div := range divs {
		tagA := htmlquery.FindOne(div, "//*[@class=\"content_1YWBm\"]/a")
		hotDiv := htmlquery.FindOne(div, "//*[@class=\"hot-index_1Bl1a\"]")
		hotStr := htmlText(hotDiv)
		var paper model.Paper
		paper.Title = htmlText(tagA)
		paper.Url = htmlquery.SelectAttr(tagA, "href")
		paper.Hot, _ = strconv.Atoi(hotStr)
		list.PushFront(paper)
	}
	return *list
}

func htmlText(node *html.Node) string {
	return strings.Trim(htmlquery.InnerText(node), " ")
}

func main() {
	papers := getPapers()
	for paper := papers.Front(); paper != nil; paper = paper.Next() {
		log.Println(paper)
	}
}
