package main

import (
	"fmt"
	"os"
	"strings"
)

type Blog struct {
	Title string
	Date  string
	Text  string
}

func makeTitles(posts []Blog) (titles []string) {
	for _, post := range posts {
		titles = append(titles, "<div class='blog'><br><a href='https://minetest-j45.github.io/blog/post_"+strings.ToLower(strings.ReplaceAll(post.Title, " ", "_"))+".html' style='text-decoration: none; color: #000;'><mybutton>"+post.Title+"</mybutton></a></div>")
	}
	return titles
}

func makePost(post Blog) (content string) {
	return "<bigtext>" + post.Title + "</bigtext> - " + post.Date + "<a href='https://minetest-j45.github.io/blog/'><exit style='color: #000;'>X</exit></a><br><p>" + post.Text + "</p><br><myButton onclick='function copy() { navigator.clipboard.writeText(`https://minetest-j45.github.io/blog/post_" + strings.ToLower(strings.ReplaceAll(post.Title, " ", "_")) + ".html`); alert(`Copied blog post link to clipboard`); }; copy();'>Share this blog post!</myButton><br><a href='https://minetest-j45.github.io/blog/'><img src='https://minetest-j45.github.io/images/return.png' style='cursor: pointer;'></a>"
}

func main() {
	files, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "post_") && strings.HasSuffix(file.Name(), ".html") {
			os.Remove(file.Name())
		}
	}

	dat, err := os.ReadFile("blogs.csv")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")

	var posts []Blog
	lines = lines[1:] // first line is the layout so remove it
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		fields := strings.Split(line, "|")
		if len(fields) < 3 { // invalid line
			continue
		}

		post := Blog{
			Title: fields[0],
			Date:  fields[1],
			Text:  fields[2],
		}
		posts = append(posts, post)

		template_post, err := os.ReadFile("template_post.html")
		if err != nil {
			panic(err)
		}

		var wanted string = "{{ post }}"
		before, after, found := strings.Cut(string(template_post), wanted)
		if found == false {
			fmt.Println(wanted + " not found in provided template")
		}

		err = os.WriteFile("post_"+strings.ToLower(strings.ReplaceAll(post.Title, " ", "_"))+".html", []byte(before+makePost(post)+after), 0644)
		if err != nil {
			panic(err)
		}
	}

	template_index, err := os.ReadFile("template_index.html")
	if err != nil {
		panic(err)
	}

	var wanted string = "{{ posts }}"
	before, after, found := strings.Cut(string(template_index), wanted)
	if found == false {
		fmt.Println(wanted + " not found in provided template")
	}

	err = os.WriteFile("index.html", []byte(before+strings.Join(makeTitles(posts), "\n")+after), 0644)
	if err != nil {
		panic(err)
	}
}
