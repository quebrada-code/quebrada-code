package ioc

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/gorm"
	"quebrada_api/internal/communication"

	"quebrada_api/internal/domain/entities"
	domain "quebrada_api/internal/domain/repositories"
	"quebrada_api/internal/domain/services"
	"quebrada_api/internal/infrastructure/senders"
	"quebrada_api/pkg/identity"
)

func ProvideAuthService(
	repository domain.IRepository[entities.User],
	passwordHasher identity.IPasswordHasher,
	DB *gorm.DB,
	sender senders.ISender) services.IAuthService {
	return services.NewAuthService(repository, passwordHasher, sender, DB)
}

func ProvideProblemService(
	repository domain.IRepository[entities.Problem],
	kafkaProducer *kafka.Producer) services.IProblemService {
	publisher := communication.NewKafkaPublisher(kafkaProducer)

	return services.NewProblemService(repository, publisher)
}
