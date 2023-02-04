package telegram

import (
	"net/url"
	"strconv"
)

type Params struct {
	per_page int
	page     int
	search   string
}

func NewParams() *Params {
	return &Params{per_page: 5, page: 0, search: ""}
}

func (p *Params) setPage(page int) {
	p.page = page
}

func (p *Params) setSearch(search string) {
	p.search = search
}

func (p *Params) getQueryString() string {
	return "?per_page=" + strconv.Itoa(p.per_page) +
		"&page=" + strconv.Itoa(p.page) +
		"&text=" + url.QueryEscape(p.search)
}

func (p *Params) clearParams() {
	p.page = 0
	p.search = ""
}
