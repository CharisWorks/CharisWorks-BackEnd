package authstatus

func ExampleAuthStatus(email string) bool {
	return true
}

type ExampleAuthStatusRequests struct {
}

func (a ExampleAuthStatusRequests) Check(email string) (bool, error) {
	return ExampleAuthStatus(email), nil
}
