package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type discuss struct {
	PrimaryGroups []struct {
		FlairBgColor string `json:"flair_bg_color"`
		FlairColor   string `json:"flair_color"`
		FlairURL     string `json:"flair_url"`
		ID           int64  `json:"id"`
		Name         string `json:"name"`
	} `json:"primary_groups"`
	TopicList struct {
		CanCreateTopic bool        `json:"can_create_topic"`
		Draft          interface{} `json:"draft"`
		DraftKey       string      `json:"draft_key"`
		DraftSequence  interface{} `json:"draft_sequence"`
		MoreTopicsURL  string      `json:"more_topics_url"`
		PerPage        int64       `json:"per_page"`
		TopTags        []string    `json:"top_tags"`
		Topics         []struct {
			Archetype          string      `json:"archetype"`
			Archived           bool        `json:"archived"`
			Bookmarked         interface{} `json:"bookmarked"`
			Bumped             bool        `json:"bumped"`
			BumpedAt           string      `json:"bumped_at"`
			CanVote            bool        `json:"can_vote"`
			CategoryID         int64       `json:"category_id"`
			Closed             bool        `json:"closed"`
			CreatedAt          string      `json:"created_at"`
			Excerpt            string      `json:"excerpt"`
			FancyTitle         string      `json:"fancy_title"`
			FeaturedLink       interface{} `json:"featured_link"`
			HasAcceptedAnswer  bool        `json:"has_accepted_answer"`
			HasSummary         bool        `json:"has_summary"`
			HighestPostNumber  int64       `json:"highest_post_number"`
			ID                 int64       `json:"id"`
			ImageURL           string      `json:"image_url"`
			LastPostedAt       string      `json:"last_posted_at"`
			LastPosterUsername string      `json:"last_poster_username"`
			LikeCount          int64       `json:"like_count"`
			Liked              interface{} `json:"liked"`
			Pinned             bool        `json:"pinned"`
			PinnedGlobally     bool        `json:"pinned_globally"`
			Posters            []struct {
				Description    string `json:"description"`
				Extras         string `json:"extras"`
				PrimaryGroupID int64  `json:"primary_group_id"`
				UserID         int64  `json:"user_id"`
			} `json:"posters"`
			PostsCount int64       `json:"posts_count"`
			ReplyCount int64       `json:"reply_count"`
			Slug       string      `json:"slug"`
			Tags       []string    `json:"tags"`
			Title      string      `json:"title"`
			Unpinned   interface{} `json:"unpinned"`
			Unseen     bool        `json:"unseen"`
			UserVoted  interface{} `json:"user_voted"`
			Views      int64       `json:"views"`
			Visible    bool        `json:"visible"`
			VoteCount  interface{} `json:"vote_count"`
		} `json:"topics"`
	} `json:"topic_list"`
	Users []struct {
		AvatarTemplate string `json:"avatar_template"`
		ID             int64  `json:"id"`
		Name           string `json:"name"`
		Username       string `json:"username"`
	} `json:"users"`
}

func main() {
	fmt.Printf("%-24s %s\n", "Date", "Title")
	fmt.Printf("%-24s %s\n", "====", "=====")
	for i := 0; i < 10; i++ {
		unanswered(i)
	}
}

func unanswered(page int) {
	response, err := http.Get("https://discuss.elastic.co/c/elastic-stack/beats/l/latest.json?page=" + strconv.Itoa(page))
	checkErr(err)

	responseData, err := ioutil.ReadAll(response.Body)
	checkErr(err)

	var responseObject discuss
	err = json.Unmarshal(responseData, &responseObject)
	checkErr(err)

	for _, topic := range responseObject.TopicList.Topics {
		if topic.PostsCount == 1 {
			fmt.Printf("%24s %s\n", topic.CreatedAt, topic.Title)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
