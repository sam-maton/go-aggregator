package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{}
	feed := RSSFeed{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		return &feed, fmt.Errorf("there was an error creating the new request: %w", err)
	}

	resp, err := client.Do(req)

	if err != nil {
		return &feed, fmt.Errorf("there was an error whilst making a request to the RSS Feed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return &feed, fmt.Errorf("there was an error reading the response body: %w", err)
	}

	xml.Unmarshal(body, &feed)

	feed.unescape()

	return &feed, nil
}

func (feed *RSSFeed) unescape() {
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)

	for k, v := range feed.Channel.Item {
		feed.Channel.Item[k].Description = html.UnescapeString(v.Description)
		feed.Channel.Item[k].Title = html.UnescapeString(v.Title)
	}
}
