package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type discuss struct {
	Users []struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Name           string `json:"name"`
		AvatarTemplate string `json:"avatar_template"`
	} `json:"users"`
	PrimaryGroups []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		FlairURL     string `json:"flair_url"`
		FlairBgColor string `json:"flair_bg_color"`
		FlairColor   string `json:"flair_color"`
	} `json:"primary_groups"`
	TopicList struct {
		CanCreateTopic bool        `json:"can_create_topic"`
		MoreTopicsURL  string      `json:"more_topics_url"`
		Draft          interface{} `json:"draft"`
		DraftKey       string      `json:"draft_key"`
		DraftSequence  int         `json:"draft_sequence"`
		PerPage        int         `json:"per_page"`
		TopTags        []string    `json:"top_tags"`
		Topics         []struct {
			ID                 int           `json:"id"`
			Title              string        `json:"title"`
			FancyTitle         string        `json:"fancy_title"`
			Slug               string        `json:"slug"`
			PostsCount         int           `json:"posts_count"`
			ReplyCount         int           `json:"reply_count"`
			HighestPostNumber  int           `json:"highest_post_number"`
			ImageURL           interface{}   `json:"image_url"`
			CreatedAt          time.Time     `json:"created_at"`
			LastPostedAt       time.Time     `json:"last_posted_at"`
			Bumped             bool          `json:"bumped"`
			BumpedAt           time.Time     `json:"bumped_at"`
			Unseen             bool          `json:"unseen"`
			Pinned             bool          `json:"pinned"`
			Unpinned           interface{}   `json:"unpinned"`
			Excerpt            string        `json:"excerpt,omitempty"`
			Visible            bool          `json:"visible"`
			Closed             bool          `json:"closed"`
			Archived           bool          `json:"archived"`
			Bookmarked         interface{}   `json:"bookmarked"`
			Liked              interface{}   `json:"liked"`
			Tags               []interface{} `json:"tags"`
			Views              int           `json:"views"`
			LikeCount          int           `json:"like_count"`
			HasSummary         bool          `json:"has_summary"`
			Archetype          string        `json:"archetype"`
			LastPosterUsername string        `json:"last_poster_username"`
			CategoryID         int           `json:"category_id"`
			PinnedGlobally     bool          `json:"pinned_globally"`
			FeaturedLink       interface{}   `json:"featured_link"`
			HasAcceptedAnswer  bool          `json:"has_accepted_answer"`
			VoteCount          interface{}   `json:"vote_count"`
			CanVote            bool          `json:"can_vote"`
			UserVoted          bool          `json:"user_voted"`
			Posters            []struct {
				Extras         interface{} `json:"extras"`
				Description    string      `json:"description"`
				UserID         int         `json:"user_id"`
				PrimaryGroupID int         `json:"primary_group_id"`
			} `json:"posters"`
			LastReadPostNumber int `json:"last_read_post_number,omitempty"`
			Unread             int `json:"unread,omitempty"`
			NewPosts           int `json:"new_posts,omitempty"`
			NotificationLevel  int `json:"notification_level,omitempty"`
		} `json:"topics"`
	} `json:"topic_list"`
}

func main() {
	fmt.Printf("%-33s %s\n", "Date", "Title")
	fmt.Printf("%-33s %s\n", "====", "=====")
	for i := 0; i < 10; i++ {
		unanswered(i)
	}
}

func unanswered(page int) {
	response, err := http.Get("https://discuss.elastic.co/c/beats/l/latest.json?page=" + strconv.Itoa(page))
	checkErr(err)

	responseData, err := ioutil.ReadAll(response.Body)
	checkErr(err)

	var responseObject discuss
	err = json.Unmarshal(responseData, &responseObject)
	checkErr(err)

	for _, topic := range responseObject.TopicList.Topics {
		if topic.PostsCount == 1 {
			fmt.Printf("%33s %s\n", topic.CreatedAt, topic.Title)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
