package redditproto

import (
	"testing"
)

func TestParseComment(t *testing.T) {
	replyTree := `{
		"kind": "Listing",
		"data": {
			"children": [
				{
					"kind": "t1",
					"data": {
						"body": "reply1",
						"replies": ""
					}
				},
				{
					"kind": "t1",
					"data": {
						"body": "reply2",
						"replies": ""
					}
				}
			]
		}
	}`
	comment, err := ParseComment([]byte(
		`{
			"kind": "t1",
			"data": {
				"body": "something",
				"replies": ` + replyTree + `
			}
		}`),
	)
	if err != nil {
		t.Fatal(err)
	}

	if comment == nil {
		t.Fatal("returned comment was nil")
	}

	if len(comment.Replies) != 2 {
		t.Fatalf("got %d replies; wanted 2", len(comment.Replies))
	}

	if comment.GetBody() != "something" {
		t.Errorf("got %s; wanted something", comment.GetBody())
	}

	if comment.Replies[0].GetBody() != "reply1" {
		t.Errorf("got %s; wanted reply1", comment.Replies[0].GetBody())
	}

	if comment.Replies[1].GetBody() != "reply2" {
		t.Errorf("got %s; wanted reply2", comment.Replies[1].GetBody())
	}
}

func TestParseThread(t *testing.T) {
	link, err := ParseThread([]byte(
		`[
			{
				"kind": "Listing",
				"data": {
					"children": [
						{
							"kind": "t3",
							"data": {
								"title": "1"
							}
						}
					]
				}
			},
			{
				"kind": "Listing",
				"data": {
					"children": [
						{
							"kind": "t1",
							"data": {
								"body": "1"
							}
						},
						{
							"kind": "t1",
							"data": {
								"body": "2"
							}
						}
					]
				}
			}
		]`))
	if err != nil {
		t.Fatal(err)
	}

	if link == nil {
		t.Fatalf("returned link was nil")
	}

	if len(link.Comments) != 2 {
		t.Errorf("got %d comments; wanted 2", len(link.Comments))
	}
}

func TestParseListing(t *testing.T) {
	links, comments, messages, err := ParseListing([]byte(
		`{
			"kind": "Listing",
			"data": {
				"children": [
					{
						"kind": "t3",
						"data": {
							"title": "1"
						}
					},
					{
						"kind": "t1",
						"data": {
							"body": "1"
						}
					},
					{
						"kind": "t4",
						"data": {
							"body": "1"
						}
					}
				]
			}
		}`))
	if err != nil {
		t.Fatal(err)
	}

	if len(links) != 1 {
		t.Errorf("got %d links; wanted 1", len(links))
	}

	if len(comments) != 1 {
		t.Errorf("got %d comments; wanted 1", len(comments))
	}
	if len(messages) != 1 {
		t.Errorf("got %d messages; wanted 1", len(messages))
	}
}
