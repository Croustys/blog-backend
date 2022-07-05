package types

type UserS struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type PostS struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
