package spec

// GetUsersActivedSpec returns a Specification that is true if the user is active.
func GetUsersActivedSpec() Specification {
	return And(Equal("active", true))
}

func GetUserWithEmailSpec(email string) Specification {
	return And(Equal("email", email))
}

func GetUserNameSpec(name string) Specification {
	return And(Equal("name", name))
}
