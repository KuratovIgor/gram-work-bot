package telegram

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://api.hh.ru"

const vacancies = "/vacancies/1"

func getVacancies() (any, error) {
	res, _ := http.Get("https://api.hh.ru/vacancies")
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	value := gjson.Get(string(body), "items")
	for _, name := range value.Array() {
		log.Println(gjson.Get(name.String(), "name"))
		break
	}

	return nil, nil
}
