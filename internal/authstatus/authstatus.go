package authstatus

func AuthStatusCheck(email Email, i IAuthStatusRequests) (bool, error) {
	return i.Check(email.Email)
}
