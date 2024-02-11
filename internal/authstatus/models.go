package authstatus

type IAuthStatusRequests interface {
	Check(string) bool
}
type Email struct {
	Email string `json:"email" binding:"required"`
}
