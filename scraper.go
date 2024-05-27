package main

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/adamhu714/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) startScraping(concurrency int, timeBetweenRequests time.Duration) {
	log.Printf("Fetching feeds every %s on %v goroutines....", timeBetweenRequests, concurrency)
	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		feeds, err := apiCfg.DB.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error getting next feeds to fetch: %s\n", err.Error())
			continue
		}
		log.Printf("Found %v feeds to fetch.\n", len(feeds))

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go apiCfg.scrapeFeed(wg, feed)
		}
		wg.Wait()
	}
}

func (apiCfg *apiConfig) scrapeFeed(wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := apiCfg.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed fetched: %s", err.Error())
		return
	}

	feedData, err := rssFeedToStruct(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed, %s: %s", feed.Name, err.Error())
		return
	}

	numCreated := 0

	for _, post := range feedData.Channel.Item {
		_, err := apiCfg.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       post.Title,
			Url:         post.Link,
			Description: post.Description,
			PublishedAt: time.Now().UTC(),
			FeedID:      feed.ID,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				log.Printf("error creating post: %s", err.Error())
			}
		} else {
			numCreated++
		}
	}

	log.Printf("Feed %s fetched, %d posts found and %d new posts added to database.\n", feed.Name, len(feedData.Channel.Item), numCreated)
}
