package authstatus

func AuthStatusCheck(email Email, i IAuthStatusRequests) bool {
	return i.Check(email.Email)
}
