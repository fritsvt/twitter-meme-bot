package redditproto

import (
	"encoding/json"
)

// redditThing represents the Thing class Reddit encloses all of its responses
// in.
type redditThing struct {
	// Kind defines the type of the Data field.
	Kind string
	// Data can contain basically anything. It's exciting that way.
	Data json.RawMessage
}

// redditListing represents one of the many lovely layers of misdirection Reddit
// uses to prevent developers from inferring the types of their JSON messages.
type redditListing struct {
	// Children is a list of surprises. There's no knowing what type each
	// redditThing is! Be sure to check their Kind fields.
	Children []*redditThing
}

// listingBuffer holds slices of the children in a listing.
type listingBuffer struct {
	comments []*Comment
	links    []*Link
	messages []*Message
}

// commentResponse represents the JSON message Reddit returns to represent
// comments.
type commentResponse struct {
	ApprovedBy          *string         `json:"approved_by,omitempty"`
	Author              *string         `json:"author,omitempty"`
	AuthorFlairCssClass *string         `json:"author_flair_css_class,omitempty"`
	AuthorFlairText     *string         `json:"author_flair_text,omitempty"`
	BannedBy            *string         `json:"banned_by,omitempty"`
	Body                *string         `json:"body,omitempty"`
	BodyHtml            *string         `json:"body_html,omitempty"`
	Gilded              *int32          `json:"gilded,omitempty"`
	LinkAuthor          *string         `json:"link_author,omitempty"`
	LinkUrl             *string         `json:"link_url,omitempty"`
	NumReports          *int32          `json:"num_reports,omitempty"`
	ParentId            *string         `json:"parent_id,omitempty"`
	Replies             json.RawMessage `json:"replies,omitempty"`
	Subreddit           *string         `json:"subreddit,omitempty"`
	SubredditId         *string         `json:"subreddit_id,omitempty"`
	Distinguished       *string         `json:"distinguished,omitempty"`
	Created             *float64        `json:"created,omitempty"`
	CreatedUtc          *float64        `json:"created_utc,omitempty"`
	Ups                 *int32          `json:"ups,omitempty"`
	Downs               *int32          `json:"downs,omitempty"`
	Likes               *bool           `json:"likes,omitempty"`
	Id                  *string         `json:"id,omitempty"`
	Name                *string         `json:"name,omitempty"`
}
