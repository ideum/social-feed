package youtube

import (
	"code.google.com/p/goauth2/oauth/jwt"
	"code.google.com/p/google-api-go-client/youtube/v3"
	"fmt"
	"github.com/ideum/social-feed/social"
	"log"
	"net/url"
	"time"
)

type Api struct{ *youtube.Service }

type Credentials struct {
	Id  string
	Key string
}

func New(c *Credentials) *Api {
	t, err := jwt.NewTransport(jwt.NewToken(
		c.Id,
		youtube.YoutubeReadonlyScope,
		[]byte(c.Key),
	))
	if err != nil {
		log.Fatal(err)
	}

	api, err := youtube.New(t.Client())
	if err != nil {
		log.Fatal(err)
	}

	return &Api{api}
}

func (api *Api) GetPosts() ([]social.Post, error) {
	svc := youtube.NewPlaylistItemsService(api.Service)
	query := svc.
		List("id,snippet").
		PlaylistId("UUh1eDd1edZ8ZKjMfm6FOx2A").
		MaxResults(10)
	feed, err := query.Do()
	if err != nil {
		return nil, err
	}

	items := feed.Items

	for feed.NextPageToken != "" {
		feed, err = query.PageToken(feed.NextPageToken).Do()
		if err != nil {
			return nil, err
		}
	}

	posts := make([]social.Post, 0, len(items))
	for _, item := range items {
		c, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		u, _ := url.Parse(fmt.Sprintf(
			"https://www.youtube.com/watch?v=%s",
			item.Snippet.ResourceId.VideoId,
		))
		i, _ := url.Parse(item.Snippet.Thumbnails.Default.Url)

		posts = append(posts, social.Post{
			Source:    social.YouTube,
			CreatedAt: c,
			Text:      item.Snippet.Title,
			Url:       *u,
			Image:     *i,
		})
	}

	return posts, nil
}
