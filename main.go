// Command unanswered lists topics on discuss.elastic.co that have no replies.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const baseURL = "https://discuss.elastic.co"

type topic struct {
	CreatedAt         string `json:"created_at"`
	Title             string `json:"title"`
	PostsCount        int64  `json:"posts_count"`
	Closed            bool   `json:"closed"`
	Pinned            bool   `json:"pinned"`
	HasAcceptedAnswer bool   `json:"has_accepted_answer"`
}

type latest struct {
	TopicList struct {
		MoreTopicsURL string  `json:"more_topics_url"`
		Topics        []topic `json:"topics"`
	} `json:"topic_list"`
}

func main() {
	category := flag.String("category", "elastic-stack/beats/28", "category path, as in /c/<category>/l/latest")
	pages := flag.Int("pages", 10, "maximum number of pages to fetch")
	flag.Parse()

	client := &http.Client{Timeout: 10 * time.Second}

	fmt.Printf("%-24s %s\n", "Date", "Title")
	fmt.Printf("%-24s %s\n", "====", "=====")
	for page := 0; page < *pages; page++ {
		l, err := fetch(client, *category, page)
		if err != nil {
			log.Fatalf("page %d: %v", page, err)
		}
		for _, t := range unanswered(l.TopicList.Topics) {
			fmt.Printf("%-24s %s\n", t.CreatedAt, t.Title)
		}
		if l.TopicList.MoreTopicsURL == "" || len(l.TopicList.Topics) == 0 {
			break
		}
	}
}

func fetch(client *http.Client, category string, page int) (*latest, error) {
	u := fmt.Sprintf("%s/c/%s/l/latest.json?page=%d", baseURL, category, page)
	resp, err := client.Get(u)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %s", u, resp.Status)
	}
	return decode(resp.Body)
}

func decode(r io.Reader) (*latest, error) {
	var l latest
	if err := json.NewDecoder(r).Decode(&l); err != nil {
		return nil, err
	}
	return &l, nil
}

// unanswered returns the open, unpinned topics that have only the original post.
func unanswered(topics []topic) []topic {
	var out []topic
	for _, t := range topics {
		if t.PostsCount == 1 && !t.Closed && !t.Pinned && !t.HasAcceptedAnswer {
			out = append(out, t)
		}
	}
	return out
}
