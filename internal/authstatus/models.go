package authstatus

type IAuthStatusRequests interface {
	Check(string) (bool, error)
}
type Email struct {
	Email string `json:"email" binding:"required"`
}
