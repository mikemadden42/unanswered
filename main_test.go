package main

import (
	"os"
	"testing"
)

func TestDecodeAndFilter(t *testing.T) {
	f, err := os.Open("testdata/latest.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = f.Close() }()

	l, err := decode(f)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	if l.TopicList.MoreTopicsURL == "" {
		t.Error("more_topics_url not decoded")
	}
	if got := len(l.TopicList.Topics); got != 4 {
		t.Fatalf("got %d topics, want 4", got)
	}

	got := unanswered(l.TopicList.Topics)
	if len(got) != 1 || got[0].Title != "No replies yet" {
		t.Errorf("unanswered = %+v, want only %q", got, "No replies yet")
	}
}
