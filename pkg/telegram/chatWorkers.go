package telegram

import (
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) displayAuthorizeMessage(authorizeLink string, message *tgbotapi.Message) error {
	var button = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Авторизироваться", authorizeLink),
		),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет!\nДля начала поиска тебе нужно авторизироваться.\nПожалуйста, перейди по ссылке, нажав на кнопку ниже.\n\nПосле авторизации введи команду /start для начала работы.")
	msg.ReplyMarkup = button

	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) displayVacancies(vacancies headhunter.Vacancies, message *tgbotapi.Message) error {
	for _, item := range vacancies.Items {
		var buttons = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("Перейти", item.AlternateUrl),
			),
		)

		msg := tgbotapi.NewMessage(message.Chat.ID, "ДОЛЖНОСТЬ:\n"+item.Name+"\n\nЗАРПЛАТА:\n от "+item.Salary.From+" до "+item.Salary.To+" "+item.Salary.Currency+"\n\nГород:\n"+item.Area+"\n\nРАБОТОДАТЕЛЬ:\n"+item.Employer+"\n\nОПИСАНИЕ:\n"+item.Responsibility+"\n\nТРЕБОВАНИЯ:\n"+item.Requirement+"\n\nАДРЕС:\n"+item.Address.City+" "+item.Address.Street+" "+item.Address.Building+"\n\nГРАФИК:\n"+item.Schedule+"\n\nОПУБЛИКОВАНО:\n"+item.PublishedAt)
		msg.ReplyMarkup = buttons

		_, err := b.bot.Send(msg)
		if err != nil {
			return nil
		}
	}

	return nil
}
