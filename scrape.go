package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
)

func main() {
	var posts []post
	c := colly.NewCollector()

	c.OnHTML("a.title", func(e *colly.HTMLElement) {
		link := getURL(e.Attr("href"))

		posts = append(posts, post{Title: e.Text, URL: link})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://old.reddit.com/r/TurnipExchange/new/")

	text := ""
	for _, post := range posts {
		text += fmt.Sprintf("%s \n\t%s\n\n", post.Title, post.URL)
		fmt.Printf("%s \n\t%s\n\n", post.Title, post.URL)
	}
	byte_text := []byte(text)
    err := ioutil.WriteFile("turnip_prices", byte_text, 0644)
	if err != nil {
        panic(err)
    }
}
func getURL(url string) string {
	if url[0:3] == "/r/" {
		return "https://www.reddit.com" + url
	}
	return url
}

type post struct {
	Title string
    URL string
}
