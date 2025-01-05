package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", req.UserAgent())
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var rsfeed *RSSFeed
	if err = xml.Unmarshal(data, &rsfeed); err != nil {
		return nil, err
	}

	rsfeed.Channel.Title = html.UnescapeString(rsfeed.Channel.Title)
	rsfeed.Channel.Description = html.UnescapeString(rsfeed.Channel.Description)
	// rsfeed.Channel.Item = html.UnescapeString(rsfeed.Channel.Description)
	for _, item := range rsfeed.Channel.Item {
		item.Description = html.UnescapeString(item.Description)
		item.Title = html.UnescapeString(item.Title)
	}
	return rsfeed, nil

}
