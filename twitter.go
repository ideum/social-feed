package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/ideum/social-feed/social"
	"net/url"
)

var twitterApi *anaconda.TwitterApi

func init() {
	anaconda.SetConsumerKey(cfg.Twitter.ConsumerKey)
	anaconda.SetConsumerSecret(cfg.Twitter.ConsumerSecret)

	twitterApi = anaconda.NewTwitterApi(cfg.Twitter.AccessToken, cfg.Twitter.AccessTokenSecret)
}

func GetTwitterPosts() ([]social.Post, error) {
	tweets, err := twitterApi.GetUserTimeline(nil)

	if err != nil {
		return nil, err
	}

	res := make([]social.Post, 0, len(tweets))

	for _, tweet := range tweets {
		t, _ := tweet.CreatedAtTime()
		u, _ := url.Parse("https://twitter.com/ideum/status/" + tweet.IdStr)

		res = append(res, social.Post{
			Source:    social.Twitter,
			CreatedAt: t,
			Text:      tweet.Text,
			Url:       *u,
		})
	}

	return res, nil
}
