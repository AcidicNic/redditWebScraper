# 🕷 Reddit WebScraper

[![Go Report Card](https://goreportcard.com/badge/github.com/acidicnic/redditWebScraper)](https://goreportcard.com/report/github.com/acidicnic/redditWebScraper)

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started)
3. [Flags](#flags)
4. [Output](#output.json)

## Project Structure

```bash
📂 redditWebScraper
├── README.md
└── scrape.go
```

## Getting Started

```bash
git clone https://github.com/AcidicNic/redditWebScraper.git
cd redditWebScraper
go run scrape.go
```

## Flags


```bash
// This will get the 10 newest posts from /r/TurnipExchange, excluding posts with the flair CLOSED
go run scrape.go -subreddit="TurnipExchange/new" -lmt=10 -filter="CLOSED"
```

- __subreddit__ _string_
    - The name of a subreddit and category if you'd like. (_default "TurnipExchange/new"_)
- __filter__ _string_
	- A flair you'd like to exclude from the results.
- __lmt__ _int_
	- How many posts off the first page you want returned.


## output.json
```json
[
	{
        /* Post title */
		"title": "This is the title of a super interesting post!",

        /* Direct link to the post */
		"postUrl": "https://www.reddit.com/r/<subreddit>/comments/<post ID>/<post_slug>/",

        /* Time the post was uploaded to Reddit */
		"time": "Thu Jul 16 00:45:14 2020 UTC",

        /* Author's username */
		"author": "RedditUser31498234",

        /* Number of comments */
		"comments": 0,

        /* Linked URL (Empty String for text posts) */
		"source": "",

        /* Flair (Empty String for unflaired posts) */
		"flair": ""
	},
	{
        ...
    },
    ...
]
```
