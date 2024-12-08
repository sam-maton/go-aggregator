package rss

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sam-maton/go-aggregator/internal/database"
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

func ScrapeFeeds(ctx context.Context, db *database.Queries) error {
	feed, err := db.GetNextFeedToFetch(ctx)

	if err != nil {
		return fmt.Errorf("there was an error fetching the next feed: %w", err)
	}

	markArgs := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: feed.ID,
	}
	err = db.MarkFeedFetched(ctx, markArgs)

	if err != nil {
		return fmt.Errorf("there was an error marking the feed as fetched: %w", err)
	}

	rss, err := FetchFeed(ctx, feed.Url)

	if err != nil {
		return err
	}

	// layout := "01/02 03:04:05PM '06 -0700"

	for _, r := range rss.Channel.Item {

		// publishedAt, err := time.Parse(layout, r.PubDate)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return err
		// }

		postArgs := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       r.Title,
			Url:         r.Link,
			Description: r.Description,
			PublishedAt: time.Now(),
			FeedID:      feed.ID,
		}

		_, err = db.CreatePost(ctx, postArgs)

		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
