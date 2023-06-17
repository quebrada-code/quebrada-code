package communication

type Subscriber[T interface{}] interface {
	Subscribe(topic string, handler func(event T) error) error
}
