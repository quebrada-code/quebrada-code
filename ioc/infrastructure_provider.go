package ioc

import (
	"quebrada_api/internal/config"
	"quebrada_api/internal/infrastructure/senders"
	"quebrada_api/pkg/identity"
)

func ProvidePasswordHash(salt string) identity.IPasswordHasher {
	return identity.NewPasswordHasher("")
}

func ProviderEmailSender(config config.STMPConfig) *senders.EmailSender {
	return senders.NewEmailSender(config)
}
