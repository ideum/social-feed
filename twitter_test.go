package main

import (
	"testing"
)

func TestFetchingTwitterFeed(t *testing.T) {
	if _, err := GetTwitterPosts(); err != nil {
		t.Fatal(err)
	}
}
