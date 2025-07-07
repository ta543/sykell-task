package models

import "time"

type URLStatus string

const (
        StatusQueued     URLStatus = "queued"
        StatusProcessing URLStatus = "processing"
        StatusDone       URLStatus = "done"
        StatusError      URLStatus = "error"
        StatusStopped    URLStatus = "stopped"
)

type URL struct {
	ID            int64     `json:"id" db:"id"`
	Address       string    `json:"address" db:"address"`
	Title         string    `json:"title" db:"title"`
	HTMLVersion   string    `json:"html_version" db:"html_version"`
	H1Count       int       `json:"h1_count" db:"h1_count"`
	H2Count       int       `json:"h2_count" db:"h2_count"`
	H3Count       int       `json:"h3_count" db:"h3_count"`
	InternalLinks int       `json:"internal_links" db:"internal_links"`
	ExternalLinks int       `json:"external_links" db:"external_links"`
        BrokenLinks   int       `json:"broken_links" db:"broken_links"`
        BrokenLinksDetail []BrokenLink `json:"broken_links_detail,omitempty" db:"-"`
        HasLoginForm  bool      `json:"has_login_form" db:"has_login_form"`
	Status        URLStatus `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
