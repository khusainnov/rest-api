package posts

type Post struct {
	PostId   int    `json:"post_id"`
	Header   string `json:"header"`
	Body     string `json:"body"`
	AuthorId int    `json:"author_id"`
}

//type Author struct {
//	AuthorID   int       `json:"author_id"`
//	AuthorName string `json:"author_name"`
//}
