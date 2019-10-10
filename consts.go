package goapi

import "time"

const PublishedRequest string = "https://dev.to/api/articles?"

// Articles is a list of articles. Yeah...
type Articles []struct {
	TypeOf                 string       `json:"type_of"`
	ID                     int          `json:"id"`
	Title                  string       `json:"title"`
	Description            string       `json:"description"`
	CoverImage             string       `json:"cover_image"`
	Published              bool         `json:"published"`
	PublishedAt            time.Time    `json:"published_at"`
	TagList                []string     `json:"tag_list"`
	Tags                   []string     `json:"tags"`
	Slug                   string       `json:"slug"`
	Path                   string       `json:"path"`
	URL                    string       `json:"url"`
	CanonicalURL           string       `json:"canonical_url"`
	CommentsCount          int          `json:"comments_count"`
	PositiveReactionsCount int          `json:"positive_reactions_count"`
	PublishedTimestamp     time.Time    `json:"published_timestamp"`
	User                   User         `json:"user"`
	Organization           Organization `json:"organization,omitempty"`
	FlareTag               FlareTag     `json:"flare_tag,omitempty"`
}

// Article just an article
type Article struct {
	TypeOf                 string       `json:"type_of"`
	ID                     int          `json:"id"`
	Title                  string       `json:"title"`
	Description            string       `json:"description"`
	CoverImage             string       `json:"cover_image"`
	Published              bool         `json:"published"`
	PublishedAt            time.Time    `json:"published_at"`
	TagList                string       `json:"tag_list"`
	Tags                   []string     `json:"tags"`
	Slug                   string       `json:"slug"`
	Path                   string       `json:"path"`
	URL                    string       `json:"url"`
	CanonicalURL           string       `json:"canonical_url"`
	CommentsCount          int          `json:"comments_count"`
	PositiveReactionsCount int          `json:"positive_reactions_count"`
	PublishedTimestamp     time.Time    `json:"published_timestamp"`
	User                   User         `json:"user"`
	Organization           Organization `json:"organization,omitempty"`
	FlareTag               FlareTag     `json:"flare_tag,omitempty"`
}

// User struct contains information about user
type User struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	WebsiteURL      string `json:"website_url"`
	ProfileImage    string `json:"profile_image"`    // 640x640
	ProfileImage90  string `json:"profile_image_90"` // 90x90
}

// Organization struct contains information about organization (may be empty)
type Organization struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Slug           string `json:"slug"`
	ProfileImage   string `json:"profile_image"`    // 640x640
	ProfileImage90 string `json:"profile_image_90"` // 90x90
}

// FlareTag struct contains information about flare tag (may be empty)
type FlareTag struct {
	Name         string `json:"name"`
	BgColorHex   string `json:"bg_color_hex"`
	TextColorHex string `json:"text_color_hex"`
}

// QueryArticle used in GetPublishedArticles function
type QueryArticle struct {
	Page     int32  // Pagination page
	Tag      string // Articles that contain the requested tag
	Username string // Articles belonging to a User or Organization
	State    string // which articles are fresh or rising [fresh/rising]
	Top      int32  // most popular articles in the last N days
}

// Payload for CreateNewArticle function
type Payload struct {
	Article NewArticle `json:"article"`
}

// NewArticle struct used in Payload struct
type NewArticle struct {
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	BodyMarkdown   string   `json:"body_markdown"`
	Published      bool     `json:"published"`
	Series         string   `json:"series"`
	MainImage      string   `json:"main_image"`
	CanonicalURL   string   `json:"canonical_url"`
	Tags           []string `json:"tags"`
	OrganizationID int32    `json:"organization_id"`
}
