package api

import (
	"net/url"
	"strconv"
)

type Params struct {
	per_page int
	Page     int
	search   string
	salary   string
	currency string
}

func NewParams() *Params {
	return &Params{per_page: 5, Page: 0, search: "", salary: "", currency: "RUR"}
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

func (p *Params) GetQueryString() string {
	params := "?per_page=" + strconv.Itoa(p.per_page) +
		"&page=" + strconv.Itoa(p.Page) +
		"&text=" + url.QueryEscape(p.search) +
		"&currency=" + url.QueryEscape(p.currency)

	if p.salary != "" {
		params = params + "&salary=" + url.QueryEscape(p.salary)
	}

	return params
}

func (p *Params) ClearParams() {
	p.Page = 0
	p.search = ""
	p.salary = ""
}

func (p *Params) ClearFilters() {
	p.salary = ""
}
