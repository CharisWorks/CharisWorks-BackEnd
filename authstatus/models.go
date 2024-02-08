package authstatus

type IAuthStatusRequests interface {
	Check(string) bool
}
