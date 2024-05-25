package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/adamhu714/rssagg/internal/database"
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

	for _, post := range feedData.Channel.Item {
		log.Printf("Found post: %s\n", post.Title)
	}

	log.Printf("Feed %s fetched, %d posts found.\n", feed.Name, len(feedData.Channel.Item))
}
