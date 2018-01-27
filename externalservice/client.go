package externalservice

import (
	//"io"
)

// Post is the data structure representing the data sent and received from the
// external service
type Post struct {
	ID int `json:"id"` // the primary key

	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// Client represents the client interface to the external service
type Client interface {
	GET(id int) (*Post, error)
	POST(id int, post *Post) (*Post, error)
}
