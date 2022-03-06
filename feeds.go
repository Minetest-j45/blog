package main

import (
	"io/ioutil"
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
		Link:        &feeds.Link{Href: "https://j1233.minetest.land/blog/"},
		Description: "j45's Blog",
		Author:      &feeds.Author{Name: "j45"},
		Created:     time.Now(),
	}

	var blogs []Blog

	var feedItems []*feeds.Item

	dat, err := ioutil.ReadFile("blogs.csv")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")

	lines = lines[1 : len(lines)-1]
	//for _, line := range lines {
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		fields := strings.Split(line, "|")
		blog := Blog{
			Title: fields[0],
			Text:  fields[2],
		}
		blogs = append(blogs, blog)
	}

	for _, blog := range blogs {
		titleslice := strings.Split(blog.Title, " ")
		link := "https://j1233.minetest.land/blog/#" + strings.Join(titleslice, "_")
		feedItems = append(feedItems,
			&feeds.Item{
				Title:       blog.Title,
				Link:        &feeds.Link{Href: link},
				Description: blog.Text,
			})
	}
	feed.Items = feedItems

	rssFeed, err := feed.ToRss()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("feed.rss", []byte(rssFeed), 0644)
	if err != nil {
		panic(err)
	}

	jsonFeed, err := feed.ToJSON()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("feed.json", []byte(jsonFeed), 0644)
	if err != nil {
		panic(err)
	}
}
