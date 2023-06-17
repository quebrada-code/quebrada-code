package communication

type Publisher interface {
	Send(to string, data interface{}, key string) error
}
