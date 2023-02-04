package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var params = NewParams()

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Command() {
	case commandStart:
		msg.Text = ""
		b.openBaseKeyboard(message)
		_, error := b.bot.Send(msg)
		return error
	default:
		_, error := b.bot.Send(msg)
		return error
	}
}

func (b *Bot) handleBaseKeybord(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Text {
	case baseCommands[0]:
		params.setPage(0)
		error := b.handleGetVacancies(message)
		b.openVacanciesKeyboard(message)
		return error
	case baseCommands[1]:
		msg.Text = "Создаем резюме..."
		_, error := b.bot.Send(msg)
		return error
	}

	return nil
}

func (b *Bot) handleVacanciesKeybord(message *tgbotapi.Message) error {
	switch message.Text {
	case vacanciesCommands[0]:
		params.setPage(params.page + 1)
		error := b.handleGetVacancies(message)
		return error
	case vacanciesCommands[2]:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Введите должность для поиска")
		b.bot.Send(msg)
		b.mode = "search"
	case vacanciesCommands[3]:
		b.mode = ""
		params.clearParams()
		b.openBaseKeyboard(message)
	}

	return nil
}

func (b *Bot) handleGetVacancies(message *tgbotapi.Message) error {
	vacancies, err := getVacancies(params)

	if err != nil {
		return err
	}

	for _, item := range vacancies.items {
		var buttons = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Откликнуться", "Откликнуться"),
				tgbotapi.NewInlineKeyboardButtonURL("Перейти", item.alternateUrl),
			),
		)

		msg := tgbotapi.NewMessage(message.Chat.ID, "ДОЛЖНОСТЬ:\n"+item.name+"\n\nЗАРПЛАТА:\n от "+item.salary.from+" до "+item.salary.to+" "+item.salary.currency+"\n\nРАБОТОДАТЕЛЬ:\n"+item.employer+"\n\nОПИСАНИЕ:\n"+item.responsibility+"\n\nТРЕБОВАНИЯ:\n"+item.requirement+"\n\nАДРЕС:\n"+item.address.city+" "+item.address.street+" "+item.address.building+"\n\nГРАФИК:\n"+item.schedule+"\n\nОПУБЛИКОВАНО:\n"+item.publishedAt)
		msg.ReplyMarkup = buttons
		b.bot.Send(msg)
	}

	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	switch b.mode {
	case "search":
		params.setSearch(msg.Text)
		error := b.handleGetVacancies(message)
		return error
	}

	return nil
}
