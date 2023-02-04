package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

// ГЛАВНЫЕ КОМАНДЫ
var baseCommands = []string{"Поиск вакансий", "Создать резюме"}
var baseCommandsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(baseCommands[0]),
		tgbotapi.NewKeyboardButton(baseCommands[1]),
	),
)

func (b *Bot) openBaseKeyboard(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты можешь пользоваться всем функционалом, используя команды!")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = baseCommandsKeyboard

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

// КОМАНДЫ ВЗАИМОДЕЙСТВИЯ С ВАКАНСИЯМИ
var vacanciesCommands = []string{"Больше вакансий", "Фильтры", "Поиск", "Главная"}
var vacanciesCommandsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(vacanciesCommands[0]),
		tgbotapi.NewKeyboardButton(vacanciesCommands[1]),
		tgbotapi.NewKeyboardButton(vacanciesCommands[2]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(vacanciesCommands[3]),
	),
)

func (b *Bot) openVacanciesKeyboard(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "-------------------------------------------------------")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = vacanciesCommandsKeyboard

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
