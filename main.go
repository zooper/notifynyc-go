package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func main() {
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.Get("https://a858-nycnotify.nyc.gov/RSS/NotifyNYC?lang=en")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var rss RSS
	if err := xml.Unmarshal(body, &rss); err != nil {
		panic(err)
	}
	for _, item := range rss.Channel.Items {

		file, err := os.ReadFile("log.txt")
		if err != nil {
			fmt.Println(err)
		}
		if strings.Contains(string(file), item.PubDate) {
			break

		} else {
			// Write the new entries to file
			f.WriteString(item.PubDate + "\n")
			matrix(item.PubDate, item.Description, item.Title)
		}
	}
	f.Close()
}
