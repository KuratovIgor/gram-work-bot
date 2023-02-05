package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleGetVacancies(message *tgbotapi.Message) error {
	vacancies, err := api.VacanciesApi.GetVacancies(params)

	if err != nil {
		return err
	}

	for _, item := range vacancies.Items {
		var buttons = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Откликнуться", "Откликнуться"),
				tgbotapi.NewInlineKeyboardButtonURL("Перейти", item.AlternateUrl),
			),
		)

		msg := tgbotapi.NewMessage(message.Chat.ID, "ДОЛЖНОСТЬ:\n"+item.Name+"\n\nЗАРПЛАТА:\n от "+item.Salary.From+" до "+item.Salary.To+" "+item.Salary.Currency+"\n\nРАБОТОДАТЕЛЬ:\n"+item.Employer+"\n\nОПИСАНИЕ:\n"+item.Responsibility+"\n\nТРЕБОВАНИЯ:\n"+item.Requirement+"\n\nАДРЕС:\n"+item.Address.City+" "+item.Address.Street+" "+item.Address.Building+"\n\nГРАФИК:\n"+item.Schedule+"\n\nОПУБЛИКОВАНО:\n"+item.PublishedAt)
		msg.ReplyMarkup = buttons
		b.bot.Send(msg)
	}

	return nil
}
