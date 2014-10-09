package flickr

import (
	"encoding/xml"
	"fmt"
	"github.com/ideum/social-feed/social"
	"github.com/mncaudill/go-flickr"
	"net/url"
	"time"
)

type Api struct {
	key, secret string
}

type Credentials struct {
	Key, Secret string
}

func New(c *Credentials) *Api {
	return &Api{key: c.Key, secret: c.Secret}
}

func (api *Api) GetPosts() ([]social.Post, error) {
	req := flickr.Request{
		ApiKey: api.key,
		Method: "flickr.people.getPublicPhotos",
		Args: map[string]string{
			"user_id":        "45874208@N00", // Ideum photos
			"privacy_filter": "1",            // public only
			"extras":         "date_upload",
		},
	}

	req.Sign(api.secret)

	res, err := req.Execute()
	if err != nil {
		return nil, err
	}

	var data struct {
		Photos []struct {
			Id         int    `xml:"id,attr"`
			Owner      string `xml:"owner,attr"`
			Secret     string `xml:"secret,attr"`
			Server     int    `xml:"server,attr"`
			Farm       int    `xml:"farm,attr"`
			Title      string `xml:"title,attr"`
			DateUpload int64  `xml:"dateupload,attr"`
		} `xml:"photos>photo"`
	}

	err = xml.Unmarshal([]byte(res), &data)
	if err != nil {
		return nil, err
	}

	posts := make([]social.Post, 0, len(data.Photos))
	for _, photo := range data.Photos {
		t := time.Unix(photo.DateUpload, 0)
		u, _ := url.Parse(fmt.Sprintf(
			"https://www.flickr.com/photos/%s/%d",
			photo.Owner,
			photo.Id,
		))
		i, _ := url.Parse(fmt.Sprintf(
			"https://farm%d.staticflickr.com/%d/%d_%s.jpg",
			photo.Farm,
			photo.Server,
			photo.Id,
			photo.Secret,
		))

		posts = append(posts, social.Post{
			Source:    social.Flickr,
			CreatedAt: t,
			Text:      photo.Title,
			Url:       *u,
			Image:     *i,
		})
	}

	return posts, nil
}
