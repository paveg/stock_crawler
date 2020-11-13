package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/gocolly/colly"
	"github.com/paveg/ps5_crawler/api"
	log "github.com/sirupsen/logrus"
)

const URL = "https://www.amazon.co.jp/dp/B08GGGBKRQ"
const WithEcoBagURL = "https://www.amazon.co.jp/dp/B08GGGCH3Y"

const OutOfStackMessage = "現在在庫切れです。"
const ExhibitionSellersMessage = "出品者からお求めいただけます。"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	logLevel := log.InfoLevel
	if os.Getenv("DEBUG_MODE") != "" {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)
	log.Debugln("PS5 crawler has initialized.")
}

func main() {
	log.Infoln("executing...")
	result := false
	urls := []string{URL, WithEcoBagURL}

	for _, url := range urls {
		result = crawl(url)
		if result {
			err := api.Notify(fmt.Sprintf("In stock\n%s", url))
			if err != nil {
				log.Errorln(err)
			}
		}
	}

	log.Infoln("PS5 crawler has finished")
}

func cleanText(text string) string {
	result := strings.Replace(text, "\n", "", -1)
	return strings.TrimSpace(result)
}

func crawl(url string) bool {
	stockResult := false
	c := colly.NewCollector()

	spanSelector := "#availability > span:first-child"
	c.OnHTML(spanSelector, func(e *colly.HTMLElement) {
		text := e.DOM.Text()
		result := cleanText(text)

		if result != OutOfStackMessage && result != ExhibitionSellersMessage {
			log.Infoln("ok")
			stockResult = true
		}
	})

	err := c.Visit(url)
	if err != nil {
		log.Errorln(err)
	}

	return stockResult
}
