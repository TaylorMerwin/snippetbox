package models

import (
	"database/sql"
	"time"
)

// Define a Snippet type to hold the data for an individial snippet.

type Snippet struct {
	ID	int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet to DB.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// Return snippet based on id.
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// Return 10 more recently created snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}