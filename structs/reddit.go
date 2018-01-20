package structs

type Thread = struct {
	ImageUrl string `json:"image_url"`
	Title string `json:"title"`
	Id string `json:"id"`
	Author string `json:"author"`
}

type RedditResponse struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash         string `json:"modhash"`
		WhitelistStatus string `json:"whitelist_status"`
		Threads
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

type Threads struct {
	Children        []struct {
		Kind string `json:"kind"`
		Data struct {
			Domain              string        `json:"domain"`
			ApprovedAtUtc       interface{}   `json:"approved_at_utc"`
			ModReasonBy         interface{}   `json:"mod_reason_by"`
			BannedBy            interface{}   `json:"banned_by"`
			NumReports          interface{}   `json:"num_reports"`
			SubredditID         string        `json:"subreddit_id"`
			ThumbnailWidth      int           `json:"thumbnail_width"`
			Subreddit           string        `json:"subreddit"`
			SelftextHTML        interface{}   `json:"selftext_html"`
			Selftext            string        `json:"selftext"`
			Likes               interface{}   `json:"likes"`
			SuggestedSort       string        `json:"suggested_sort"`
			UserReports         []interface{} `json:"user_reports"`
			SecureMedia         interface{}   `json:"secure_media"`
			IsRedditMediaDomain bool          `json:"is_reddit_media_domain"`
			LinkFlairText       interface{}   `json:"link_flair_text"`
			ID                  string        `json:"id"`
			BannedAtUtc         interface{}   `json:"banned_at_utc"`
			ModReasonTitle      interface{}   `json:"mod_reason_title"`
			ViewCount           interface{}   `json:"view_count"`
			Archived            bool          `json:"archived"`
			Clicked             bool          `json:"clicked"`
			MediaEmbed          struct {
			} `json:"media_embed"`
			ReportReasons   interface{}   `json:"report_reasons"`
			Author          string        `json:"author"`
			NumCrossposts   int           `json:"num_crossposts"`
			Saved           bool          `json:"saved"`
			ModReports      []interface{} `json:"mod_reports"`
			CanModPost      bool          `json:"can_mod_post"`
			IsCrosspostable bool          `json:"is_crosspostable"`
			Pinned          bool          `json:"pinned"`
			Score           int           `json:"score"`
			ApprovedBy      interface{}   `json:"approved_by"`
			Over18          bool          `json:"over_18"`
			Hidden          bool          `json:"hidden"`
			Preview         struct {
				Images []struct {
					Source struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"source"`
					Resolutions []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"resolutions"`
					Variants struct {
					} `json:"variants"`
					ID string `json:"id"`
				} `json:"images"`
				Enabled bool `json:"enabled"`
			} `json:"preview"`
			Thumbnail           string      `json:"thumbnail"`
			Edited              bool        `json:"edited"`
			LinkFlairCSSClass   interface{} `json:"link_flair_css_class"`
			AuthorFlairCSSClass interface{} `json:"author_flair_css_class"`
			ContestMode         bool        `json:"contest_mode"`
			Gilded              int         `json:"gilded"`
			Downs               int         `json:"downs"`
			BrandSafe           bool        `json:"brand_safe"`
			SecureMediaEmbed    struct {
			} `json:"secure_media_embed"`
			RemovalReason         interface{} `json:"removal_reason"`
			PostHint              string      `json:"post_hint"`
			AuthorFlairText       interface{} `json:"author_flair_text"`
			Stickied              bool        `json:"stickied"`
			CanGild               bool        `json:"can_gild"`
			ThumbnailHeight       int         `json:"thumbnail_height"`
			ParentWhitelistStatus string      `json:"parent_whitelist_status"`
			Name                  string      `json:"name"`
			Spoiler               bool        `json:"spoiler"`
			Permalink             string      `json:"permalink"`
			SubredditType         string      `json:"subreddit_type"`
			Locked                bool        `json:"locked"`
			HideScore             bool        `json:"hide_score"`
			Created               float64     `json:"created"`
			URL                   string      `json:"url"`
			WhitelistStatus       string      `json:"whitelist_status"`
			Quarantine            bool        `json:"quarantine"`
			Title                 string      `json:"title"`
			CreatedUtc            float64     `json:"created_utc"`
			SubredditNamePrefixed string      `json:"subreddit_name_prefixed"`
			Ups                   int         `json:"ups"`
			Media                 interface{} `json:"media"`
			NumComments           int         `json:"num_comments"`
			IsSelf                bool        `json:"is_self"`
			Visited               bool        `json:"visited"`
			ModNote               interface{} `json:"mod_note"`
			IsVideo               bool        `json:"is_video"`
			Distinguished         interface{} `json:"distinguished"`
		} `json:"data"`
	} `json:"children"`
}