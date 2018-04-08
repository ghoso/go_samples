package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

func main(){
	// Get args
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: rss.get [rss_urls...]\n")
		os.Exit(1)
	}
		
	for _,url := range urls {
		// Get RSS
		rss_data,err := GetRSS(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "rss.get Error: %s \n",err)
			continue
		}
		
		// Output
		ShowRss(rss_data.Channel)
	}
}

func GetRSS(url string) (RssFeed,error) {
	res,err := http.Get(url)
	if err != nil {
		return RssFeed{},err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	rss_res,err := ParseRSS(body)
	if err != nil {
		return RssFeed{},err
	}
	return rss_res,nil
}

func ShowRss(rss RssChannel){
	fmt.Println("Channel Info")
	fmt.Println("Title:", rss.Title)
	fmt.Println("Link:", rss.Link)
	fmt.Println("Last build date:", rss.LastBuildDate)
	fmt.Println("Pub date:", rss.PubDate)
	fmt.Println("Language:", rss.Language)
	fmt.Println("Description:", rss.Description)

	for _,entry := range rss.Items {
		fmt.Println("Title:", entry.Title)
		fmt.Println("Link:", entry.Link)
		fmt.Println("Pub date:", entry.PubDate)
	}
}

type RssFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel RssChannel `xml:"channel"`
}

type RssChannel struct {
	XMLName xml.Name `xml:"channel"`
	Title  string  `xml:"title"`
	Description string `xml:"description"`
	Link string  `xml:"link"`
	Language string  `xml:"language"`
	PubDate  string  `xml:"pubDate"`
	LastBuildDate string  `xml:"lastBuildDate"`
	Items []RssItem `xml:"item"`
}

type RssItem struct {
	Title  string  `xml:"title"`
	Link string  `xml:"link"`
	PubDate  string  `xml:"pubDate"`
}

func ParseRSS(data []byte) (RssFeed,error){
	feed := RssFeed{}
	err := xml.Unmarshal(data, &feed)
	if err != nil {
		return RssFeed{},err
	}
	return feed,nil
}

