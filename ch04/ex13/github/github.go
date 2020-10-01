// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
	Query      string `json:"-"`
}

type Issue struct {
	Number        int
	RepositoryURL string `json:"repository_url"`
	HTMLURL       string `json:"html_url"`
	Title         string
	User          *User
	Milestone     *Milestone
	CreatedAt     time.Time `json:"created_at"`
	Body          string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Title       string
	Number      int
	Description string
	HTMLURL     string `json:"html_url"`
}

//!-
