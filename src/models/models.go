package models

type Target struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Err   string `json:"err,omitempty"`
}
