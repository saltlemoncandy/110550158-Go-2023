package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func main() {
	maxComments := flag.Int("max", 10, "Max number of comments to show")
	flag.Parse()

	c := colly.NewCollector()

	commentCount := 0 

	c.OnHTML(".push", func(e *colly.HTMLElement) {
		if commentCount < *maxComments { 
			name := e.ChildText(".push-userid") 
			content := strings.TrimSpace(e.ChildText(".push-content")) 
			time := e.ChildText(".push-ipdatetime") 
			if name != "" && content != "" && time != "" {
				commentCount++
				fmt.Printf("%d. 名字：%s，留言%s，時間： %s\n", commentCount, name, content, time)
			}
		}
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	url := "https://www.ptt.cc/bbs/joke/M.1481217639.A.4DF.html"
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
