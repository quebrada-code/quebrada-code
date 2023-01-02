package domain_errors

type UserAlreadyExistsError struct{}

func (UserAlreadyExistsError) Error() string {
	return "Já existe usuário cadastrado com esse e-mail"
}

type HashPasswordError struct{}

func (HashPasswordError) Error() string {
	return "Falha em encriptar senha."
}

type InsertUserError struct{}

func (InsertUserError) Error() string {
	return "Falha ao inserir usuário"
}
