package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Text    string     `xml:",chardata"`
	Version string     `xml:"version,attr"`
	Channel RssChannel `xml:"channel"`
}

type RssChannel struct {
	Text        string    `xml:",chardata"`
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Language    string    `xml:"language"`
	Item        []RssItem `xml:"item"`
}
type RssItem struct {
	Text        string  `xml:",chardata"`
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	Description string  `xml:"description"`
	Guid        RssGuid `xml:"guid"`
	PubDate     string  `xml:"pubDate"`
}

type RssGuid struct {
	Text        string `xml:",chardata"`
	IsPermaLink string `xml:"isPermaLink,attr"`
}

func rssFeedToStruct(url string) (RssFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RssFeed{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RssFeed{}, nil
	}

	rssFeed := RssFeed{}

	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RssFeed{}, nil
	}

	return rssFeed, nil
}
