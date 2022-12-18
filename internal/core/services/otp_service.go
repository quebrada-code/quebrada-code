package services

type Generator interface {
	RandomSecret(length int) string
}
