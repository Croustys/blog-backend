package types

type UserS struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type PostS struct {
	AuthorEmail string `json:"authorEmail"`
	Username    string `json:"username"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}
type PostResponse struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
