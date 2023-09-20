package web

type NoteUpdateRequest struct {
	Id    int      `validate:"required"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
	Body  string   `json:"body"`
}
