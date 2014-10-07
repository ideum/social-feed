package main

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestSocialPostJSONMarshalling(t *testing.T) {
	u, _ := url.Parse("http://www.google.com")
	p := SocialPost{Text: "test", Url: *u}

	j, err := p.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	var d interface{}
	json.Unmarshal(j, &d)
	m := d.(map[string]interface{})

	if m["text"].(string) != "test" {
		t.Fail()
	}
	if m["url"].(string) != "http://www.google.com" {
		t.Fail()
	}
	if m["image"].(string) != "" {
		t.Fail()
	}
}
