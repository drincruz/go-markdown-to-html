package main

import (
	"log"
	"testing"
)

func TestGetMostRecentArray(t *testing.T) {
	log.Printf("[TestGetMostRecentArray]\n")

	posts := []Blog{
		{Title: "Engineering Guilds", Url: "/2023/10/21/engineering-guilds-at-work.html"},
		{Title: "Learn Go", Url: "/2023/10/20/why-you-should-learn-go.html"},
		{Title: "Learn Python", Url: "/2023/10/19/why-you-should-learn-python.html"},
	}
	want := []string{
		"markdown/2023/10/21/engineering-guilds-at-work.markdown",
		"markdown/2023/10/20/why-you-should-learn-go.markdown",
		"markdown/2023/10/19/why-you-should-learn-python.markdown",
	}
	got := getMostRecentArray(posts)

	for index, val := range want {
		if got[index] != val {
			t.Errorf("getMostRecentArray: %s != %s", got[index], val)
		}
	}
}
