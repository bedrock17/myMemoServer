package data

type memo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type memoList struct {
	Memos []string `json:"memos"`
}
