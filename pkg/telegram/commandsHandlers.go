package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"time"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		_, error := b.bot.Send(msg)
		return error
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	b.client.UrlParams.ClearParams()

	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message)
	}

	b.openBaseKeyboard(message)

	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch b.mode {
	case "search":
		return b.handleSearch(message)
	case "salary":
		return b.handleFilterBySalary(message)
	case "area":
		return b.handleFilterByArea(message)
	case "schedule":
		return b.handleFilterBySchedule(message)
	case "experience":
		return b.handleFilterByExperience(message)
	case "message":
		b.applyMessage = message.Text
		applyErr := b.handleApplyToJob(message)
		if applyErr != nil {
			return applyErr
		}
	}

	return nil
}

func (b *Bot) handleInlineCommand(update tgbotapi.Update) error {
	switch b.mode {
	case "chooseResume":
		applyErr := b.handleApplyToJobByResume(update.CallbackQuery.Message, update.CallbackQuery.Data)
		if applyErr != nil {
			return applyErr
		}

		vacancy, reqErr := b.getVacancy(update.CallbackQuery.Message.Chat.ID, b.chosenVacancyId)
		if reqErr != nil {
			return reqErr
		}

		salaryFrom, _ := strconv.Atoi(vacancy.Salary.From)
		salaryTo, _ := strconv.Atoi(vacancy.Salary.To)

		infoAboutMe, infoErr := b.getInfoAboutMe(update.CallbackQuery.Message.Chat.ID)
		if infoErr != nil {
			return infoErr
		}

		savingErr := b.graphqlRepository.SaveApplyToJob(infoAboutMe.UserID, vacancy.Id, vacancy.Name, vacancy.Employer, vacancy.AlternateUrl, vacancy.Area, "Отклик", salaryFrom, salaryTo, time.Now().Format("01-02-2006"))
		if savingErr != nil {
			return savingErr
		}

		b.chosenVacancyId = ""
	case "apply":
		b.handleSendApplyMessage(update.CallbackQuery.Message, update.CallbackQuery.Data)
	default:
		b.handleSendApplyMessage(update.CallbackQuery.Message, update.CallbackQuery.Data)
	}

	return nil
}
