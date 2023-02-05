package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) handleBaseKeyboard(message *tgbotapi.Message) error {
	switch message.Text {
	case baseCommands[0]:
		params.SetPage(0)
		error := b.handleGetVacancies(message)
		b.openVacanciesKeyboard(message)
		return error
	}

	return nil
}

func (b *Bot) handleVacanciesKeyboard(message *tgbotapi.Message) error {
	switch message.Text {
	case vacanciesCommands[0]:
		params.SetPage(params.Page + 1)
		error := b.handleGetVacancies(message)
		return error
	case vacanciesCommands[1]:
		b.openFilterKeyboard(message)
	case vacanciesCommands[2]:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Введите должность для поиска")
		b.bot.Send(msg)
		b.mode = "search"
	case vacanciesCommands[3]:
		b.mode = ""
		params.ClearParams()
		b.openBaseKeyboard(message)
	}

	return nil
}

func (b *Bot) handleFiltersKeyboard(message *tgbotapi.Message) error {
	switch message.Text {
	case filterCommands[0]:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Введите сумму в рублях")
		b.bot.Send(msg)
		b.mode = "salary"
	case filterCommands[1]:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Введите название города")
		b.bot.Send(msg)
		b.mode = "area"
	case filterCommands[2]:
		b.openScheduleKeyboard(message)
	case filterCommands[3]:
		b.openExperienceKeyboard(message)
	case filterCommands[4]:
		params.ClearFilters()
		b.mode = ""
		error := b.handleGetVacancies(message)
		return error
	case filterCommands[5]:
		b.mode = ""
		b.openVacanciesKeyboard(message)
	}

	return nil
}

func (b *Bot) handleScheduleKeyboard(message *tgbotapi.Message) error {
	switch message.Text {
	case scheduleCommands[0], scheduleCommands[1], scheduleCommands[2], scheduleCommands[3], scheduleCommands[4]:
		msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		b.mode = "schedule"
		b.bot.Send(msg)
	case scheduleCommands[5]:
		b.mode = ""
		b.openFilterKeyboard(message)
	}

	return nil
}

func (b *Bot) handleExperienceKeyboard(message *tgbotapi.Message) error {
	switch message.Text {
	case experienceCommands[0], experienceCommands[1], experienceCommands[2], experienceCommands[3]:
		msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		b.mode = "experience"
		b.bot.Send(msg)
	case experienceCommands[4]:
		b.mode = ""
		b.openFilterKeyboard(message)
	}

	return nil
}