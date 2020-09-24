package main

import (
	"github.com/gocolly/colly"
	"github.com/paveg/ps5_crawler/api"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

const URL = "https://www.amazon.co.jp/dp/B08GGGBKRQ"
const WithEcoBagURL = "https://www.amazon.co.jp/dp/B08GGGCH3Y"

const OutOfStackMessage = "現在在庫切れです。"

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
	log.Infoln("execute main.")
	result := crawl()
	if result {
		err := api.Notify("[test]In stock"); if err != nil {
			log.Errorln(err)
		}
	}
}

func cleanText(text string) string {
	result := strings.Replace(text, "\n", "", -1)
	return strings.TrimSpace(result)
}

// FIXME: I'd like to return any URLs I have in stock.
func crawl() bool {
	v := false
	c := colly.NewCollector()

	spanSelector := "#availability > span:first-child"
	c.OnHTML(spanSelector, func(e *colly.HTMLElement) {
		text := e.DOM.Text()
		result := cleanText(text)

		if result != OutOfStackMessage {
			log.Infoln("ok")
			v = true
		}
	})

	err := c.Visit(URL)
	if err != nil {
		log.Warnln(err)
	}
	err = c.Visit(WithEcoBagURL)
	if err != nil {
		log.Warnln(err)
	}

	return v
}
