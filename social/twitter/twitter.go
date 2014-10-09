package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/ideum/social-feed/social"
	"net/url"
)

type Api struct{ *anaconda.TwitterApi }

type Credentials struct {
	ConsumerKey, ConsumerSecret    string
	AccessToken, AccessTokenSecret string
}

func New(c *Credentials) *Api {
	anaconda.SetConsumerKey(c.ConsumerKey)
	anaconda.SetConsumerSecret(c.ConsumerSecret)

	api := anaconda.NewTwitterApi(c.AccessToken, c.AccessTokenSecret)

	return &Api{api}
}

func (api *Api) GetPosts() ([]social.Post, error) {
	tweets, err := api.GetUserTimeline(nil)

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
