package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
)

const authURI = "https://hh.ru/oauth/authorize"

func (b *Bot) generateAuthorizationLink(config *config.Config) string {
	return authURI + "?response_type=code&client_id=" + config.ClientID
}
