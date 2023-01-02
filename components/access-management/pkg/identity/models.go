package identity

type Result struct {
	Succeeded bool
	Errors    []string
}

type EmailCredential struct {
	Email    string
	Password string
}
