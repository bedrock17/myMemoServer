package data

import "time"

type memo struct {
	path       string
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CratedTime time.Time `json:"createdtime"`
}

type memoList struct {
	Memos []string `json:"memos"`
}
