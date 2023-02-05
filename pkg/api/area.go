package api

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/types"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

var Areas []types.AreaType

func GetAllAreas() error {
	res, _ := http.Get(baseURL + areaURL)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	largeAreas := gjson.Parse(string(body))

	for _, largeArea := range largeAreas.Array() {
		areas := gjson.Get(largeArea.String(), "areas")

		for _, item := range areas.Array() {
			var area types.AreaType

			area.Name = gjson.Get(item.String(), "name").String()
			area.Id = gjson.Get(item.String(), "id").String()

			Areas = append(Areas, area)
		}
	}

	return nil
}
