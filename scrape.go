package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"encoding/json"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
    subreddit := flag.String("subreddit", "TurnipExchange/new",
		"The name of a subreddit and category if you'd like.")
	lmt := flag.Int("lmt", 0,
		"How many posts off the first page you want returned.")
	filterFlair := flag.String("filter", "",
		"A flair you'd like to exclude from the results.")
	flag.Parse()

	posts := scrapeReddit(*subreddit, *lmt, *filterFlair)

	prettyPrintPosts(posts)

	savePostsToJson(posts)
}

type post struct {
	Title string `json:"title"`
    PostURL string `json:"postUrl"`
    Time string `json:"time"`
    Author string `json:"author"`
    Comments int `json:"comments"`
    Source string `json:"source"`
    Flair string `json:"flair"`
}

func scrapeReddit(subreddit string, limit int, filterFlair string) []post {
	var posts []post
	c := colly.NewCollector()

	c.OnHTML("div.top-matter", func(e *colly.HTMLElement) {
		flair := e.ChildAttr("span.linkflairlabel", "title")

		if filterFlair != "" && flair != filterFlair {
			posts = append(posts, post{
				Title: e.ChildText("a.title"),
				PostURL: oldUrlToNew(e.ChildAttr("a.comments", "href")),
				Flair: flair,
				Time: e.ChildAttr("time.live-timestamp", "title"),
				Author: e.ChildText("a.author"),
				Comments: getCommentNum(e.ChildText("a.comments")),
				Source: getLinkedURL(e.ChildAttr("a.title", "href")),
			})
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	subreddit = strings.Trim(subreddit, "/")
	c.Visit(fmt.Sprintf("https://old.reddit.com/r/%s/", subreddit))

	if len(posts) > limit && limit > 0 {
		return posts[:limit]
	}
	fmt.Println(len(posts))
	return posts
}

func getCommentNum(commentText string) int {
	// No comments
	if commentText == "comment" {
		return 0
	}
	// only 1 comment
	if commentText == "1 comment" {
        return 1
    }
	// commentsText will be "(some int) comments" so return the int
	CommentNum, err := strconv.Atoi(commentText[:len(commentText)-9])
    checkErr(err)
	return CommentNum
}

func oldUrlToNew(url string) string {
	// replaces 'old' with 'www'
	// 'https://old.reddit.com/...' --> 'https://www.reddit.com/...'
	return fmt.Sprintf("%swww%s", url[:8], url[11:])
}

func getLinkedURL(url string) string {
	// if this is a text post, return empty string
	if url[:3] == "/r/" {
		return ""
	}
	// if there is a link return the link
	return url
}

func prettyPrintPosts(posts []post) {
	for _, post := range posts {
		postText := fmt.Sprintf(
			"%s\n\tPost URL: %s\n\tTime: %s\n\tAuthor: /u/%s\n\tComments: %d",
			post.Title, post.PostURL, post.Time, post.Author, post.Comments)
		if post.Source != "" {
			postText += "\n\tSource: " + post.Source
		}
		if post.Flair != "" {
			postText += "\n\tFlair: " + post.Flair
		}
		fmt.Println(postText + "\n")
	}
}

func savePostsToJson(posts []post) {
	postsJson, err := json.MarshalIndent(posts, "", "	")
	checkErr(err)
	ioutil.WriteFile("output.json", postsJson, 0644)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
