package authstatus

type AuthStatusRequests struct {
}

func (a AuthStatusRequests) Check(email string) bool {
	return ExampleAuthStatus(email)
}

func AuthStatusCheck(email string, i IAuthStatusRequests) bool {
	return i.Check(email)
}
