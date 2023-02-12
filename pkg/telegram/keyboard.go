package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

// ГЛАВНЫЕ КОМАНДЫ
var baseCommands = []string{"Поиск вакансий", "Мои резюме"}
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
	msg := tgbotapi.NewMessage(message.Chat.ID, "Выбери действие")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = vacanciesCommandsKeyboard

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

// КОМАНДЫ ДЛЯ ФИЛЬТРАЦИИ ВАКАНСИЙ
var filterCommands = []string{"З/П", "Город", "График", "Опыт", "Сбросить фильтры", "К вакансиям"}
var filterCommandsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(filterCommands[0]),
		tgbotapi.NewKeyboardButton(filterCommands[1]),
		tgbotapi.NewKeyboardButton(filterCommands[2]),
		tgbotapi.NewKeyboardButton(filterCommands[3]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(filterCommands[4]),
		tgbotapi.NewKeyboardButton(filterCommands[5]),
	),
)

func (b *Bot) openFilterKeyboard(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите параметр для фильтрации")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = filterCommandsKeyboard

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

// КОМАНДЫ ДЛЯ ВЫБОРА ГРАФИКА
var scheduleCommands = []string{"Полный день", "Сменный график", "Гибкий график", "Удаленная работа", "Вахтовый метод", "К фильтрам"}
var scheduleCommandsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(scheduleCommands[0]),
		tgbotapi.NewKeyboardButton(scheduleCommands[1]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(scheduleCommands[2]),
		tgbotapi.NewKeyboardButton(scheduleCommands[3]),
		tgbotapi.NewKeyboardButton(scheduleCommands[4]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(scheduleCommands[5]),
	),
)

func (b *Bot) openScheduleKeyboard(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Укажите график")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = scheduleCommandsKeyboard

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

// КОМАНДЫ ДЛЯ ВЫБОРА ОПЫТА РАБОТЫ
var experienceCommands = []string{"Нет опыта", "От 1 года до 3 лет", "От 3 лет до 6 лет", "Более 6 лет", "Вернуться"}
var experienceCommandsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(experienceCommands[0]),
		tgbotapi.NewKeyboardButton(experienceCommands[1]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(experienceCommands[2]),
		tgbotapi.NewKeyboardButton(experienceCommands[3]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(experienceCommands[4]),
	),
)

func (b *Bot) openExperienceKeyboard(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Укажите опыт работы")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = experienceCommandsKeyboard

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
