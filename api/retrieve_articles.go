package api

type RetrieveArticlesOption struct {
	Page         int
	PerPage      int
	Tag          string
	Username     string
	State        State
	Top          int
	CollectionId int
}

type State string

const (
	StateFresh  State = "fresh"
	StateRising       = "rising"
	StateAll          = "all"
)

type Article struct {
	TypeOf               string       `json:"type_of"`
	Id                   int          `json:"id"`
	Title                string       `json:"title"`
	Description          string       `json:"description"`
	CoverImage           string       `json:"cover_image"`
	ReadablePublishData  string       `json:"readable_publish_data"`
	SocialImage          string       `json:"social_image"`
	TagList              []string     `json:"tag_list"`
	Tags                 string       `json:"tags"`
	Slug                 string       `json:"slug"`
	Path                 string       `json:"path"`
	Url                  string       `json:"url"`
	CanonicalUrl         string       `json:"canonical_url"`
	CommentsCount        int          `json:"comments_count"`
	PublicReactionsCount int          `json:"public_reactions_count"`
	CreatedAt            string       `json:"created_at"`
	EditedAt             string       `json:"edited_at"`
	CrosspostedAt        string       `json:"crossposted_at"`
	PublishedAt          string       `json:"published_at"`
	LastCommentAt        string       `json:"last_comment_at"`
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

type Articles []*Article

func (c *Client) RetrieveArticles(option *RetrieveArticlesOption) (*Articles, error) {
	res := &Articles{}
	err := GetQuery("/articles", option, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
