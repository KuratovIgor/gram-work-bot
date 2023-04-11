package graphqldb

import (
	"context"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	"github.com/machinebox/graphql"
	"strconv"
)

var createFiltersRequest = graphql.NewRequest(`
	mutation CreateFilters($chat_id: String!) {
  		default_filterCreate(data: {chat_id:$chat_id}) {
    		chat_id
  		}
	}
`)

var getFiltersRequest = graphql.NewRequest(`
	query GetFilters($chat_id: String!) {
  		default_filter(chat_id:$chat_id) {
			search
			salary
			schedule
			experience
			area_id
  		}
	}
`)

func (g *GraphqlRepository) CreateDefaultFilters(chatId int64) error {
	ctx := context.Background()

	createFiltersRequest.Var("chat_id", strconv.FormatInt(chatId, 10))

	var respData map[string]map[string]string
	err := g.graphqlClient.Run(ctx, createFiltersRequest, &respData)
	if err != nil {
		return err
	}

	return nil
}

func (g *GraphqlRepository) GetDefaultFilters(chatId int64) (*repository.DefaultFilterType, error) {
	ctx := context.Background()

	getFiltersRequest.Var("chat_id", strconv.FormatInt(chatId, 10))

	var respData map[string]repository.DefaultFilterType
	err := g.graphqlClient.Run(ctx, getFiltersRequest, &respData)
	if err != nil {
		return nil, err
	}

	filters := &repository.DefaultFilterType{
		Search:     respData["default_filter"].Search,
		Salary:     respData["default_filter"].Salary,
		Schedule:   respData["default_filter"].Schedule,
		Experience: respData["default_filter"].Experience,
		AreaId:     respData["default_filter"].AreaId,
	}

	return filters, nil
}
