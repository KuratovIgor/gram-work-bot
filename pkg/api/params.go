package api

import (
	"net/url"
	"strconv"
)

type Params struct {
	per_page   int
	Page       int
	search     string
	salary     string
	currency   string
	area       string
	schedule   []string
	experience []string
}

func NewParams() *Params {
	return &Params{per_page: 5, Page: 0, search: "", salary: "", currency: "RUR", area: "113", schedule: []string{}, experience: []string{}}
}

func (p *Params) SetPage(page int) {
	p.Page = page
}

func (p *Params) SetSearch(search string) {
	p.search = search
}

func (p *Params) SetSalary(salary string) {
	p.salary = salary
}

func (p *Params) SetArea(area string) {
	p.area = area
}

func (p *Params) SetSchedule(schedule string) {
	p.schedule = append(p.schedule, schedule)
}

func (p *Params) SetExperience(experience string) {
	p.experience = append(p.experience, experience)
}

func (p *Params) GetQueryString() string {
	params := "?per_page=" + strconv.Itoa(p.per_page) +
		"&page=" + strconv.Itoa(p.Page) +
		"&text=" + url.QueryEscape(p.search) +
		"&currency=" + url.QueryEscape(p.currency) +
		"&area=" + url.QueryEscape(p.area)

	if p.salary != "" {
		params = params + "&salary=" + url.QueryEscape(p.salary)
	}

	if len(p.schedule) != 0 {
		for i, _ := range p.schedule {
			params = params + "&schedule=" + url.QueryEscape(p.schedule[i])
		}
	}

	if len(p.experience) != 0 {
		for i, _ := range p.experience {
			params = params + "&experience=" + url.QueryEscape(p.experience[i])
		}
	}

	return params
}

func (p *Params) ClearParams() {
	p.Page = 0
	p.search = ""
	p.salary = ""
	p.area = "113"
	p.schedule = []string{}
	p.experience = []string{}
}

func (p *Params) ClearFilters() {
	p.salary = ""
	p.area = "113"
	p.schedule = []string{}
	p.experience = []string{}
}
