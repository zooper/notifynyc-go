package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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
	for {
		f, err := os.OpenFile("/log/log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println(err)
		}

		resp, err := http.Get("https://a858-nycnotify.nyc.gov/RSS/NotifyNYC?lang=en")
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		var rss RSS
		if err := xml.Unmarshal(body, &rss); err != nil {
			fmt.Println("Cant parse rss: ")
			log.Println(err)

		} else {
			for _, item := range rss.Channel.Items {

				file, err := os.ReadFile("/log/log.txt")
				if err != nil {
					fmt.Println(err)
				}
				if strings.Contains(string(file), item.PubDate) {
					break

				} else {
					// Write the new entries to file
					f.WriteString(item.PubDate + "\n")
					// Remove the translation junk in the end of the message
					clean_Description := strings.Split(item.Description, "To view")[0]
					telegram(item.PubDate, item.Title, clean_Description)
				}
			}
		}
		f.Close()
		// Sleep for 5 minutes
		time.Sleep(5 * time.Minute)
	}
}
