package telegram

import (
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const baseURL = "https://api.hh.ru"
const vacancies = "/vacancies"

func getVacancies(params *Params) (Vacancies, error) {
	res, _ := http.Get(baseURL + vacancies + params.getQueryString())
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var vacancies Vacancies

	value := gjson.Get(string(body), "items")
	for _, item := range value.Array() {
		var vacancy Vacancy
		vacancy.id = gjson.Get(item.String(), "id").String()
		vacancy.name = gjson.Get(item.String(), "name").String()
		vacancy.salary.from = gjson.Get(item.String(), "salary.from").String()
		vacancy.salary.to = gjson.Get(item.String(), "salary.to").String()
		vacancy.salary.currency = gjson.Get(item.String(), "salary.currency").String()
		vacancy.address.city = gjson.Get(item.String(), "address.city").String()
		vacancy.address.street = gjson.Get(item.String(), "address.street").String()
		vacancy.address.building = gjson.Get(item.String(), "address.building").String()
		vacancy.publishedAt = gjson.Get(item.String(), "published_at").String()
		vacancy.employer = gjson.Get(item.String(), "employer.name").String()
		vacancy.requirement = gjson.Get(item.String(), "snippet.requirement").String()
		vacancy.responsibility = gjson.Get(item.String(), "snippet.responsibility").String()
		vacancy.schedule = gjson.Get(item.String(), "schedule.name").String()
		vacancy.alternateUrl = gjson.Get(item.String(), "alternate_url").String()

		vacancies.items = append(vacancies.items, vacancy)
	}

	return vacancies, nil
}
