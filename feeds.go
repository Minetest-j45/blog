package main

import (
	"os"
	"strings"
	"time"

	"github.com/gorilla/feeds"
)

type Blog struct {
	Title string
	Date  string
	Text  string
}

func main() {
	feed := &feeds.Feed{
		Title:       "j45's Blog",
		Link:        &feeds.Link{Href: "https://minetest-j45.github.io/blog/"},
		Description: "j45's Blog",
		Author:      &feeds.Author{Name: "j45"},
		Created:     time.Now(),
	}

	var blogs []Blog

	var feedItems []*feeds.Item

	dat, err := os.ReadFile("blogs.csv")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")

	lines = lines[1:] // first line is the layout so remove it
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		fields := strings.Split(line, "|")
		if len(fields) < 3 { // invalid line
			continue
		}

		blog := Blog{
			Title: fields[0],
			Text:  fields[2],
		}
		blogs = append(blogs, blog)
	}

	for _, blog := range blogs {
		feedItems = append(feedItems,
			&feeds.Item{
				Title:       blog.Title,
				Link:        &feeds.Link{Href: "https://minetest-j45.github.io/blog/post_" + strings.ToLower(strings.ReplaceAll(blog.Title, " ", "_")) + ".html"},
				Description: blog.Text,
			})
	}
	feed.Items = feedItems

	rssFeed, err := feed.ToRss()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("feed.rss", []byte(rssFeed), 0644)
	if err != nil {
		panic(err)
	}

	jsonFeed, err := feed.ToJSON()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("feed.json", []byte(jsonFeed), 0644)
	if err != nil {
		panic(err)
	}
}
