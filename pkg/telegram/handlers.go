package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Command() {
	case commandStart:
		msg.Text = ""
		b.client.UrlParams.ClearParams()
		return b.handleStartCommand(message)
	default:
		_, error := b.bot.Send(msg)
		return error
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)

	if err != nil {
		return b.initAuthorizationProcess(message)
	}

	b.openBaseKeyboard(message)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	switch b.mode {
	case "search":
		b.client.UrlParams.SetSearch(msg.Text)
		error := b.handleGetVacancies(message)
		return error
	case "salary":
		b.client.UrlParams.SetSalary(msg.Text)
		error := b.handleGetVacancies(message)
		return error
	case "area":
		areaId := SearchAreaByName(msg.Text)
		b.client.UrlParams.SetArea(areaId)
		error := b.handleGetVacancies(message)
		return error
	case "schedule":
		scheduleId := getScheduleId(message.Text)

		if scheduleId == "unknown" {
			msg.Text = "Ты ввел неизвестный график :("
			b.bot.Send(msg)
		} else {
			b.client.UrlParams.SetSchedule(scheduleId)
			error := b.handleGetVacancies(message)
			return error
		}
	case "experience":
		experienceId := getExperienceId(message.Text)

		if experienceId == "unknown" {
			msg.Text = "Ты ввел неизвестный опыт работы :("
			b.bot.Send(msg)
		} else {
			b.client.UrlParams.SetExperience(experienceId)
			error := b.handleGetVacancies(message)
			return error
		}
	}

	return nil
}

func SearchAreaByName(name string) string {
	for _, area := range AllAreas {
		if strings.Contains(area.Name, name) {
			return area.Id
		}
	}

	return "113"
}

func getScheduleId(schedule string) string {
	switch schedule {
	case scheduleCommands[0]:
		return "fullDay"
	case scheduleCommands[1]:
		return "shift"
	case scheduleCommands[2]:
		return "flexible"
	case scheduleCommands[3]:
		return "remote"
	case scheduleCommands[4]:
		return "flyInFlyOut"
	}

	return "unknown"
}

func getExperienceId(experience string) string {
	switch experience {
	case experienceCommands[0]:
		return "noExperience"
	case experienceCommands[1]:
		return "between1And3"
	case experienceCommands[2]:
		return "between3And6"
	case experienceCommands[3]:
		return "moreThan6"

	}

	return "unknown"
}
