package api

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/types"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const baseURL = "https://api.hh.ru"
const vacancies = "/vacancies"

type VacanciesApiType struct {
	baseURL      string
	vacanciesURL string
}

func newVacanciesApi() *VacanciesApiType {
	return &VacanciesApiType{baseURL: baseURL, vacanciesURL: vacancies}
}

func (v *VacanciesApiType) GetVacancies(params *Params) (types.Vacancies, error) {
	res, _ := http.Get(baseURL + vacancies + params.GetQueryString())
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var vacancies types.Vacancies

	value := gjson.Get(string(body), "items")
	for _, item := range value.Array() {
		var vacancy types.Vacancy
		vacancy.Id = gjson.Get(item.String(), "id").String()
		vacancy.Name = gjson.Get(item.String(), "name").String()
		vacancy.Salary.From = gjson.Get(item.String(), "salary.from").String()
		vacancy.Salary.To = gjson.Get(item.String(), "salary.to").String()
		vacancy.Salary.Currency = gjson.Get(item.String(), "salary.currency").String()
		vacancy.Address.City = gjson.Get(item.String(), "address.city").String()
		vacancy.Address.Street = gjson.Get(item.String(), "address.street").String()
		vacancy.Address.Building = gjson.Get(item.String(), "address.building").String()
		vacancy.PublishedAt = gjson.Get(item.String(), "published_at").String()
		vacancy.Employer = gjson.Get(item.String(), "employer.name").String()
		vacancy.Requirement = gjson.Get(item.String(), "snippet.requirement").String()
		vacancy.Responsibility = gjson.Get(item.String(), "snippet.responsibility").String()
		vacancy.Schedule = gjson.Get(item.String(), "schedule.name").String()
		vacancy.AlternateUrl = gjson.Get(item.String(), "alternate_url").String()

		vacancies.Items = append(vacancies.Items, vacancy)
	}

	return vacancies, nil
}

var VacanciesApi = newVacanciesApi()
