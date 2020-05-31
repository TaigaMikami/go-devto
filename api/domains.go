package api

type Article struct {
	TypeOf               string       `json:"type_of"`
	Id                   int          `json:"id"`
	Title                string       `json:"title"`
	Description          string       `json:"description"`
	CoverImage           string       `json:"cover_image"`
	Slug                 string       `json:"slug"`
	Path                 string       `json:"path"`
	Url                  string       `json:"url"`
	CanonicalUrl         string       `json:"canonical_url"`
	CommentsCount        int          `json:"comments_count"`
	PublicReactionsCount int          `json:"public_reactions_count"`
	PublishedAt          string       `json:"published_at"`
	PublishedTimestamp   string       `json:"published_timestamp"`
	User                 User         `json:"user"`
	Organization         Organization `json:"organization"`
	FlareTag             FlareTag     `json:"flare_tag"`
}

type User struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	WebsiteUrl      string `json:"website_url"`
	ProfileImage    string `json:"profile_image"`
	ProfileImage90  string `json:"profile_image_90"`
}

type Organization struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Slug           string `json:"slug"`
	ProfileImage   string `json:"profile_image"`
	ProfileImage90 string `json:"profile_image_90"`
}

type FlareTag struct {
	Name         string `json:"name"`
	BgColorHex   string `json:"bg_color_hex"`
	TextColorHex string `json:"text_color_hex"`
}

type VideoArticle struct {
	TypeOf                 string `json:"type_of"`
	Id                     int    `json:"id"`
	Path                   string `json:"path"`
	CloudinaryVideoUrl     string `json:"cloudinary_video_url"`
	Title                  string `json:"title"`
	UserId                 int    `json:"user_id"`
	VideoDurationInMinutes string `json:"video_duration_in_minutes"`
	User                   User   `json:"user"`
}

type Comment struct {
	TypeOf   string     `json:"type_of"`
	IdCode   string     `json:"id_code"`
	BodyHtml string     `json:"body_html"`
	User     User       `json:"user"`
	Children []*Comment `json:"children"`
}

type Follower struct {
	TypeOf       string `json:"type_of"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Username     string `json:"username"`
	ProfileImage string `json:"profile_image"`
}
