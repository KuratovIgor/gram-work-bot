package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var commands = []string{"Поиск вакансий", "Создать резюме"}

var commandsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(commands[0]),
		tgbotapi.NewKeyboardButton(commands[1]),
	),
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Command() {
	case commandStart:
		msg.Text = ""
		b.openKeyboard(message)
		_, error := b.bot.Send(msg)
		return error
	default:
		_, error := b.bot.Send(msg)
		return error
	}
}

func (b *Bot) handleCommandFromKeybord(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Text {
	case commands[0]:
		b.handleGetVacancies(&msg)
		_, error := b.bot.Send(msg)
		return error
	case commands[1]:
		msg.Text = "Создаем резюме..."
		_, error := b.bot.Send(msg)
		return error
	default:
		_, error := b.bot.Send(msg)
		return error
	}
}

func (b *Bot) openKeyboard(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты можешь пользоваться всем функционалом, используя команды!")
	msg.ReplyMarkup = commandsKeyboard

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.bot.Send(msg)
}

func (b *Bot) handleGetVacancies(message *tgbotapi.MessageConfig) {
	_, err := getVacancies()

	if err != nil {
		log.Println(err)
	}

	//log.Println(response)
}
