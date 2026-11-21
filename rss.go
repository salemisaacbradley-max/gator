package main

import ("fmt"
		"context"
		"net/http"
		"io"
		"html")

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
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error with fetching feed: %w", err)
	}
	request.Header.Set("User-Agent", "gator")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := http.Client.Do(requestPntr)
	if err != nil {
		return nil, fmt.Errorf("Error with Doing: %w", err)
	}
	defer resp.Body.Close()
	
	bite, err := io.readAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error Reading: %w", err)
	}
	newFeed := RSSFeed{}
	err = xml.Unmarshal(bite, &newFeed)
	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling: %w", err)
	}
	loopFeed := RSSFeed{}
	for _, items := range newFeed.Channel {
		chTitle := html.UnescapeString(items.Title)
		chDes := html.UnescapeString(items.Des)
		chItem := RSSItem{}
		for _, item := range items {
			title := html.UnescapeString(items.Title)
			des := html.UnescapeString(items.Description)
			chItem = {title, newFeed.Channel.Link, des, newFeed.Channel.Pubdate}
		}
	}
	return loopFeed, nil
}

