package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

var params = api.NewParams()

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Command() {
	case commandStart:
		msg.Text = ""
		params.ClearParams()
		b.openBaseKeyboard(message)
		_, error := b.bot.Send(msg)
		return error
	default:
		_, error := b.bot.Send(msg)
		return error
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	switch b.mode {
	case "search":
		params.SetSearch(msg.Text)
		error := b.handleGetVacancies(message)
		return error
	case "salary":
		params.SetSalary(msg.Text)
		error := b.handleGetVacancies(message)
		return error
	case "area":
		areaId := SearchAreaByName(msg.Text)
		params.SetArea(areaId)
		error := b.handleGetVacancies(message)
		return error
	case "schedule":
		scheduleId := getScheduleId(message.Text)

		if scheduleId == "unknown" {
			msg.Text = "Ты ввел неизвестный график :("
			b.bot.Send(msg)
		} else {
			params.SetSchedule(getScheduleId(message.Text))
			error := b.handleGetVacancies(message)
			return error
		}
	}

	return nil
}

func SearchAreaByName(name string) string {
	for _, area := range api.Areas {
		if strings.Contains(area.Name, name) {
			return area.Id
		}
	}

	return "113"
}

func getScheduleId(schedule string) string {
	switch schedule {
	case "Полный день":
		return "fullDay"
	case "Сменный график":
		return "shift"
	case "Гибкий график":
		return "flexible"
	case "Удаленная работа":
		return "remote"
	case "Вахтовый метод":
		return "flyInFlyOut"
	}

	return "unknown"
}
