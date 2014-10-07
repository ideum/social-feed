package main

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

var twitterApi *anaconda.TwitterApi

func init() {
	anaconda.SetConsumerKey(cfg.Twitter.ConsumerKey)
	anaconda.SetConsumerSecret(cfg.Twitter.ConsumerSecret)

	twitterApi = anaconda.NewTwitterApi(cfg.Twitter.AccessToken, cfg.Twitter.AccessTokenSecret)
}

func GetTwitterPosts() ([]SocialPost, error) {
	tweets, err := twitterApi.GetUserTimeline(nil)

	if err != nil {
		return nil, err
	}

	res := make([]SocialPost, 0, len(tweets))

	for _, tweet := range tweets {
		t, _ := tweet.CreatedAtTime()
		u, _ := url.Parse("https://twitter.com/ideum/status/" + tweet.IdStr)

		res = append(res, SocialPost{
			Source:    TwitterSource,
			CreatedAt: t,
			Text:      tweet.Text,
			Url:       *u,
		})
	}

	return res, nil
}
