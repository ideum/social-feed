package main

import "github.com/ChimeraCoder/anaconda"

var twitterApi *anaconda.TwitterApi

func init() {
	anaconda.SetConsumerKey(cfg.Twitter.ConsumerKey)
	anaconda.SetConsumerSecret(cfg.Twitter.ConsumerSecret)

	twitterApi = anaconda.NewTwitterApi(cfg.Twitter.AccessToken, cfg.Twitter.AccessTokenSecret)
}

func GetTwitterPosts() ([]SocialPost, error) {
	tweets, err := twitterApi.GetSearch("ideum", nil)

	if err != nil {
		return nil, err
	}

	res := make([]SocialPost, 0, len(tweets))

	for _, tweet := range tweets {
		res = append(res, SocialPost{
			Text: tweet.Text,
		})
	}

	return res, nil
}
