package web

type NoteResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Tags []string `json:"tags"`
	Body string `json:"body"`
}
