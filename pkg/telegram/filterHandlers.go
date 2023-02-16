package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleSearch(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.client.UrlParams.SetSearch(msg.Text)

	return b.displayVacancies(message)
}

func (b *Bot) handleFilterBySalary(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.client.UrlParams.SetSalary(msg.Text)

	return b.displayVacancies(message)
}

func (b *Bot) handleFilterByArea(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	areaId := searchAreaByName(msg.Text)
	b.client.UrlParams.SetArea(areaId)

	return b.displayVacancies(message)
}

func (b *Bot) handleFilterBySchedule(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	scheduleId := getScheduleIdByText(message.Text)
	if scheduleId != "unknown" {
		b.client.UrlParams.SetSchedule(scheduleId)
		return b.displayVacancies(message)
	}

	msg.Text = "Ты ввел неизвестный график :("
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleFilterByExperience(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	experienceId := getExperienceIdByText(message.Text)

	if experienceId != "unknown" {
		b.client.UrlParams.SetExperience(experienceId)
		return b.displayVacancies(message)
	}

	msg.Text = "Ты ввел неизвестный опыт работы :("
	_, err := b.bot.Send(msg)

	return err
}
