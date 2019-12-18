package users

type FindRequest struct {
	ID    int    `uri:"id"`
	Email string `form:"email"`
}
