package identity

type ISender interface {
	Send(to []string, subject string, template string) error
}
