package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/api"
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

var params = api.NewParams()

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message, config *config.Config) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Command() {
	case commandStart:
		msg.Text = ""
		params.ClearParams()
		return b.handleStartCommand(message, config)
	default:
		_, error := b.bot.Send(msg)
		return error
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message, config *config.Config) error {
	_, err := b.getAccessToken(message.Chat.ID)

	if err != nil {
		return b.initAuthorizationProcess(config, message)
	}

	b.openBaseKeyboard(message)
	return err
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
			params.SetSchedule(scheduleId)
			error := b.handleGetVacancies(message)
			return error
		}
	case "experience":
		experienceId := getExperienceId(message.Text)

		if experienceId == "unknown" {
			msg.Text = "Ты ввел неизвестный опыт работы :("
			b.bot.Send(msg)
		} else {
			params.SetExperience(experienceId)
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
